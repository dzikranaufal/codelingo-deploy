package media

import (
	"bytes"
	"errors"
	"image"
	"image/png"
	"testing"
	"time"
)

func pngBytes(t *testing.T, w, h int) []byte {
	t.Helper()
	var buf bytes.Buffer
	if err := png.Encode(&buf, image.NewRGBA(image.Rect(0, 0, w, h))); err != nil {
		t.Fatal(err)
	}
	return buf.Bytes()
}

func TestValidate(t *testing.T) {
	if ct, err := Validate(pngBytes(t, 4, 4)); err != nil || ct != "image/png" {
		t.Errorf("PNG valid: got %q, %v", ct, err)
	}
	if _, err := Validate(nil); !errors.Is(err, ErrInvalidImage) {
		t.Errorf("kosong: got %v", err)
	}
	if _, err := Validate([]byte("<svg onload=alert(1)></svg>")); !errors.Is(err, ErrUnsupported) {
		t.Errorf("SVG/teks harus ditolak: got %v", err)
	}
	// Header PNG rusak: terdeteksi PNG tetapi tidak bisa didekode.
	if _, err := Validate([]byte("\x89PNG\r\n\x1a\nrusak")); !errors.Is(err, ErrInvalidImage) {
		t.Errorf("PNG rusak: got %v", err)
	}
}

func TestMaxBytesAndKind(t *testing.T) {
	if MaxBytes(KindAvatar) != MaxAvatarBytes || MaxBytes(KindBanner) != MaxBannerBytes || MaxBytes("x") != 0 {
		t.Error("MaxBytes salah")
	}
	if !ValidKind(KindAvatar) || !ValidKind(KindBanner) || ValidKind("../etc") {
		t.Error("ValidKind salah")
	}
}

func TestURL(t *testing.T) {
	got := URL("budi", KindAvatar, time.UnixMilli(1700000000123))
	want := "/api/v1/media/users/budi/avatar?v=1700000000123"
	if got != want {
		t.Errorf("URL = %q, mau %q", got, want)
	}
}
