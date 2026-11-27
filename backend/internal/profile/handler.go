package profile

import (
	"errors"
	"io"
	"net/http"
	"time"

	"codelingo/backend/internal/httpx"
	"codelingo/backend/internal/media"
	"codelingo/backend/internal/middleware"
)

type Handler struct {
	Service *Service
}

type updateInput struct {
	Name string `json:"name"`
}

type passwordInput struct {
	CurrentPassword string `json:"current_password"`
	NewPassword     string `json:"new_password"`
}

// Profile menangani GET dan PATCH /api/v1/profile (wajib dibungkus Authenticate).
func (h *Handler) Profile(w http.ResponseWriter, r *http.Request) {
	user, ok := middleware.UserFromContext(r.Context())
	if !ok {
		httpx.WriteError(w, http.StatusUnauthorized, "token tidak valid")
		return
	}

	switch r.Method {
	case http.MethodGet:
		httpx.WriteJSON(w, http.StatusOK, h.Service.Build(user))

	case http.MethodPatch:
		var input updateInput
		if err := httpx.DecodeJSON(w, r, &input); err != nil {
			httpx.WriteError(w, http.StatusBadRequest, "format JSON tidak valid")
			return
		}
		if err := h.Service.UpdateName(user, input.Name); err != nil {
			writeServiceError(w, err, "nama harus 2 sampai 100 karakter")
			return
		}
		httpx.WriteJSON(w, http.StatusOK, h.Service.Build(user))

	default:
		httpx.MethodNotAllowed(w, http.MethodGet, http.MethodPatch)
	}
}

// Password menangani PUT /api/v1/profile/password (wajib dibungkus Authenticate).
func (h *Handler) Password(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		httpx.MethodNotAllowed(w, http.MethodPut)
		return
	}

	user, ok := middleware.UserFromContext(r.Context())
	if !ok {
		httpx.WriteError(w, http.StatusUnauthorized, "token tidak valid")
		return
	}

	var input passwordInput
	if err := httpx.DecodeJSON(w, r, &input); err != nil {
		httpx.WriteError(w, http.StatusBadRequest, "format JSON tidak valid")
		return
	}

	if err := h.Service.ChangePassword(user, input.CurrentPassword, input.NewPassword); err != nil {
		writeServiceError(w, err, "password saat ini wajib diisi dan password baru harus 8 sampai 72 karakter")
		return
	}
	httpx.WriteJSON(w, http.StatusOK, map[string]string{"message": "password berhasil diubah"})
}

// Avatar menangani POST (unggah/ganti) dan DELETE /api/v1/profile/avatar.
func (h *Handler) Avatar(w http.ResponseWriter, r *http.Request) { h.mediaRoute(w, r, media.KindAvatar) }

// Banner menangani POST (unggah/ganti) dan DELETE /api/v1/profile/banner.
func (h *Handler) Banner(w http.ResponseWriter, r *http.Request) { h.mediaRoute(w, r, media.KindBanner) }

func (h *Handler) mediaRoute(w http.ResponseWriter, r *http.Request, kind string) {
	user, ok := middleware.UserFromContext(r.Context())
	if !ok {
		httpx.WriteError(w, http.StatusUnauthorized, "token tidak valid")
		return
	}
	svc := &media.Service{DB: h.Service.DB}

	switch r.Method {
	case http.MethodPost:
		// Server memakai timeout 10 detik; unggahan gambar butuh lebih lama.
		rc := http.NewResponseController(w)
		_ = rc.SetReadDeadline(time.Now().Add(60 * time.Second))
		_ = rc.SetWriteDeadline(time.Now().Add(60 * time.Second))

		limit := media.MaxBytes(kind)
		r.Body = http.MaxBytesReader(w, r.Body, int64(limit)+(1<<20)) // + ruang untuk header multipart

		data, err := readUpload(r, limit)
		if err != nil {
			writeMediaError(w, err)
			return
		}
		if _, err := svc.Save(user.ID, kind, data); err != nil {
			writeMediaError(w, err)
			return
		}
		httpx.WriteJSON(w, http.StatusOK, h.Service.Build(user))

	case http.MethodDelete:
		if err := svc.Delete(user.ID, kind); err != nil {
			writeMediaError(w, err)
			return
		}
		httpx.WriteJSON(w, http.StatusOK, h.Service.Build(user))

	default:
		httpx.MethodNotAllowed(w, http.MethodPost, http.MethodDelete)
	}
}

// readUpload membaca field multipart "file" secara streaming (tanpa file sementara)
// dan menolak file yang lebih besar dari limit.
func readUpload(r *http.Request, limit int) ([]byte, error) {
	mr, err := r.MultipartReader()
	if err != nil {
		return nil, media.ErrInvalidImage
	}
	for {
		part, err := mr.NextPart()
		if errors.Is(err, io.EOF) {
			return nil, media.ErrInvalidImage // field "file" tidak ada
		}
		if err != nil {
			return nil, mapReadErr(err)
		}
		if part.FormName() != "file" {
			_ = part.Close()
			continue
		}
		data, err := io.ReadAll(io.LimitReader(part, int64(limit)+1))
		_ = part.Close()
		if err != nil {
			return nil, mapReadErr(err)
		}
		if len(data) > limit {
			return nil, media.ErrTooLarge
		}
		return data, nil
	}
}

func mapReadErr(err error) error {
	var tooBig *http.MaxBytesError
	if errors.As(err, &tooBig) {
		return media.ErrTooLarge
	}
	return media.ErrInvalidImage
}

func writeMediaError(w http.ResponseWriter, err error) {
	switch {
	case errors.Is(err, media.ErrTooLarge):
		httpx.WriteError(w, http.StatusRequestEntityTooLarge, "ukuran gambar terlalu besar (foto maks. 2 MB, banner maks. 4 MB)")
	case errors.Is(err, media.ErrUnsupported):
		httpx.WriteError(w, http.StatusUnsupportedMediaType, "format gambar harus JPEG, PNG, GIF, atau WebP")
	case errors.Is(err, media.ErrInvalidImage), errors.Is(err, media.ErrInvalidKind):
		httpx.WriteError(w, http.StatusBadRequest, "file bukan gambar yang valid")
	default:
		httpx.WriteError(w, http.StatusInternalServerError, "terjadi kesalahan internal")
	}
}

// Password salah dibalas 400 (bukan 401) agar interceptor frontend tidak
// salah mengira sesi login sudah habis.
func writeServiceError(w http.ResponseWriter, err error, invalidMessage string) {
	switch {
	case errors.Is(err, ErrInvalidInput):
		httpx.WriteError(w, http.StatusBadRequest, invalidMessage)
	case errors.Is(err, ErrWrongPassword):
		httpx.WriteError(w, http.StatusBadRequest, "password saat ini salah")
	case errors.Is(err, ErrSamePassword):
		httpx.WriteError(w, http.StatusBadRequest, "password baru harus berbeda dari password saat ini")
	default:
		httpx.WriteError(w, http.StatusInternalServerError, "terjadi kesalahan internal")
	}
}