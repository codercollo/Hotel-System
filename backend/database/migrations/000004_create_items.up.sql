CREATE TABLE IF NOT EXISTS items (
    id          TEXT        PRIMARY KEY,
    name        TEXT        NOT NULL,
    description TEXT        NOT NULL DEFAULT '',
    price       BIGINT      NOT NULL DEFAULT 0,      -- stored in smallest currency unit
    currency    TEXT        NOT NULL DEFAULT 'USD',
    status      TEXT        NOT NULL DEFAULT 'active', -- active | inactive | archived
    stock       INTEGER     NOT NULL DEFAULT 0,
    images      TEXT[]      NOT NULL DEFAULT '{}',
    metadata    JSONB       NOT NULL DEFAULT '{}',
    search_vec  TSVECTOR,
    created_by  TEXT        REFERENCES users(id),
    created_at  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at  TIMESTAMPTZ
);

CREATE INDEX IF NOT EXISTS idx_items_status     ON items (status);
CREATE INDEX IF NOT EXISTS idx_items_deleted_at ON items (deleted_at) WHERE deleted_at IS NULL;
CREATE INDEX IF NOT EXISTS idx_items_search     ON items USING GIN (search_vec);
CREATE INDEX IF NOT EXISTS idx_items_metadata   ON items USING GIN (metadata);

-- Populate search vector from name and description
CREATE OR REPLACE FUNCTION items_search_vector_update()
RETURNS TRIGGER AS $$
BEGIN
    NEW.search_vec := to_tsvector('english',
        COALESCE(NEW.name, '') || ' ' || COALESCE(NEW.description, '')
    );
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER items_search_vec
    BEFORE INSERT OR UPDATE ON items
    FOR EACH ROW EXECUTE FUNCTION items_search_vector_update();

CREATE TRIGGER items_updated_at
    BEFORE UPDATE ON items
    FOR EACH ROW EXECUTE FUNCTION update_updated_at();