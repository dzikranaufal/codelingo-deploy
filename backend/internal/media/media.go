// Package media menyimpan dan menyajikan foto profil dan banner user.
// Gambar disimpan di tabel user_media (PostgreSQL/Supabase) sebagai bytea.
// Semua validasi (ukuran, tipe, isi gambar) dilakukan di sini, bukan di frontend.
package media

import (
	"bytes"
	"errors"
	"fmt"
	"image"
	_ "image/gif"  // decoder untuk image.DecodeConfig
	_ "image/jpeg" // decoder untuk image.DecodeConfig
	_ "image/png"  // decoder untuk image.DecodeConfig
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"codelingo/backend/internal/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

const (
	KindAvatar = "avatar"
	KindBanner = "banner"

	MaxAvatarBytes = 2 << 20 // 2 MiB
	MaxBannerBytes = 4 << 20 // 4 MiB
	maxDimension   = 6000    // piksel, sisi terpanjang

	// PathPrefix = awalan URL publik: /api/v1/media/users/{username}/{kind}
	PathPrefix = "/api/v1/media/users/"
)

var (
	ErrInvalidKind  = errors.New("jenis gambar tidak valid")
	ErrTooLarge     = errors.New("ukuran gambar terlalu besar")
	ErrUnsupported  = errors.New("format gambar harus JPEG, PNG, GIF, atau WebP")
	ErrInvalidImage = errors.New("file bukan gambar yang valid")
	ErrInternal     = errors.New("terjadi kesalahan internal")
)

func ValidKind(kind string) bool { return kind == KindAvatar || kind == KindBanner }

// MaxBytes = batas ukuran file per jenis (0 jika jenis tidak valid).
func MaxBytes(kind string) int {
	switch kind {
	case KindAvatar:
		return MaxAvatarBytes
	case KindBanner:
		return MaxBannerBytes
	}
	return 0
}

// Validate memeriksa isi file (bukan nama atau header dari klien) dan
// mengembalikan content-type yang aman. SVG dan tipe lain ditolak.
func Validate(data []byte) (string, error) {
	if len(data) == 0 {
		return "", ErrInvalidImage
	}
	ct := http.DetectContentType(data)
	switch ct {
	case "image/jpeg", "image/png", "image/gif":
		cfg, _, err := image.DecodeConfig(bytes.NewReader(data))
		if err != nil {
			return "", ErrInvalidImage
		}
		if cfg.Width < 1 || cfg.Height < 1 || cfg.Width > maxDimension || cfg.Height > maxDimension {
			return "", ErrInvalidImage
		}
	case "image/webp":
		// Tidak ada decoder WebP di stdlib; tanda tangan RIFF/WEBP sudah dicek
		// oleh DetectContentType.
	default:
		return "", ErrUnsupported
	}
	return ct, nil
}

// URL membentuk alamat publik gambar. v = penanda versi agar cache browser
// otomatis berganti setelah gambar diganti.
func URL(username, kind string, v time.Time) string {
	return fmt.Sprintf("%s%s/%s?v=%d", PathPrefix, url.PathEscape(username), kind, v.UnixMilli())
}

type Service struct {
	DB *gorm.DB
}

// Save mengganti (atau membuat) gambar user. Mengembalikan waktu versi baru.
func (s *Service) Save(userID uint, kind string, data []byte) (time.Time, error) {
	if !ValidKind(kind) {
		return time.Time{}, ErrInvalidKind
	}
	if len(data) > MaxBytes(kind) {
		return time.Time{}, ErrTooLarge
	}
	ct, err := Validate(data)
	if err != nil {
		return time.Time{}, err
	}
	now := time.Now()
	row := models.UserMedia{
		UserID: userID, Kind: kind, ContentType: ct,
		Data: data, SizeBytes: len(data), UpdatedAt: now,
	}
	err = s.DB.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "user_id"}, {Name: "kind"}},
		DoUpdates: clause.AssignmentColumns([]string{"content_type", "data", "size_bytes", "updated_at"}),
	}).Omit(clause.Associations).Create(&row).Error
	if err != nil {
		log.Printf("media: gagal menyimpan %s: %v", kind, err)
		return time.Time{}, ErrInternal
	}
	return now, nil
}

// Delete menghapus gambar user (tidak error jika memang belum ada).
func (s *Service) Delete(userID uint, kind string) error {
	if !ValidKind(kind) {
		return ErrInvalidKind
	}
	if err := s.DB.Where("user_id = ? AND kind = ?", userID, kind).Delete(&models.UserMedia{}).Error; err != nil {
		log.Printf("media: gagal menghapus %s: %v", kind, err)
		return ErrInternal
	}
	return nil
}

// URLs mengembalikan alamat foto profil dan banner satu user ("" jika belum ada).
func (s *Service) URLs(userID uint, username string) (avatar, banner string) {
	var rows []models.UserMedia
	err := s.DB.Select("user_id", "kind", "updated_at").Where("user_id = ?", userID).Find(&rows).Error
	if err != nil {
		log.Printf("media: gagal membaca daftar gambar: %v", err)
		return "", ""
	}
	for _, m := range rows {
		switch m.Kind {
		case KindAvatar:
			avatar = URL(username, KindAvatar, m.UpdatedAt)
		case KindBanner:
			banner = URL(username, KindBanner, m.UpdatedAt)
		}
	}
	return avatar, banner
}

// AvatarURLs = versi massal untuk daftar user (id -> alamat foto). User tanpa foto tidak ada di map.
func AvatarURLs(db *gorm.DB, usernames map[uint]string) map[uint]string {
	out := make(map[uint]string, len(usernames))
	if len(usernames) == 0 {
		return out
	}
	ids := make([]uint, 0, len(usernames))
	for id := range usernames {
		ids = append(ids, id)
	}
	var rows []models.UserMedia
	err := db.Select("user_id", "kind", "updated_at").
		Where("user_id IN ? AND kind = ?", ids, KindAvatar).Find(&rows).Error
	if err != nil {
		log.Printf("media: gagal membaca foto massal: %v", err)
		return out
	}
	for _, m := range rows {
		out[m.UserID] = URL(usernames[m.UserID], KindAvatar, m.UpdatedAt)
	}
	return out
}

// Serve menangani GET /api/v1/media/users/{username}/{avatar|banner}.
// Publik (tag <img> tidak bisa mengirim header Authorization); isinya hanya
// foto profil dan banner, bukan data akun.
func (s *Service) Serve(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet && r.Method != http.MethodHead {
		w.Header().Set("Allow", "GET, HEAD")
		http.Error(w, "method tidak diizinkan", http.StatusMethodNotAllowed)
		return
	}
	parts := strings.Split(strings.Trim(strings.TrimPrefix(r.URL.Path, PathPrefix), "/"), "/")
	if len(parts) != 2 || parts[0] == "" || !ValidKind(parts[1]) {
		http.NotFound(w, r)
		return
	}

	var u models.User
	if err := s.DB.Select("id").Where("username = ?", parts[0]).Take(&u).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			http.NotFound(w, r)
			return
		}
		log.Printf("media: gagal mencari user: %v", err)
		http.Error(w, "terjadi kesalahan internal", http.StatusInternalServerError)
		return
	}

	var m models.UserMedia
	if err := s.DB.Where("user_id = ? AND kind = ?", u.ID, parts[1]).Take(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			http.NotFound(w, r)
			return
		}
		log.Printf("media: gagal membaca gambar: %v", err)
		http.Error(w, "terjadi kesalahan internal", http.StatusInternalServerError)
		return
	}

	etag := `"` + strconv.FormatInt(m.UpdatedAt.UnixMilli(), 10) + `"`
	h := w.Header()
	h.Set("Content-Type", m.ContentType)
	h.Set("X-Content-Type-Options", "nosniff")
	h.Set("Content-Security-Policy", "default-src 'none'; sandbox")
	h.Set("Cross-Origin-Resource-Policy", "cross-origin")
	h.Set("Cache-Control", "public, max-age=3600")
	h.Set("ETag", etag)
	if r.Header.Get("If-None-Match") == etag {
		w.WriteHeader(http.StatusNotModified)
		return
	}
	h.Set("Content-Length", strconv.Itoa(len(m.Data)))
	w.WriteHeader(http.StatusOK)
	if r.Method == http.MethodGet {
		_, _ = w.Write(m.Data)
	}
}
