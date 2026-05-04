package handlers

import (
	"context"

	"github.com/codercollo/hotel-system/backend/events"
	"github.com/codercollo/hotel-system/backend/pkg/logger"
)

// SendPush is a placeholder for FCM/APNs push dispatch. Phase 6 wires the real adapter.
func SendPush() events.Handler {
	return func(_ context.Context, ev events.Event) error {
		logger.Log().Info().Str("topic", string(ev.Topic)).Msg("push handler: stub (Phase 6)")
		return nil
	}
}
