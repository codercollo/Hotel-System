package auth

import (
	"context"
	"net/http"
	"time"

	"github.com/codercollo/hotel-system/backend/pkg/apierror"
	pkgjwt "github.com/codercollo/hotel-system/backend/pkg/jwt"
	"github.com/codercollo/hotel-system/backend/pkg/password"
	"github.com/google/uuid"
)

// UserLookup is satisfied by users.Repository — kept here to avoid import cycles.
type UserLookup interface {
	FindByEmail(ctx context.Context, email string) (AuthUser, error)
}

// Service handles authentication business logic.
type Service struct {
	users  UserLookup
	repo   Repository
	jwtMgr *pkgjwt.Manager
}

// NewService constructs an auth Service.
// Signature: NewService(userRepo, tokenRepo, jwtManager) — matches main.go.
func NewService(users UserLookup, repo Repository, jwtMgr *pkgjwt.Manager) *Service {
	return &Service{users: users, repo: repo, jwtMgr: jwtMgr}
}

// Login validates credentials and issues a token pair.
func (s *Service) Login(ctx context.Context, inp LoginInput, userAgent, clientIP string) (*TokenPair, error) {
	user, err := s.users.FindByEmail(ctx, inp.Email)
	if err != nil {
		return nil, ErrInvalidCredentials
	}
	if !user.IsActive {
		return nil, ErrAccountInactive
	}
	if err := password.Compare(user.PasswordHash, inp.Password); err != nil {
		return nil, ErrInvalidCredentials
	}
	return s.issueTokenPair(ctx, user, userAgent, clientIP)
}

// Register hashes the password and hands off to the users service via the handler.
// The actual user row is created in users.Service; this method exists for symmetry.
func (s *Service) Register(_ context.Context, _ RegisterInput) error { return nil }

// Logout revokes the supplied refresh token.
func (s *Service) Logout(ctx context.Context, refreshToken string) error {
	return s.repo.BlockSession(ctx, refreshToken)
}

// Refresh validates a refresh token and issues a new token pair (token rotation).
func (s *Service) Refresh(ctx context.Context, refreshToken, userAgent, clientIP string) (*TokenPair, error) {
	session, err := s.repo.GetSession(ctx, refreshToken)
	if err != nil {
		return nil, apierror.ErrUnauthorized
	}
	if session.IsBlocked {
		return nil, ErrSessionBlocked
	}
	if time.Now().After(session.ExpiresAt) {
		return nil, ErrSessionExpired
	}

	claims, err := s.jwtMgr.ParseRefresh(refreshToken)
	if err != nil {
		return nil, apierror.ErrUnauthorized
	}

	// Rotate — block old session before issuing a new one.
	_ = s.repo.BlockSession(ctx, refreshToken)

	user := AuthUser{ID: claims.UserID, Email: claims.Email, Role: claims.Role}
	return s.issueTokenPair(ctx, user, userAgent, clientIP)
}

// VerifyToken parses and validates a raw access token string.
func (s *Service) VerifyToken(tokenStr string) (*pkgjwt.Claims, error) {
	return s.jwtMgr.Parse(tokenStr)
}

// IssueForUser is called by the handler after creating a new user via users.Service.
func (s *Service) IssueForUser(ctx context.Context, user AuthUser, userAgent, clientIP string) (*TokenPair, error) {
	return s.issueTokenPair(ctx, user, userAgent, clientIP)
}

// issueTokenPair mints access + refresh tokens and persists the session.
func (s *Service) issueTokenPair(ctx context.Context, user AuthUser, userAgent, clientIP string) (*TokenPair, error) {
	accessToken, err := s.jwtMgr.SignAccess(user.ID, user.Email, user.Role)
	if err != nil {
		return nil, apierror.ErrInternal
	}
	refreshToken, err := s.jwtMgr.SignRefresh(user.ID, user.Email, user.Role)
	if err != nil {
		return nil, apierror.ErrInternal
	}

	expiresAt := time.Now().Add(s.jwtMgr.RefreshTTL())

	if err := s.repo.CreateSession(ctx, Session{
		ID:           uuid.NewString(),
		UserID:       user.ID,
		RefreshToken: refreshToken,
		UserAgent:    userAgent,
		ClientIP:     clientIP,
		IsBlocked:    false,
		ExpiresAt:    expiresAt,
		CreatedAt:    time.Now(),
	}); err != nil {
		return nil, apierror.New(http.StatusInternalServerError, apierror.CodeInternal, "could not persist session")
	}

	return &TokenPair{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiresAt:    expiresAt,
	}, nil
}
