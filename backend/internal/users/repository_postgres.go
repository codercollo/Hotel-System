package users

import (
	"context"
	"errors"
	"time"

	"github.com/codercollo/hotel-system/backend/pkg/apierror"
	"github.com/codercollo/hotel-system/backend/pkg/idgen"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type postgresRepository struct {
	db *pgxpool.Pool
}

// NewPostgresRepository creates a Postgres-backed users repository.
func NewPostgresRepository(db *pgxpool.Pool) Repository {
	return &postgresRepository{db: db}
}

const cols = `id, email, name, password, role, is_active, metadata, created_at, updated_at`

func scan(row pgx.Row) (*User, error) {
	u := &User{}
	err := row.Scan(
		&u.ID, &u.Email, &u.Name, &u.Password,
		&u.Role, &u.IsActive, &u.Metadata, &u.CreatedAt, &u.UpdatedAt,
	)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, apierror.NotFound(apierror.CodeUserNotFound, "user not found")
	}
	if err != nil {
		return nil, apierror.Internal(err)
	}
	return u, nil
}

func (r *postgresRepository) FindByID(ctx context.Context, id string) (*User, error) {
	row := r.db.QueryRow(ctx,
		`SELECT `+cols+` FROM users WHERE id=$1 AND deleted_at IS NULL`, id)
	return scan(row)
}

func (r *postgresRepository) FindByEmail(ctx context.Context, email string) (*User, error) {
	row := r.db.QueryRow(ctx,
		`SELECT `+cols+` FROM users WHERE email=$1 AND deleted_at IS NULL`, email)
	return scan(row)
}

func (r *postgresRepository) List(ctx context.Context, limit, offset int) ([]*User, int, error) {
	rows, err := r.db.Query(ctx,
		`SELECT `+cols+` FROM users WHERE deleted_at IS NULL ORDER BY created_at DESC LIMIT $1 OFFSET $2`,
		limit, offset,
	)
	if err != nil {
		return nil, 0, apierror.Internal(err)
	}
	defer rows.Close()

	var users []*User
	for rows.Next() {
		u := &User{}
		if err := rows.Scan(
			&u.ID, &u.Email, &u.Name, &u.Password,
			&u.Role, &u.IsActive, &u.Metadata, &u.CreatedAt, &u.UpdatedAt,
		); err != nil {
			return nil, 0, apierror.Internal(err)
		}
		users = append(users, u)
	}

	var total int
	_ = r.db.QueryRow(ctx, `SELECT COUNT(*) FROM users WHERE deleted_at IS NULL`).Scan(&total)
	return users, total, nil
}

func (r *postgresRepository) Create(ctx context.Context, input CreateUserInput) (*User, error) {
	id := idgen.NewULID()
	row := r.db.QueryRow(ctx,
		`INSERT INTO users (id, email, name, password, role, is_active, metadata)
		 VALUES ($1,$2,$3,$4,$5,true,'{}')
		 RETURNING `+cols,
		id, input.Email, input.Name, input.Password, input.Role,
	)
	return scan(row)
}

func (r *postgresRepository) Update(ctx context.Context, id string, input UpdateUserInput) (*User, error) {
	user, err := r.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if input.Name != nil {
		user.Name = *input.Name
	}
	if input.Email != nil {
		user.Email = *input.Email
	}
	if input.IsActive != nil {
		user.IsActive = *input.IsActive
	}
	if input.Metadata != nil {
		user.Metadata = input.Metadata
	}

	row := r.db.QueryRow(ctx,
		`UPDATE users SET name=$1, email=$2, is_active=$3, metadata=$4, updated_at=$5
		 WHERE id=$6 AND deleted_at IS NULL RETURNING `+cols,
		user.Name, user.Email, user.IsActive, user.Metadata, time.Now(), id,
	)
	return scan(row)
}

func (r *postgresRepository) Delete(ctx context.Context, id string) error {
	cmd, err := r.db.Exec(ctx,
		`UPDATE users SET deleted_at=$1 WHERE id=$2 AND deleted_at IS NULL`,
		time.Now(), id,
	)
	if err != nil {
		return apierror.Internal(err)
	}
	if cmd.RowsAffected() == 0 {
		return apierror.NotFound(apierror.CodeUserNotFound, "user not found")
	}
	return nil
}
