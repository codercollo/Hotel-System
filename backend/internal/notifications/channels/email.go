package channels

import (
	"context"
	"fmt"
	"net/smtp"
)

// EmailChannel delivers notifications via SMTP.
type EmailChannel struct {
	host     string
	port     string
	user     string
	password string
	from     string
}

// NewEmailChannel constructs an SMTP email channel.
func NewEmailChannel(host, port, user, password, from string) *EmailChannel {
	return &EmailChannel{host: host, port: port, user: user, password: password, from: from}
}

func (c *EmailChannel) Name() string { return "email" }

// Send delivers an email via SMTP. In Phase 6 this will be wired to SendGrid/SES.
func (c *EmailChannel) Send(_ context.Context, msg Message) error {
	addr := fmt.Sprintf("%s:%s", c.host, c.port)
	body := fmt.Sprintf(
		"From: %s\r\nTo: %s\r\nSubject: %s\r\n\r\n%s",
		c.from, msg.To, msg.Subject, msg.Body,
	)

	var auth smtp.Auth
	if c.user != "" {
		auth = smtp.PlainAuth("", c.user, c.password, c.host)
	}

	return smtp.SendMail(addr, auth, c.from, []string{msg.To}, []byte(body))
}
