// Package follow berisi fitur follow/unfollow dan profil publik pengguna.
// Semua aturan (tidak boleh follow diri sendiri, admin tidak bisa diikuti,
// follow ganda diabaikan) dijalankan di backend.
package follow

import (
	"errors"
	"log"
	"net/http"
	"time"

	"codelingo/backend/internal/media"
	"codelingo/backend/internal/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

const (
	DefaultLimit = 20
	MaxLimit     = 100
)

type APIError struct {
	Status  int
	Message string
}

func (e *APIError) Error() string { return e.Message }

func newErr(status int, msg string) *APIError { return &APIError{Status: status, Message: msg} }

func internalErr(op string, err error) *APIError {
	log.Printf("follow: %s: %v", op, err)
	return newErr(http.StatusInternalServerError, "terjadi kesalahan internal")
}

type Service struct {
	DB *gorm.DB
}

// PublicProfile hanya berisi data yang aman dilihat semua user
// (tanpa email, umur, role).
type PublicProfile struct {
	Username       string    `json:"username"`
	Name           string    `json:"name"`
	Level          int       `json:"level"`
	XP             int       `json:"xp"`
	AvatarURL      string    `json:"avatar_url"`
	BannerURL      string    `json:"banner_url"`
	JoinedAt       time.Time `json:"joined_at"`
	FollowersCount int64     `json:"followers_count"`
	FollowingCount int64     `json:"following_count"`
	IsFollowing    bool      `json:"is_following"` // viewer mengikuti user ini
	FollowsYou     bool      `json:"follows_you"`  // user ini mengikuti viewer
	IsMe           bool      `json:"is_me"`
}

type Card struct {
	Username    string `json:"username"`
	Name        string `json:"name"`
	XP          int    `json:"xp"`
	AvatarURL   string `json:"avatar_url"`
	IsFollowing bool   `json:"is_following"` // viewer mengikuti user ini
	IsMe        bool   `json:"is_me"`
}

type ListResponse struct {
	Users  []Card `json:"users"`
	Total  int64  `json:"total"`
	Limit  int    `json:"limit"`
	Offset int    `json:"offset"`
}

// target memuat user yang dituju. Akun admin dianggap tidak ada.
func (s *Service) target(username string) (*models.User, error) {
	var u models.User
	err := s.DB.Where("username = ? AND role <> ?", username, models.RoleAdmin).Take(&u).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, newErr(http.StatusNotFound, "pengguna tidak ditemukan")
		}
		return nil, internalErr("find user", err)
	}
	return &u, nil
}

func (s *Service) count(where string, id uint) (int64, error) {
	var n int64
	if err := s.DB.Model(&models.Follow{}).Where(where, id).Count(&n).Error; err != nil {
		return 0, err
	}
	return n, nil
}

func (s *Service) exists(followerID, followingID uint) (bool, error) {
	var n int64
	err := s.DB.Model(&models.Follow{}).
		Where("follower_id = ? AND following_id = ?", followerID, followingID).Count(&n).Error
	return n > 0, err
}

// Profile = profil publik satu user beserta jumlah followers/following.
func (s *Service) Profile(viewer *models.User, username string) (*PublicProfile, error) {
	u, err := s.target(username)
	if err != nil {
		return nil, err
	}
	followers, err := s.count("following_id = ?", u.ID)
	if err != nil {
		return nil, internalErr("count followers", err)
	}
	following, err := s.count("follower_id = ?", u.ID)
	if err != nil {
		return nil, internalErr("count following", err)
	}

	p := &PublicProfile{
		Username: u.Username, Name: u.Name, Level: u.Level, XP: u.XP,
		JoinedAt: u.CreatedAt, FollowersCount: followers, FollowingCount: following,
		IsMe: u.ID == viewer.ID,
	}
	p.AvatarURL, p.BannerURL = (&media.Service{DB: s.DB}).URLs(u.ID, u.Username)

	if !p.IsMe {
		if p.IsFollowing, err = s.exists(viewer.ID, u.ID); err != nil {
			return nil, internalErr("check following", err)
		}
		if p.FollowsYou, err = s.exists(u.ID, viewer.ID); err != nil {
			return nil, internalErr("check follows you", err)
		}
	}
	return p, nil
}

// Follow: idempotent. Mengikuti dua kali tetap berhasil dan tidak membuat baris ganda.
func (s *Service) Follow(viewer *models.User, username string) (*PublicProfile, error) {
	u, err := s.target(username)
	if err != nil {
		return nil, err
	}
	if u.ID == viewer.ID {
		return nil, newErr(http.StatusBadRequest, "kamu tidak bisa mengikuti diri sendiri")
	}
	row := models.Follow{FollowerID: viewer.ID, FollowingID: u.ID, CreatedAt: time.Now()}
	err = s.DB.Clauses(clause.OnConflict{DoNothing: true}).Omit(clause.Associations).Create(&row).Error
	if err != nil {
		return nil, internalErr("create follow", err)
	}
	return s.Profile(viewer, username)
}

// Unfollow: idempotent.
func (s *Service) Unfollow(viewer *models.User, username string) (*PublicProfile, error) {
	u, err := s.target(username)
	if err != nil {
		return nil, err
	}
	err = s.DB.Where("follower_id = ? AND following_id = ?", viewer.ID, u.ID).Delete(&models.Follow{}).Error
	if err != nil {
		return nil, internalErr("delete follow", err)
	}
	return s.Profile(viewer, username)
}

type cardRow struct {
	ID       uint
	Username string
	Name     string
	XP       int
}

// List: kind "followers" (siapa yang mengikuti username) atau "following"
// (siapa yang diikuti username). Yang terbaru lebih dulu.
func (s *Service) List(viewer *models.User, username, kind string, limit, offset int) (*ListResponse, error) {
	u, err := s.target(username)
	if err != nil {
		return nil, err
	}

	// otherCol = kolom yang menunjuk user pada daftar; selfCol = kolom milik u.
	var otherCol, selfCol string
	switch kind {
	case "followers":
		otherCol, selfCol = "follower_id", "following_id"
	case "following":
		otherCol, selfCol = "following_id", "follower_id"
	default:
		return nil, newErr(http.StatusNotFound, "daftar tidak ditemukan")
	}

	var total int64
	if err := s.DB.Model(&models.Follow{}).Where(selfCol+" = ?", u.ID).Count(&total).Error; err != nil {
		return nil, internalErr("count list", err)
	}

	var rows []cardRow
	err = s.DB.Table("follows AS f").
		Select("us.id AS id, us.username AS username, us.name AS name, us.xp AS xp").
		Joins("JOIN users us ON us.id = f."+otherCol).
		Where("f."+selfCol+" = ? AND us.role <> ?", u.ID, models.RoleAdmin).
		Order("f.created_at DESC, us.id DESC").
		Limit(limit).Offset(offset).
		Scan(&rows).Error
	if err != nil {
		return nil, internalErr("list follows", err)
	}

	ids := make([]uint, 0, len(rows))
	names := make(map[uint]string, len(rows))
	for _, r := range rows {
		ids = append(ids, r.ID)
		names[r.ID] = r.Username
	}
	avatars := media.AvatarURLs(s.DB, names)

	// Siapa saja dari daftar ini yang diikuti viewer (satu query).
	followed := map[uint]bool{}
	if len(ids) > 0 {
		var fl []uint
		err := s.DB.Model(&models.Follow{}).
			Where("follower_id = ? AND following_id IN ?", viewer.ID, ids).
			Pluck("following_id", &fl).Error
		if err != nil {
			return nil, internalErr("viewer follows", err)
		}
		for _, id := range fl {
			followed[id] = true
		}
	}

	cards := make([]Card, 0, len(rows))
	for _, r := range rows {
		cards = append(cards, Card{
			Username: r.Username, Name: r.Name, XP: r.XP,
			AvatarURL: avatars[r.ID], IsFollowing: followed[r.ID], IsMe: r.ID == viewer.ID,
		})
	}
	return &ListResponse{Users: cards, Total: total, Limit: limit, Offset: offset}, nil
}
