package items

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/codercollo/hotel-system/backend/pkg/apierror"
	"github.com/codercollo/hotel-system/backend/pkg/idgen"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type postgresRepository struct {
	db *pgxpool.Pool
}

// NewPostgresRepository creates a Postgres-backed items repository.
func NewPostgresRepository(db *pgxpool.Pool) Repository {
	return &postgresRepository{db: db}
}

const cols = `id, name, description, price, currency, status, stock, images, metadata, created_by, created_at, updated_at`

func scanItem(row pgx.Row) (*Item, error) {
	item := &Item{}
	err := row.Scan(
		&item.ID, &item.Name, &item.Description, &item.Price, &item.Currency,
		&item.Status, &item.Stock, &item.Images, &item.Metadata,
		&item.CreatedBy, &item.CreatedAt, &item.UpdatedAt,
	)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, apierror.NotFound(apierror.CodeItemNotFound, "item not found")
	}
	if err != nil {
		return nil, apierror.Internal(err)
	}
	return item, nil
}

func (r *postgresRepository) FindByID(ctx context.Context, id string) (*Item, error) {
	row := r.db.QueryRow(ctx,
		`SELECT `+cols+` FROM items WHERE id=$1 AND deleted_at IS NULL`, id)
	return scanItem(row)
}

func (r *postgresRepository) List(ctx context.Context, f ListFilters) ([]*Item, int, error) {
	query := `SELECT ` + cols + ` FROM items WHERE deleted_at IS NULL`
	args := []any{}

	if f.Status != "" {
		args = append(args, f.Status)
		query += ` AND status=$1`
	}
	query += ` ORDER BY created_at DESC LIMIT $` + argN(len(args)+1) + ` OFFSET $` + argN(len(args)+2)
	args = append(args, f.Limit, f.Offset)

	rows, err := r.db.Query(ctx, query, args...)
	if err != nil {
		return nil, 0, apierror.Internal(err)
	}
	defer rows.Close()

	var result []*Item
	for rows.Next() {
		item := &Item{}
		if err := rows.Scan(
			&item.ID, &item.Name, &item.Description, &item.Price, &item.Currency,
			&item.Status, &item.Stock, &item.Images, &item.Metadata,
			&item.CreatedBy, &item.CreatedAt, &item.UpdatedAt,
		); err != nil {
			return nil, 0, apierror.Internal(err)
		}
		result = append(result, item)
	}

	var total int
	_ = r.db.QueryRow(ctx, `SELECT COUNT(*) FROM items WHERE deleted_at IS NULL`).Scan(&total)
	return result, total, nil
}

func (r *postgresRepository) Create(ctx context.Context, input CreateItemInput) (*Item, error) {
	id := idgen.NewULID()
	cur := "USD"
	if input.Currency != "" {
		cur = input.Currency
	}
	if input.Images == nil {
		input.Images = []string{}
	}
	if input.Metadata == nil {
		input.Metadata = map[string]any{}
	}

	row := r.db.QueryRow(ctx,
		`INSERT INTO items (id, name, description, price, currency, stock, images, metadata, created_by)
		 VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9)
		 RETURNING `+cols,
		id, input.Name, input.Description, input.Price, cur,
		input.Stock, input.Images, input.Metadata, input.CreatedBy,
	)
	return scanItem(row)
}

func (r *postgresRepository) Update(ctx context.Context, id string, input UpdateItemInput) (*Item, error) {
	item, err := r.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if input.Name != nil {
		item.Name = *input.Name
	}
	if input.Description != nil {
		item.Description = *input.Description
	}
	if input.Price != nil {
		item.Price = *input.Price
	}
	if input.Status != nil {
		item.Status = *input.Status
	}
	if input.Stock != nil {
		item.Stock = *input.Stock
	}
	if input.Images != nil {
		item.Images = input.Images
	}
	if input.Metadata != nil {
		item.Metadata = input.Metadata
	}

	row := r.db.QueryRow(ctx,
		`UPDATE items SET name=$1, description=$2, price=$3, status=$4, stock=$5,
		 images=$6, metadata=$7, updated_at=$8
		 WHERE id=$9 AND deleted_at IS NULL RETURNING `+cols,
		item.Name, item.Description, item.Price, item.Status, item.Stock,
		item.Images, item.Metadata, time.Now(), id,
	)
	return scanItem(row)
}

func (r *postgresRepository) Delete(ctx context.Context, id string) error {
	cmd, err := r.db.Exec(ctx,
		`UPDATE items SET deleted_at=$1 WHERE id=$2 AND deleted_at IS NULL`,
		time.Now(), id,
	)
	if err != nil {
		return apierror.Internal(err)
	}
	if cmd.RowsAffected() == 0 {
		return apierror.NotFound(apierror.CodeItemNotFound, "item not found")
	}
	return nil
}

func (r *postgresRepository) Search(ctx context.Context, query string, limit, offset int) ([]*Item, int, error) {
	rows, err := r.db.Query(ctx,
		`SELECT `+cols+` FROM items
		 WHERE deleted_at IS NULL AND status='active'
		   AND (search_vec @@ plainto_tsquery('english', $1) OR name ILIKE '%' || $1 || '%')
		 ORDER BY ts_rank(search_vec, plainto_tsquery('english', $1)) DESC
		 LIMIT $2 OFFSET $3`,
		query, limit, offset,
	)
	if err != nil {
		return nil, 0, apierror.Internal(err)
	}
	defer rows.Close()

	var result []*Item
	for rows.Next() {
		item := &Item{}
		if err := rows.Scan(
			&item.ID, &item.Name, &item.Description, &item.Price, &item.Currency,
			&item.Status, &item.Stock, &item.Images, &item.Metadata,
			&item.CreatedBy, &item.CreatedAt, &item.UpdatedAt,
		); err != nil {
			continue
		}
		result = append(result, item)
	}

	var total int
	_ = r.db.QueryRow(ctx,
		`SELECT COUNT(*) FROM items WHERE deleted_at IS NULL AND status='active'
		 AND (search_vec @@ plainto_tsquery('english', $1) OR name ILIKE '%' || $1 || '%')`,
		query,
	).Scan(&total)

	return result, total, nil
}

func argN(n int) string {
	return fmt.Sprintf("%d", n)
}
