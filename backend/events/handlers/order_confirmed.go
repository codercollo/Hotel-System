package handlers

import (
	"context"
	"fmt"

	"github.com/codercollo/hotel-system/backend/events"
	"github.com/codercollo/hotel-system/backend/internal/notifications"
	"github.com/codercollo/hotel-system/backend/internal/orders"
	ws "github.com/codercollo/hotel-system/backend/internal/websocket"
	"github.com/codercollo/hotel-system/backend/pkg/logger"
)

func OrderConfirmed(notifSvc *notifications.Service) events.Handler {
	return func(ctx context.Context, ev events.Event) error {
		var order orders.Order
		if err := ev.Decode(&order); err != nil {
			logger.Error(err, "handler: order_confirmed: decode")
			return nil
		}
		_, err := notifSvc.Send(ctx, notifications.SendInput{
			UserID:   order.UserID,
			Type:     "order_confirmed",
			Channel:  "in_app",
			Title:    "Booking Confirmed 🏨",
			Body:     fmt.Sprintf("Your booking #%s has been confirmed. We look forward to welcoming you!", order.ID),
			Metadata: map[string]any{"order_id": order.ID},
		})
		if err != nil {
			logger.Error(err, "handler: order_confirmed: create notification")
		}
		return nil
	}
}

func OrderConfirmedWithWS(notifSvc *notifications.Service, hub *ws.Hub) events.Handler {
	return func(ctx context.Context, ev events.Event) error {
		var order orders.Order
		if err := ev.Decode(&order); err != nil {
			logger.Error(err, "handler: order_confirmed: decode")
			return nil
		}
		n, err := notifSvc.Send(ctx, notifications.SendInput{
			UserID:   order.UserID,
			Type:     "order_confirmed",
			Channel:  "in_app",
			Title:    "Booking Confirmed 🏨",
			Body:     fmt.Sprintf("Your booking #%s has been confirmed.", order.ID),
			Metadata: map[string]any{"order_id": order.ID},
		})
		if err != nil {
			logger.Error(err, "handler: order_confirmed: notify")
		}
		if n != nil {
			msg, _ := ws.NewMessage(ws.TypeNotification, n)
			hub.SendToUser(order.UserID, msg)
		}
		statusMsg, _ := ws.NewMessage(ws.TypeOrderStatus, map[string]any{
			"order_id": order.ID,
			"status":   order.Status,
		})
		hub.SendToUser(order.UserID, statusMsg)
		return nil
	}
}
