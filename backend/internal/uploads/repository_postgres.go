package uploads

import (
	"context"
	"errors"

	"github.com/codercollo/hotel-system/backend/pkg/apierror"
	"github.com/codercollo/hotel-system/backend/pkg/idgen"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type postgresRepository struct{ db *pgxpool.Pool }

// NewPostgresRepository creates a Postgres-backed uploads repository.
func NewPostgresRepository(db *pgxpool.Pool) Repository {
	return &postgresRepository{db: db}
}

func scanUpload(row pgx.Row) (*Upload, error) {
	u := &Upload{}
	err := row.Scan(&u.ID, &u.UserID, &u.Filename, &u.OriginalName,
		&u.MimeType, &u.Size, &u.Provider, &u.Path, &u.URL, &u.CreatedAt)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, apierror.NotFound(apierror.CodeUploadFailed, "upload not found")
	}
	if err != nil {
		return nil, apierror.Internal(err)
	}
	return u, nil
}

func (r *postgresRepository) Create(ctx context.Context, u *Upload) (*Upload, error) {
	u.ID = idgen.NewULID()
	return scanUpload(r.db.QueryRow(ctx,
		`INSERT INTO uploads (id, user_id, filename, original_name, mime_type, size, provider, path, url)
		 VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9)
		 RETURNING id, user_id, filename, original_name, mime_type, size, provider, path, url, created_at`,
		u.ID, u.UserID, u.Filename, u.OriginalName, u.MimeType,
		u.Size, u.Provider, u.Path, u.URL,
	))
}

func (r *postgresRepository) FindByID(ctx context.Context, id string) (*Upload, error) {
	return scanUpload(r.db.QueryRow(ctx,
		`SELECT id, user_id, filename, original_name, mime_type, size, provider, path, url, created_at
		 FROM uploads WHERE id=$1`, id,
	))
}

func (r *postgresRepository) ListForUser(ctx context.Context, userID string) ([]*Upload, error) {
	rows, err := r.db.Query(ctx,
		`SELECT id, user_id, filename, original_name, mime_type, size, provider, path, url, created_at
         FROM uploads WHERE user_id=$1 ORDER BY created_at DESC`, userID)
	if err != nil {
		return nil, apierror.Internal(err)
	}
	defer rows.Close()
	var result []*Upload
	for rows.Next() {
		u := &Upload{}
		if err := rows.Scan(&u.ID, &u.UserID, &u.Filename, &u.OriginalName,
			&u.MimeType, &u.Size, &u.Provider, &u.Path, &u.URL, &u.CreatedAt); err != nil {
			continue
		}
		result = append(result, u)
	}
	return result, nil
}

func (r *postgresRepository) Delete(ctx context.Context, id string) error {
	_, err := r.db.Exec(ctx, `DELETE FROM uploads WHERE id=$1`, id)
	return err
}
