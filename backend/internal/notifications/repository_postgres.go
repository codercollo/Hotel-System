package notifications

import (
	"context"
	"time"

	"github.com/codercollo/hotel-system/backend/pkg/apierror"
	"github.com/codercollo/hotel-system/backend/pkg/idgen"
	"github.com/jackc/pgx/v5/pgxpool"
)

type postgresRepository struct{ db *pgxpool.Pool }

// NewPostgresRepository creates a Postgres-backed notifications repository.
func NewPostgresRepository(db *pgxpool.Pool) Repository {
	return &postgresRepository{db: db}
}

func (r *postgresRepository) Create(ctx context.Context, n *Notification) (*Notification, error) {
	n.ID = idgen.NewULID()
	if n.Metadata == nil {
		n.Metadata = map[string]any{}
	}
	_, err := r.db.Exec(ctx,
		`INSERT INTO notifications (id, user_id, type, channel, title, body, metadata)
		 VALUES ($1,$2,$3,$4,$5,$6,$7)`,
		n.ID, n.UserID, n.Type, n.Channel, n.Title, n.Body, n.Metadata,
	)
	if err != nil {
		return nil, apierror.Internal(err)
	}
	n.CreatedAt = time.Now()
	return n, nil
}

func (r *postgresRepository) ListForUser(ctx context.Context, userID string, limit, offset int) ([]*Notification, int, error) {
	rows, err := r.db.Query(ctx,
		`SELECT id, user_id, type, channel, title, body, is_read, metadata, created_at, read_at
		 FROM notifications WHERE user_id=$1 ORDER BY created_at DESC LIMIT $2 OFFSET $3`,
		userID, limit, offset,
	)
	if err != nil {
		return nil, 0, apierror.Internal(err)
	}
	defer rows.Close()

	var result []*Notification
	for rows.Next() {
		n := &Notification{}
		if err := rows.Scan(&n.ID, &n.UserID, &n.Type, &n.Channel,
			&n.Title, &n.Body, &n.IsRead, &n.Metadata, &n.CreatedAt, &n.ReadAt); err != nil {
			return nil, 0, apierror.Internal(err)
		}
		result = append(result, n)
	}

	var total int
	_ = r.db.QueryRow(ctx, `SELECT COUNT(*) FROM notifications WHERE user_id=$1`, userID).Scan(&total)
	return result, total, nil
}

func (r *postgresRepository) MarkRead(ctx context.Context, id, userID string) error {
	now := time.Now()
	_, err := r.db.Exec(ctx,
		`UPDATE notifications SET is_read=true, read_at=$1 WHERE id=$2 AND user_id=$3`,
		now, id, userID,
	)
	return err
}

func (r *postgresRepository) MarkAllRead(ctx context.Context, userID string) error {
	now := time.Now()
	_, err := r.db.Exec(ctx,
		`UPDATE notifications SET is_read=true, read_at=$1 WHERE user_id=$2 AND is_read=false`,
		now, userID,
	)
	return err
}
