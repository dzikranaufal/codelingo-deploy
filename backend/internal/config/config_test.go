package config

import (
	"net/url"
	"testing"
	"time"
)

func parse(t *testing.T, dsn string) *url.URL {
	t.Helper()
	u, err := url.Parse(dsn)
	if err != nil {
		t.Fatalf("DSN tidak bisa di-parse: %v", err)
	}
	return u
}

func TestDSNFromComponentsEscapesPassword(t *testing.T) {
	c := Config{
		DBHost: "db.example.com", DBPort: "5432", DBUser: "postgres.ref",
		DBPassword: "p@ss:w/rd#1", DBName: "postgres",
		DBSSLMode: "require", Timezone: "Asia/Jakarta",
	}

	u := parse(t, c.DSN())
	pw, _ := u.User.Password()

	if u.User.Username() != "postgres.ref" || pw != "p@ss:w/rd#1" {
		t.Fatalf("kredensial salah setelah di-parse: %q / %q", u.User.Username(), pw)
	}
	if u.Host != "db.example.com:5432" || u.Path != "/postgres" {
		t.Fatalf("host/path salah: %s%s", u.Host, u.Path)
	}
	if u.Query().Get("sslmode") != "require" {
		t.Fatalf("sslmode salah: %s", u.RawQuery)
	}
	if u.Query().Has("TimeZone") {
		t.Fatalf("DSN tidak boleh menambahkan TimeZone otomatis: %s", u.RawQuery)
	}
}

func TestDSNFromURLAddsMissingDefaults(t *testing.T) {
	c := Config{
		DatabaseURL: "postgresql://u:p@host:5432/postgres",
		DBSSLMode:   "require",
		Timezone:    "Asia/Jakarta",
	}

	q := parse(t, c.DSN()).Query()

	if q.Get("sslmode") != "require" {
		t.Fatalf("default sslmode tidak ditambahkan: %v", q)
	}
	if q.Has("TimeZone") {
		t.Fatalf("TimeZone tidak boleh ditambahkan otomatis: %v", q)
	}
}

func TestDSNFromURLKeepsExplicitParams(t *testing.T) {
	c := Config{
		DatabaseURL: "postgresql://u:p@host:5432/postgres?sslmode=disable&TimeZone=UTC",
		DBSSLMode:   "require",
		Timezone:    "Asia/Jakarta",
	}

	q := parse(t, c.DSN()).Query()

	if q.Get("sslmode") != "disable" || q.Get("TimeZone") != "UTC" {
		t.Fatalf("parameter eksplisit tertimpa: %v", q)
	}
}

func TestApplyTimezone(t *testing.T) {
	old := time.Local
	defer func() { time.Local = old }()

	if err := (Config{Timezone: "Asia/Jakarta"}).ApplyTimezone(); err != nil {
		t.Fatal(err)
	}
	if _, off := time.Date(2026, 1, 1, 0, 0, 0, 0, time.Local).Zone(); off != 7*3600 {
		t.Fatalf("offset = %d, ingin 25200", off)
	}
	if err := (Config{Timezone: "Bukan/Zona"}).ApplyTimezone(); err == nil {
		t.Fatal("zona tidak valid seharusnya error")
	}
}
