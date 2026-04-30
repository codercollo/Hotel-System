CREATE TABLE IF NOT EXISTS payments (
    id               TEXT        PRIMARY KEY,
    order_id         TEXT        NOT NULL REFERENCES orders(id),
    provider         TEXT        NOT NULL,          -- mpesa | stripe | flutterwave | paystack
    provider_ref     TEXT,                          -- external reference ID from provider
    amount           BIGINT      NOT NULL,
    currency         TEXT        NOT NULL DEFAULT 'USD',
    status           TEXT        NOT NULL DEFAULT 'pending',
                                                    -- pending | processing | completed | failed | refunded
    webhook_payload  JSONB       NOT NULL DEFAULT '{}',
    metadata         JSONB       NOT NULL DEFAULT '{}',
    created_at       TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at       TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_payments_order_id    ON payments (order_id);
CREATE INDEX IF NOT EXISTS idx_payments_status      ON payments (status);
CREATE INDEX IF NOT EXISTS idx_payments_provider    ON payments (provider);
CREATE UNIQUE INDEX IF NOT EXISTS idx_payments_provider_ref ON payments (provider, provider_ref)
    WHERE provider_ref IS NOT NULL;

CREATE TRIGGER payments_updated_at
    BEFORE UPDATE ON payments
    FOR EACH ROW EXECUTE FUNCTION update_updated_at();