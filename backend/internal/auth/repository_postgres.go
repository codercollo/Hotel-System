package auth

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/codercollo/hotel-system/backend/pkg/apierror"
)

type postgresRepo struct {
	db *pgxpool.Pool
}

// NewPostgresRepository returns a PostgreSQL-backed Repository.
func NewPostgresRepository(db *pgxpool.Pool) Repository {
	return &postgresRepo{db: db}
}

func (r *postgresRepo) CreateSession(ctx context.Context, s Session) error {
	const q = `
		INSERT INTO sessions
			(id, user_id, refresh_token, user_agent, client_ip, is_blocked, expires_at, created_at)
		VALUES ($1,$2,$3,$4,$5,$6,$7,$8)`
	_, err := r.db.Exec(ctx, q,
		s.ID, s.UserID, s.RefreshToken, s.UserAgent, s.ClientIP,
		s.IsBlocked, s.ExpiresAt, s.CreatedAt)
	return pgErr(err)
}

func (r *postgresRepo) GetSession(ctx context.Context, refreshToken string) (*Session, error) {
	const q = `
		SELECT id, user_id, refresh_token, user_agent, client_ip, is_blocked, expires_at, created_at
		FROM sessions WHERE refresh_token = $1`
	var s Session
	err := r.db.QueryRow(ctx, q, refreshToken).Scan(
		&s.ID, &s.UserID, &s.RefreshToken, &s.UserAgent, &s.ClientIP,
		&s.IsBlocked, &s.ExpiresAt, &s.CreatedAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, apierror.ErrNotFound
		}
		return nil, pgErr(err)
	}
	return &s, nil
}

func (r *postgresRepo) BlockSession(ctx context.Context, refreshToken string) error {
	const q = `UPDATE sessions SET is_blocked = true WHERE refresh_token = $1`
	_, err := r.db.Exec(ctx, q, refreshToken)
	return pgErr(err)
}

func (r *postgresRepo) DeleteExpiredSessions(ctx context.Context) error {
	const q = `DELETE FROM sessions WHERE expires_at < NOW()`
	_, err := r.db.Exec(ctx, q)
	return pgErr(err)
}

func pgErr(err error) error {
	if err == nil {
		return nil
	}
	return apierror.New(500, apierror.CodeInternal, err.Error())
}
