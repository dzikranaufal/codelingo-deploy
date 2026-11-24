// create-admin membuat akun admin pertama (atau mempromosikan akun yang sudah
// ada). Idempoten: jika akunnya sudah admin, tidak ada yang diubah, dan
// password akun yang sudah ada TIDAK pernah ditimpa.
//
// Pakai variabel lingkungan sementara (jangan simpan password di .env):
//
//	$env:ADMIN_USERNAME = "admin"
//	$env:ADMIN_EMAIL    = "admin@example.com"
//	$env:ADMIN_PASSWORD = "minimal-8-karakter"
//	go run ./cmd/create-admin
package main

import (
	"errors"
	"log"
	"os"
	"strings"
	"unicode"

	"codelingo/backend/internal/config"
	"codelingo/backend/internal/database"
	"codelingo/backend/internal/models"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

const adminAge = 18 // kolom age wajib diisi; admin tidak punya umur bermakna

func main() {
	cfg := config.Load()
	if err := cfg.ApplyTimezone(); err != nil {
		log.Fatal(err)
	}

	username := strings.ToLower(strings.TrimSpace(os.Getenv("ADMIN_USERNAME")))
	email := strings.ToLower(strings.TrimSpace(os.Getenv("ADMIN_EMAIL")))
	name := strings.TrimSpace(os.Getenv("ADMIN_NAME"))
	password := os.Getenv("ADMIN_PASSWORD")
	if name == "" {
		name = username
	}

	// Aturan yang sama dengan pendaftaran biasa (internal/auth/service.go).
	if len(username) < 3 || len(username) > 30 {
		log.Fatal("ADMIN_USERNAME wajib 3-30 karakter")
	}
	for _, r := range username {
		if !unicode.IsLetter(r) && !unicode.IsDigit(r) && r != '_' {
			log.Fatal("ADMIN_USERNAME hanya boleh huruf, angka, dan underscore")
		}
	}
	if len(email) > 100 || !strings.Contains(email, "@") {
		log.Fatal("ADMIN_EMAIL tidak valid")
	}
	if len(name) > 100 {
		log.Fatal("ADMIN_NAME maksimal 100 karakter")
	}

	db, err := database.Connect(cfg)
	if err != nil {
		log.Fatalf("gagal terhubung ke database: %v", err)
	}

	err = db.Transaction(func(tx *gorm.DB) error {
		var found []models.User
		if err := tx.Where("username = ? OR email = ?", username, email).Find(&found).Error; err != nil {
			return err
		}

		switch len(found) {
		case 0:
			if len(password) < 8 || len(password) > 72 {
				return errors.New("ADMIN_PASSWORD wajib 8-72 karakter untuk membuat akun baru")
			}
			hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
			if err != nil {
				return err
			}
			user := models.User{
				Name: name, Username: username, Email: email, Age: adminAge,
				Password: string(hash), Role: models.RoleAdmin,
			}
			if err := tx.Create(&user).Error; err != nil {
				return err
			}
			log.Printf("admin dibuat: %s <%s>", username, email)

		case 1:
			u := found[0]
			if u.Username != username || u.Email != email {
				return errors.New("username atau email sudah dipakai akun lain yang berbeda; periksa ADMIN_USERNAME dan ADMIN_EMAIL")
			}
			if u.Role == models.RoleAdmin {
				log.Printf("%s sudah admin, tidak ada yang diubah", username)
				return nil
			}
			if err := tx.Model(&models.User{}).Where("id = ?", u.ID).Update("role", models.RoleAdmin).Error; err != nil {
				return err
			}
			log.Printf("akun %s dipromosikan menjadi admin (password tidak diubah)", username)

		default:
			return errors.New("username dan email cocok dengan dua akun berbeda; selesaikan secara manual")
		}
		return nil
	})
	if err != nil {
		log.Fatalf("create-admin gagal: %v", err)
	}
}
