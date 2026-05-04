package channels

import (
	"context"
	"fmt"
)

// PushChannel delivers push notifications via FCM / APNs.
// Full implementation in Phase 6.
type PushChannel struct {
	serverKey string
}

// NewPushChannel constructs a push notification channel.
func NewPushChannel(serverKey string) *PushChannel {
	return &PushChannel{serverKey: serverKey}
}

func (c *PushChannel) Name() string { return "push" }

// Send delivers a push notification. Phase 6 calls the FCM HTTP v1 API.
func (c *PushChannel) Send(_ context.Context, msg Message) error {
	fmt.Printf("[PUSH stub] To: %s | Subject: %s | Body: %s\n", msg.To, msg.Subject, msg.Body)
	return nil
}
