package orders

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

// NewPostgresRepository creates a Postgres-backed orders repository.
func NewPostgresRepository(db *pgxpool.Pool) Repository {
	return &postgresRepository{db: db}
}

func scanOrder(row pgx.Row) (*Order, error) {
	o := &Order{}
	err := row.Scan(
		&o.ID, &o.UserID, &o.Status, &o.Total, &o.Currency,
		&o.Notes, &o.Metadata, &o.CreatedAt, &o.UpdatedAt,
	)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, apierror.NotFound(apierror.CodeOrderNotFound, "order not found")
	}
	if err != nil {
		return nil, apierror.Internal(err)
	}
	return o, nil
}

func (r *postgresRepository) FindByID(ctx context.Context, id string) (*Order, error) {
	row := r.db.QueryRow(ctx,
		`SELECT id, user_id, status, total, currency, notes, metadata, created_at, updated_at
		 FROM orders WHERE id=$1 AND deleted_at IS NULL`, id)
	order, err := scanOrder(row)
	if err != nil {
		return nil, err
	}
	order.Items, _ = r.loadItems(ctx, id)
	return order, nil
}

func (r *postgresRepository) ListForUser(ctx context.Context, userID string, limit, offset int) ([]*Order, int, error) {
	return r.listOrders(ctx, &userID, limit, offset)
}

func (r *postgresRepository) ListAll(ctx context.Context, limit, offset int) ([]*Order, int, error) {
	return r.listOrders(ctx, nil, limit, offset)
}

func (r *postgresRepository) listOrders(ctx context.Context, userID *string, limit, offset int) ([]*Order, int, error) {
	var (
		rows interface {
			Close()
			Next() bool
			Err() error
			Scan(...any) error
		}
		err error
	)
	if userID != nil {
		rows, err = r.db.Query(ctx,
			`SELECT id, user_id, status, total, currency, notes, metadata, created_at, updated_at
			 FROM orders WHERE user_id=$1 AND deleted_at IS NULL
			 ORDER BY created_at DESC LIMIT $2 OFFSET $3`, *userID, limit, offset)
	} else {
		rows, err = r.db.Query(ctx,
			`SELECT id, user_id, status, total, currency, notes, metadata, created_at, updated_at
			 FROM orders WHERE deleted_at IS NULL ORDER BY created_at DESC LIMIT $1 OFFSET $2`, limit, offset)
	}
	if err != nil {
		return nil, 0, apierror.Internal(err)
	}
	defer rows.Close()

	var result []*Order
	for rows.Next() {
		o := &Order{}
		if err := rows.Scan(&o.ID, &o.UserID, &o.Status, &o.Total, &o.Currency,
			&o.Notes, &o.Metadata, &o.CreatedAt, &o.UpdatedAt); err != nil {
			return nil, 0, apierror.Internal(err)
		}
		result = append(result, o)
	}

	var total int
	if userID != nil {
		_ = r.db.QueryRow(ctx, `SELECT COUNT(*) FROM orders WHERE user_id=$1 AND deleted_at IS NULL`, *userID).Scan(&total)
	} else {
		_ = r.db.QueryRow(ctx, `SELECT COUNT(*) FROM orders WHERE deleted_at IS NULL`).Scan(&total)
	}

	return result, total, nil
}

func (r *postgresRepository) Create(ctx context.Context, order *Order, items []OrderItem) (*Order, error) {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return nil, apierror.Internal(err)
	}
	defer tx.Rollback(ctx)

	order.ID = idgen.NewULID()
	if order.Currency == "" {
		order.Currency = "USD"
	}
	if order.Metadata == nil {
		order.Metadata = map[string]any{}
	}

	_, err = tx.Exec(ctx,
		`INSERT INTO orders (id, user_id, status, total, currency, notes, metadata)
		 VALUES ($1,$2,$3,$4,$5,$6,$7)`,
		order.ID, order.UserID, StatusPending, order.Total,
		order.Currency, order.Notes, order.Metadata,
	)
	if err != nil {
		return nil, apierror.Internal(err)
	}

	for i := range items {
		items[i].ID = idgen.NewULID()
		items[i].OrderID = order.ID
		_, err = tx.Exec(ctx,
			`INSERT INTO order_items (id, order_id, item_id, name, price, quantity, subtotal)
			 VALUES ($1,$2,$3,$4,$5,$6,$7)`,
			items[i].ID, items[i].OrderID, items[i].ItemID,
			items[i].Name, items[i].Price, items[i].Quantity, items[i].Subtotal,
		)
		if err != nil {
			return nil, apierror.Internal(err)
		}
	}

	if err := tx.Commit(ctx); err != nil {
		return nil, apierror.Internal(err)
	}

	order.Status = StatusPending
	order.Items = items
	order.CreatedAt = time.Now()
	order.UpdatedAt = time.Now()
	return order, nil
}

func (r *postgresRepository) UpdateStatus(ctx context.Context, id, status string) (*Order, error) {
	row := r.db.QueryRow(ctx,
		`UPDATE orders SET status=$1, updated_at=$2
		 WHERE id=$3 AND deleted_at IS NULL
		 RETURNING id, user_id, status, total, currency, notes, metadata, created_at, updated_at`,
		status, time.Now(), id,
	)
	return scanOrder(row)
}

func (r *postgresRepository) loadItems(ctx context.Context, orderID string) ([]OrderItem, error) {
	rows, err := r.db.Query(ctx,
		`SELECT id, order_id, item_id, name, price, quantity, subtotal FROM order_items WHERE order_id=$1`,
		orderID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []OrderItem
	for rows.Next() {
		oi := OrderItem{}
		if err := rows.Scan(&oi.ID, &oi.OrderID, &oi.ItemID, &oi.Name, &oi.Price, &oi.Quantity, &oi.Subtotal); err != nil {
			continue
		}
		items = append(items, oi)
	}
	return items, nil
}
