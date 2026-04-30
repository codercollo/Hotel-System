CREATE TABLE IF NOT EXISTS uploads (
    id           TEXT        PRIMARY KEY,
    user_id      TEXT        REFERENCES users(id) ON DELETE SET NULL,
    filename     TEXT        NOT NULL,
    original_name TEXT       NOT NULL,
    mime_type    TEXT        NOT NULL,
    size         BIGINT      NOT NULL DEFAULT 0,  -- bytes
    provider     TEXT        NOT NULL DEFAULT 'local', -- local | s3
    path         TEXT        NOT NULL,            -- storage path or S3 key
    url          TEXT        NOT NULL DEFAULT '', -- public URL if applicable
    metadata     JSONB       NOT NULL DEFAULT '{}',
    created_at   TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_uploads_user_id ON uploads (user_id);
CREATE INDEX IF NOT EXISTS idx_uploads_provider ON uploads (provider);