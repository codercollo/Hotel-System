package handlers

import (
	"context"

	"github.com/codercollo/hotel-system/backend/events"
	"github.com/codercollo/hotel-system/backend/pkg/logger"
)

// SendSMS is a placeholder for SMS dispatch. Phase 6 wires Africa's Talking.
func SendSMS() events.Handler {
	return func(_ context.Context, ev events.Event) error {
		logger.Log().Info().Str("topic", string(ev.Topic)).Msg("sms handler: stub (Phase 6)")
		return nil
	}
}
