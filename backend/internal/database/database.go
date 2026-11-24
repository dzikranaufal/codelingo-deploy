package database

import (
	"fmt"

	"codelingo/backend/internal/config"
	"codelingo/backend/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Connect membuka koneksi ke PostgreSQL (Supabase atau PostgreSQL lokal).
// Database harus sudah ada; skema dibuat oleh Migrate.
func Connect(cfg config.Config) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: cfg.DSN(),
		// Wajib true untuk Transaction pooler Supabase (port 6543).
		PreferSimpleProtocol: cfg.DBSimpleProtocol,
	}), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	sqlDB.SetMaxOpenConns(cfg.DBMaxOpenConns)
	sqlDB.SetMaxIdleConns(cfg.DBMaxIdleConns)
	sqlDB.SetConnMaxLifetime(cfg.DBConnMaxLifetime)
	return db, nil
}

// Migrate membuat/menyesuaikan tabel dari struct model.
// AutoMigrate hanya menambah tabel, kolom, index, dan constraint;
// tidak menghapus kolom atau data yang sudah ada.
func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(models.All()...)
}

// EnableRLS mengaktifkan Row Level Security pada semua tabel model, tanpa policy.
// Di Supabase, tabel schema public bisa diakses lewat Data API (PostgREST)
// memakai anon key; tanpa RLS siapa pun yang memegang key itu dapat membaca
// tabel users (termasuk hash password). Backend Go terhubung sebagai role
// postgres yang melewati RLS, jadi tidak terpengaruh. Aman dijalankan berulang.
func EnableRLS(db *gorm.DB) error {
	for _, m := range models.All() {
		stmt := &gorm.Statement{DB: db}
		if err := stmt.Parse(m); err != nil {
			return fmt.Errorf("parse model: %w", err)
		}
		table := stmt.Schema.Table
		if err := db.Exec(`ALTER TABLE "` + table + `" ENABLE ROW LEVEL SECURITY`).Error; err != nil {
			return fmt.Errorf("aktifkan RLS pada %s: %w", table, err)
		}
	}
	return nil
}
