CREATE TABLE IF NOT EXISTS orders (
    id          TEXT        PRIMARY KEY,
    user_id     TEXT        NOT NULL REFERENCES users(id),
    status      TEXT        NOT NULL DEFAULT 'pending',
                            -- pending | confirmed | processing | completed | cancelled
    total       BIGINT      NOT NULL DEFAULT 0,
    currency    TEXT        NOT NULL DEFAULT 'USD',
    notes       TEXT        NOT NULL DEFAULT '',
    metadata    JSONB       NOT NULL DEFAULT '{}',
    created_at  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at  TIMESTAMPTZ
);

CREATE TABLE IF NOT EXISTS order_items (
    id         TEXT    PRIMARY KEY,
    order_id   TEXT    NOT NULL REFERENCES orders(id) ON DELETE CASCADE,
    item_id    TEXT    NOT NULL REFERENCES items(id),
    name       TEXT    NOT NULL,            -- snapshot at time of order
    price      BIGINT  NOT NULL,
    quantity   INTEGER NOT NULL DEFAULT 1,
    subtotal   BIGINT  NOT NULL,
    metadata   JSONB   NOT NULL DEFAULT '{}'
);

CREATE INDEX IF NOT EXISTS idx_orders_user_id    ON orders (user_id);
CREATE INDEX IF NOT EXISTS idx_orders_status     ON orders (status);
CREATE INDEX IF NOT EXISTS idx_orders_deleted_at ON orders (deleted_at) WHERE deleted_at IS NULL;
CREATE INDEX IF NOT EXISTS idx_order_items_order ON order_items (order_id);

CREATE TRIGGER orders_updated_at
    BEFORE UPDATE ON orders
    FOR EACH ROW EXECUTE FUNCTION update_updated_at();