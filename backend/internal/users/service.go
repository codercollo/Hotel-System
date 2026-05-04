package users

import (
	"context"

	"github.com/codercollo/hotel-system/backend/internal/auth"
	"github.com/codercollo/hotel-system/backend/pkg/pagination"
)

// Service implements user management business logic.
type Service struct {
	repo Repository
}

// NewService constructs a users Service.
func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Create(ctx context.Context, name, email, passwordHash, role string) (auth.AuthUser, error) {
	u, err := s.repo.Create(ctx, CreateUserInput{
		Name:     name,
		Email:    email,
		Password: passwordHash,
		Role:     role,
	})
	if err != nil {
		return auth.AuthUser{}, err
	}
	return auth.AuthUser{
		ID:           u.ID,
		Email:        u.Email,
		Role:         u.Role,
		PasswordHash: u.Password,
		IsActive:     u.IsActive,
	}, nil
}

func (s *Service) Get(ctx context.Context, id string) (*User, error) {
	return s.repo.FindByID(ctx, id)
}

func (s *Service) List(ctx context.Context, p pagination.Params) ([]*User, int, error) {
	return s.repo.List(ctx, p.Limit, p.Offset)
}

func (s *Service) Update(ctx context.Context, id string, input UpdateUserInput) (*User, error) {
	return s.repo.Update(ctx, id, input)
}

func (s *Service) Delete(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}
