package search

import (
	"context"

	"github.com/codercollo/hotel-system/backend/pkg/apierror"
	"github.com/jackc/pgx/v5/pgxpool"
)

type postgresRepository struct{ db *pgxpool.Pool }

// NewPostgresRepository creates a PostgreSQL full-text search repository.
func NewPostgresRepository(db *pgxpool.Pool) Repository {
	return &postgresRepository{db: db}
}

// Search queries the items table using pg tsvector + trigram similarity.
func (r *postgresRepository) Search(ctx context.Context, q SearchQuery) ([]*SearchResult, int, error) {
	if q.Q == "" {
		return nil, 0, nil
	}

	rows, err := r.db.Query(ctx,
		`SELECT id, 'item' AS type, name, description, price, metadata,
		        ts_rank(search_vec, plainto_tsquery('english', $1)) AS score,
		        created_at
		 FROM items
		 WHERE deleted_at IS NULL AND status='active'
		   AND (search_vec @@ plainto_tsquery('english', $1) OR name ILIKE '%' || $1 || '%')
		 ORDER BY score DESC
		 LIMIT $2 OFFSET $3`,
		q.Q, q.Limit, q.Offset,
	)
	if err != nil {
		return nil, 0, apierror.Internal(err)
	}
	defer rows.Close()

	var results []*SearchResult
	for rows.Next() {
		r2 := &SearchResult{}
		if err := rows.Scan(&r2.ID, &r2.Type, &r2.Name, &r2.Description,
			&r2.Price, &r2.Metadata, &r2.Score, &r2.CreatedAt); err != nil {
			continue
		}
		results = append(results, r2)
	}

	var total int
	_ = r.db.QueryRow(ctx,
		`SELECT COUNT(*) FROM items WHERE deleted_at IS NULL AND status='active'
		 AND (search_vec @@ plainto_tsquery('english', $1) OR name ILIKE '%' || $1 || '%')`, q.Q,
	).Scan(&total)

	return results, total, nil
}
