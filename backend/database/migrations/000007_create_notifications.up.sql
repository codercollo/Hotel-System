CREATE TABLE IF NOT EXISTS notifications (
    id         TEXT        PRIMARY KEY,
    user_id    TEXT        NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    type       TEXT        NOT NULL,           -- order_confirmed | payment_completed | system | etc.
    channel    TEXT        NOT NULL DEFAULT 'in_app', -- in_app | email | sms | push
    title      TEXT        NOT NULL DEFAULT '',
    body       TEXT        NOT NULL DEFAULT '',
    is_read    BOOLEAN     NOT NULL DEFAULT FALSE,
    metadata   JSONB       NOT NULL DEFAULT '{}',
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    read_at    TIMESTAMPTZ
);

CREATE INDEX IF NOT EXISTS idx_notifications_user_id ON notifications (user_id);
CREATE INDEX IF NOT EXISTS idx_notifications_is_read ON notifications (user_id, is_read);
CREATE INDEX IF NOT EXISTS idx_notifications_type    ON notifications (type);