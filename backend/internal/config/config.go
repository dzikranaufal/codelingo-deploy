package config

import (
	"fmt"
	"net"
	"net/url"
	"os"
	"strconv"
	"time"
	_ "time/tzdata" // zona waktu tertanam, tidak bergantung pada OS (mis. Windows)

	"github.com/joho/godotenv"
)

type Config struct {
	// DatabaseURL (opsional) = connection string lengkap, mis. dari Supabase
	// (Dashboard > Connect > Session pooler). Jika diisi, DB_HOST/DB_PORT/
	// DB_USER/DB_PASSWORD/DB_NAME diabaikan.
	DatabaseURL string

	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	DBSSLMode  string

	// Pengaturan pool. Nilai kecil disengaja: database Supabase dipakai
	// bersama beberapa developer dan jumlah koneksinya terbatas.
	DBMaxOpenConns    int
	DBMaxIdleConns    int
	DBConnMaxLifetime time.Duration

	// DBSimpleProtocol = true wajib bila memakai Transaction pooler (port 6543),
	// karena mode itu tidak mendukung prepared statement.
	DBSimpleProtocol bool

	// Timezone = zona waktu aplikasi. Dipakai untuk batas hari/pekan (misi,
	// quest, streak, leaderboard) dan sebagai TimeZone sesi database.
	Timezone string

	ServerPort string
	JWTSecret  string
}

func Load() Config {
	_ = godotenv.Load()

	return Config{
		DatabaseURL: os.Getenv("DATABASE_URL"),

		DBHost:     getEnv("DB_HOST", "127.0.0.1"),
		DBPort:     getEnv("DB_PORT", "5432"),
		DBUser:     getEnv("DB_USER", "postgres"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:     getEnv("DB_NAME", "postgres"),
		DBSSLMode:  getEnv("DB_SSLMODE", "require"),

		DBMaxOpenConns:    getEnvInt("DB_MAX_OPEN_CONNS", 5),
		DBMaxIdleConns:    getEnvInt("DB_MAX_IDLE_CONNS", 2),
		DBConnMaxLifetime: time.Duration(getEnvInt("DB_CONN_MAX_LIFETIME_MINUTES", 30)) * time.Minute,
		DBSimpleProtocol:  getEnv("DB_SIMPLE_PROTOCOL", "false") == "true",

		Timezone: getEnv("APP_TIMEZONE", "Asia/Jakarta"),

		ServerPort: getEnv("SERVER_PORT", "8080"),
		JWTSecret:  os.Getenv("JWT_SECRET"),
	}
}

// DSN menghasilkan connection string PostgreSQL. Dari DATABASE_URL jika ada,
// jika tidak dirakit dari DB_*. Password dengan karakter khusus aman karena
// di-escape oleh net/url. TimeZone dan sslmode dilengkapi bila belum ada.
func (c *Config) DSN() string {
	raw := c.DatabaseURL
	if raw == "" {
		u := url.URL{
			Scheme: "postgres",
			User:   url.UserPassword(c.DBUser, c.DBPassword),
			Host:   net.JoinHostPort(c.DBHost, c.DBPort),
			Path:   "/" + c.DBName,
		}
		raw = u.String()
	}
	return withDefaults(raw, c.DBSSLMode)
}

// withDefaults menambahkan sslmode dan TimeZone ke URL jika belum ada.
// Jika raw bukan URL yang valid, dikembalikan apa adanya.
func withDefaults(raw, sslmode string) string {
	u, err := url.Parse(raw)
	if err != nil || u.Scheme == "" {
		return raw
	}

	q := u.Query()

	if q.Get("sslmode") == "" && sslmode != "" {
		q.Set("sslmode", sslmode)
	}

	u.RawQuery = q.Encode()
	return u.String()
}

// ApplyTimezone menetapkan zona waktu aplikasi (time.Local). Panggil sekali di
// awal setiap program, sebelum ada query atau goroutine berjalan.
func (c Config) ApplyTimezone() error {
	loc, err := time.LoadLocation(c.Timezone)
	if err != nil {
		return fmt.Errorf("APP_TIMEZONE %q tidak valid: %w", c.Timezone, err)
	}
	time.Local = loc
	return nil
}

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}

func getEnvInt(key string, fallback int) int {
	if value := os.Getenv(key); value != "" {
		if n, err := strconv.Atoi(value); err == nil && n > 0 {
			return n
		}
	}
	return fallback
}
