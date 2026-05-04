package handlers

import (
	"context"

	"github.com/codercollo/hotel-system/backend/events"
	"github.com/codercollo/hotel-system/backend/internal/notifications"
	"github.com/codercollo/hotel-system/backend/internal/payments"
	ws "github.com/codercollo/hotel-system/backend/internal/websocket"
	"github.com/codercollo/hotel-system/backend/pkg/logger"
)

func PaymentCompleted(notifSvc *notifications.Service) events.Handler {
	return func(ctx context.Context, ev events.Event) error {
		var payment payments.Payment
		if err := ev.Decode(&payment); err != nil {
			logger.Error(err, "handler: payment_completed: decode")
			return nil
		}
		_, err := notifSvc.Send(ctx, notifications.SendInput{
			UserID:   payment.OrderID,
			Type:     "payment_completed",
			Channel:  "in_app",
			Title:    "Payment Successful ✅",
			Body:     "Your payment has been processed successfully.",
			Metadata: map[string]any{"payment_id": payment.ID, "order_id": payment.OrderID},
		})
		if err != nil {
			logger.Error(err, "handler: payment_completed: create notification")
		}
		return nil
	}
}

func PaymentCompletedWithWS(notifSvc *notifications.Service, hub *ws.Hub) events.Handler {
	return func(ctx context.Context, ev events.Event) error {
		var payment payments.Payment
		if err := ev.Decode(&payment); err != nil {
			logger.Error(err, "handler: payment_completed: decode")
			return nil
		}
		n, err := notifSvc.Send(ctx, notifications.SendInput{
			UserID:   payment.OrderID,
			Type:     "payment_completed",
			Channel:  "in_app",
			Title:    "Payment Successful ✅",
			Body:     "Your payment has been processed successfully.",
			Metadata: map[string]any{"payment_id": payment.ID, "order_id": payment.OrderID},
		})
		if err != nil {
			logger.Error(err, "handler: payment_completed: notify")
		}
		if n != nil {
			msg, _ := ws.NewMessage(ws.TypeNotification, n)
			hub.SendToUser(payment.OrderID, msg)
		}
		statusMsg, _ := ws.NewMessage(ws.TypePaymentStatus, map[string]any{
			"payment_id": payment.ID,
			"order_id":   payment.OrderID,
			"status":     payment.Status,
		})
		hub.SendToUser(payment.OrderID, statusMsg)
		return nil
	}
}
