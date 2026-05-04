package handlers

import (
	"context"

	"github.com/codercollo/hotel-system/backend/events"
	"github.com/codercollo/hotel-system/backend/pkg/logger"
)

// AuditLog writes every event to the structured log for audit purposes.
// In Phase 12 this can be redirected to a dedicated audit database table.
func AuditLog() events.Handler {
	return func(_ context.Context, ev events.Event) error {
		logger.Log().Info().
			Str("event_id", ev.ID).
			Str("topic", string(ev.Topic)).
			Time("occurred_at", ev.OccuredAt).
			RawJSON("payload", ev.Payload).
			Msg("audit")
		return nil
	}
}
