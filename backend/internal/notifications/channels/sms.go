package channels

import (
	"context"
	"fmt"
)

// SMSChannel delivers notifications via SMS (Africa's Talking / Twilio stub).
// Full implementation in Phase 6.
type SMSChannel struct {
	apiKey   string
	username string
}

// NewSMSChannel constructs an SMS channel.
func NewSMSChannel(apiKey, username string) *SMSChannel {
	return &SMSChannel{apiKey: apiKey, username: username}
}

func (c *SMSChannel) Name() string { return "sms" }

// Send queues an SMS. Phase 6 calls the AT or Twilio REST API.
func (c *SMSChannel) Send(_ context.Context, msg Message) error {
	fmt.Printf("[SMS stub] To: %s | Body: %s\n", msg.To, msg.Body)
	return nil
}
