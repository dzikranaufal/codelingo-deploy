package follow

import (
	"errors"
	"net/http"
	"strconv"
	"strings"

	"codelingo/backend/internal/httpx"
	"codelingo/backend/internal/middleware"
)

const usersPrefix = "/api/v1/users/"

type Handler struct {
	Service *Service
}

func writeError(w http.ResponseWriter, err error) {
	var apiErr *APIError
	if errors.As(err, &apiErr) {
		httpx.WriteError(w, apiErr.Status, apiErr.Message)
		return
	}
	httpx.WriteError(w, http.StatusInternalServerError, "terjadi kesalahan internal")
}

// Routes menangani (semua wajib login):
//
//	GET    /api/v1/users/{username}            profil publik + jumlah followers/following
//	POST   /api/v1/users/{username}/follow     ikuti
//	DELETE /api/v1/users/{username}/follow     berhenti mengikuti
//	GET    /api/v1/users/{username}/followers  daftar pengikut  (?limit=&offset=)
//	GET    /api/v1/users/{username}/following  daftar yang diikuti
func (h *Handler) Routes(w http.ResponseWriter, r *http.Request) {
	viewer, ok := middleware.UserFromContext(r.Context())
	if !ok {
		httpx.WriteError(w, http.StatusUnauthorized, "token tidak valid")
		return
	}

	parts := strings.Split(strings.Trim(strings.TrimPrefix(r.URL.Path, usersPrefix), "/"), "/")
	if len(parts) == 0 || len(parts) > 2 || parts[0] == "" {
		httpx.WriteError(w, http.StatusNotFound, "alamat tidak ditemukan")
		return
	}
	username := parts[0]

	if len(parts) == 1 {
		if r.Method != http.MethodGet {
			httpx.MethodNotAllowed(w, http.MethodGet)
			return
		}
		resp, err := h.Service.Profile(viewer, username)
		if err != nil {
			writeError(w, err)
			return
		}
		httpx.WriteJSON(w, http.StatusOK, resp)
		return
	}

	switch parts[1] {
	case "follow":
		var (
			resp *PublicProfile
			err  error
		)
		switch r.Method {
		case http.MethodPost:
			resp, err = h.Service.Follow(viewer, username)
		case http.MethodDelete:
			resp, err = h.Service.Unfollow(viewer, username)
		default:
			httpx.MethodNotAllowed(w, http.MethodPost, http.MethodDelete)
			return
		}
		if err != nil {
			writeError(w, err)
			return
		}
		httpx.WriteJSON(w, http.StatusOK, resp)

	case "followers", "following":
		if r.Method != http.MethodGet {
			httpx.MethodNotAllowed(w, http.MethodGet)
			return
		}
		limit, offset := DefaultLimit, 0
		if raw := r.URL.Query().Get("limit"); raw != "" {
			n, err := strconv.Atoi(raw)
			if err != nil || n < 1 || n > MaxLimit {
				httpx.WriteError(w, http.StatusBadRequest, "limit harus angka 1 sampai 100")
				return
			}
			limit = n
		}
		if raw := r.URL.Query().Get("offset"); raw != "" {
			n, err := strconv.Atoi(raw)
			if err != nil || n < 0 {
				httpx.WriteError(w, http.StatusBadRequest, "offset harus angka 0 atau lebih")
				return
			}
			offset = n
		}
		resp, err := h.Service.List(viewer, username, parts[1], limit, offset)
		if err != nil {
			writeError(w, err)
			return
		}
		httpx.WriteJSON(w, http.StatusOK, resp)

	default:
		httpx.WriteError(w, http.StatusNotFound, "alamat tidak ditemukan")
	}
}
