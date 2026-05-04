package payments

import (
	"context"
	"errors"
	"time"

	"github.com/codercollo/hotel-system/backend/pkg/apierror"
	"github.com/codercollo/hotel-system/backend/pkg/idgen"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type postgresRepository struct{ db *pgxpool.Pool }

// NewPostgresRepository creates a Postgres-backed payments repository.
func NewPostgresRepository(db *pgxpool.Pool) Repository {
	return &postgresRepository{db: db}
}

const cols = `id, order_id, provider, provider_ref, amount, currency, status,
              webhook_payload, metadata, created_at, updated_at`

func scan(row pgx.Row) (*Payment, error) {
	p := &Payment{}
	err := row.Scan(
		&p.ID, &p.OrderID, &p.Provider, &p.ProviderRef,
		&p.Amount, &p.Currency, &p.Status,
		&p.WebhookPayload, &p.Metadata,
		&p.CreatedAt, &p.UpdatedAt,
	)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, apierror.NotFound(apierror.CodePaymentFailed, "payment not found")
	}
	if err != nil {
		return nil, apierror.Internal(err)
	}
	return p, nil
}

func (r *postgresRepository) FindByID(ctx context.Context, id string) (*Payment, error) {
	return scan(r.db.QueryRow(ctx, `SELECT `+cols+` FROM payments WHERE id=$1`, id))
}

func (r *postgresRepository) FindByOrderID(ctx context.Context, orderID string) ([]*Payment, error) {
	rows, err := r.db.Query(ctx,
		`SELECT `+cols+` FROM payments WHERE order_id=$1 ORDER BY created_at DESC`, orderID)
	if err != nil {
		return nil, apierror.Internal(err)
	}
	defer rows.Close()

	var result []*Payment
	for rows.Next() {
		p := &Payment{}
		if err := rows.Scan(&p.ID, &p.OrderID, &p.Provider, &p.ProviderRef,
			&p.Amount, &p.Currency, &p.Status,
			&p.WebhookPayload, &p.Metadata,
			&p.CreatedAt, &p.UpdatedAt); err != nil {
			return nil, apierror.Internal(err)
		}
		result = append(result, p)
	}
	return result, nil
}

// FindByProviderRef looks up a payment by the provider's own transaction reference.
// This is the correct lookup used during webhook processing.
func (r *postgresRepository) FindByProviderRef(ctx context.Context, providerRef string) (*Payment, error) {
	return scan(r.db.QueryRow(ctx,
		`SELECT `+cols+` FROM payments WHERE provider_ref=$1 LIMIT 1`, providerRef))
}

func (r *postgresRepository) Create(ctx context.Context, p *Payment) (*Payment, error) {
	p.ID = idgen.NewULID()
	if p.Currency == "" {
		p.Currency = "USD"
	}
	if p.Metadata == nil {
		p.Metadata = map[string]any{}
	}
	if p.WebhookPayload == nil {
		p.WebhookPayload = map[string]any{}
	}
	row := r.db.QueryRow(ctx,
		`INSERT INTO payments (id, order_id, provider, amount, currency, status, webhook_payload, metadata)
		 VALUES ($1,$2,$3,$4,$5,$6,$7,$8) RETURNING `+cols,
		p.ID, p.OrderID, p.Provider, p.Amount, p.Currency,
		StatusPending, p.WebhookPayload, p.Metadata,
	)
	return scan(row)
}

func (r *postgresRepository) UpdateStatus(ctx context.Context, id, status string, providerRef *string) (*Payment, error) {
	row := r.db.QueryRow(ctx,
		`UPDATE payments
		 SET status=$1, provider_ref=COALESCE($2, provider_ref), updated_at=$3
		 WHERE id=$4
		 RETURNING `+cols,
		status, providerRef, time.Now(), id,
	)
	return scan(row)
}

func (r *postgresRepository) StoreWebhookPayload(ctx context.Context, id string, payload map[string]any) error {
	_, err := r.db.Exec(ctx,
		`UPDATE payments SET webhook_payload=$1, updated_at=$2 WHERE id=$3`,
		payload, time.Now(), id,
	)
	return err
}
