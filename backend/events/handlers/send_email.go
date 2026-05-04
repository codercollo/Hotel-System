package handlers

import (
	"context"

	"github.com/codercollo/hotel-system/backend/events"
	"github.com/codercollo/hotel-system/backend/pkg/logger"
)

// SendEmail is a generic async email dispatch handler.
// Phase 6 populates this with the real email service call.
func SendEmail() events.Handler {
	return func(_ context.Context, ev events.Event) error {
		logger.Log().Info().
			Str("topic", string(ev.Topic)).
			Str("event_id", ev.ID).
			Msg("email handler: received event (Phase 6 wires SMTP)")
		return nil
	}
}
