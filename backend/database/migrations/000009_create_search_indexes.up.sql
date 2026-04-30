-- Composite trigram index for fuzzy name search across items.
-- This complements the tsvector full-text index created in migration 004.
CREATE INDEX IF NOT EXISTS idx_items_name_trgm
    ON items USING GIN (name gin_trgm_ops)
    WHERE deleted_at IS NULL;

-- Trigram index on user email for admin lookups.
CREATE INDEX IF NOT EXISTS idx_users_email_trgm
    ON users USING GIN (email gin_trgm_ops)
    WHERE deleted_at IS NULL;