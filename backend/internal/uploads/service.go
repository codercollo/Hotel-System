package uploads

import (
	"context"
	"fmt"
	"mime"
	"mime/multipart"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/codercollo/hotel-system/backend/internal/uploads/storage"
	"github.com/codercollo/hotel-system/backend/middleware"
	pkgmiddleware "github.com/codercollo/hotel-system/backend/middleware"
	"github.com/codercollo/hotel-system/backend/pkg/apierror"
	"github.com/codercollo/hotel-system/backend/pkg/idgen"
	"github.com/codercollo/hotel-system/backend/pkg/jwt"
	"github.com/codercollo/hotel-system/backend/pkg/response"
	"github.com/go-chi/chi/v5"
)

const maxUploadSize = 10 << 20 // 10 MB

// ─── Service ─────────────────────────────────────────────────────────────────

// Service handles file uploads.
type Service struct {
	repo        Repository
	storage     storage.Storage
	storageType string
}

// NewService creates an uploads Service.
func NewService(repo Repository, s storage.Storage, storageType string) *Service {
	return &Service{repo: repo, storage: s, storageType: storageType}
}

// Upload stores a multipart file and saves metadata.
func (s *Service) Upload(ctx context.Context, fh *multipart.FileHeader, userID *string) (*Upload, error) {
	f, err := fh.Open()
	if err != nil {
		return nil, apierror.BadRequest("failed to open uploaded file")
	}
	defer f.Close()

	ext := filepath.Ext(fh.Filename)
	key := fmt.Sprintf("uploads/%s%s", idgen.NewULID(), ext)
	mimeType := mime.TypeByExtension(ext)
	if mimeType == "" {
		mimeType = "application/octet-stream"
	}

	result, err := s.storage.Upload(ctx, key, f, fh.Size, mimeType)
	if err != nil {
		return nil, apierror.Internal(err)
	}

	return s.repo.Create(ctx, &Upload{
		UserID:       userID,
		Filename:     strings.TrimPrefix(key, "uploads/"),
		OriginalName: fh.Filename,
		MimeType:     mimeType,
		Size:         result.Size,
		Provider:     s.storageType,
		Path:         result.Key,
		URL:          result.URL,
	})
}

func (s *Service) ListForUser(ctx context.Context, userID string) ([]*Upload, error) {
	return s.repo.ListForUser(ctx, userID)
}

// Delete removes a file from storage and its metadata.
func (s *Service) Delete(ctx context.Context, id string) error {
	u, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return err
	}
	_ = s.storage.Delete(ctx, u.Path)
	return s.repo.Delete(ctx, id)
}

// ─── Handler ─────────────────────────────────────────────────────────────────

// Handler holds HTTP handlers for the uploads module.
type Handler struct{ svc *Service }

// NewHandler creates an uploads Handler.
func NewHandler(svc *Service) *Handler { return &Handler{svc: svc} }

// Upload godoc — POST /api/v1/uploads
func (h *Handler) Upload(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseMultipartForm(maxUploadSize); err != nil {
		response.Error(w, apierror.BadRequest("request too large or invalid multipart form"))
		return
	}

	_, fh, err := r.FormFile("file")
	if err != nil {
		response.Error(w, apierror.BadRequest("file field is required"))
		return
	}

	userID := middleware.GetUserID(r.Context())
	upload, err := h.svc.Upload(r.Context(), fh, &userID)
	if err != nil {
		response.Error(w, err)
		return
	}
	response.JSONCreated(w, upload)
}

// Delete godoc — DELETE /api/v1/uploads/{id}
func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if err := h.svc.Delete(r.Context(), id); err != nil {
		response.Error(w, err)
		return
	}
	response.NoContent(w)
}

// ─── Routes ──────────────────────────────────────────────────────────────────

// RegisterRoutes mounts upload endpoints. All require authentication.
func RegisterRoutes(r chi.Router, h *Handler, jwtManager *jwt.Manager) {
	r.Group(func(r chi.Router) {
		r.Use(pkgmiddleware.Authenticate(jwtManager))
		r.Post("/uploads", h.Upload)
		r.Delete("/uploads/{id}", h.Delete)
	})
}
