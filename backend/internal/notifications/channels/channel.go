// Package channels defines the notification delivery interface.
// Each channel (email, SMS, push) implements Channel.
package channels

import "context"

// Message is the content to be delivered.
type Message struct {
	To      string // email address, phone number, or device token
	Subject string // used by email
	Body    string
}

// Channel is the interface every notification adapter must satisfy.
type Channel interface {
	// Name returns the channel identifier: "email", "sms", "push".
	Name() string
	// Send delivers a message via this channel.
	Send(ctx context.Context, msg Message) error
}
