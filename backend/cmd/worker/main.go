package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/codercollo/hotel-system/backend/config"
	"github.com/codercollo/hotel-system/backend/database"
	"github.com/codercollo/hotel-system/backend/events"
	evhandlers "github.com/codercollo/hotel-system/backend/events/handlers"
	"github.com/codercollo/hotel-system/backend/internal/notifications"
	notifchannels "github.com/codercollo/hotel-system/backend/internal/notifications/channels"
	"github.com/codercollo/hotel-system/backend/pkg/cache"
	"github.com/codercollo/hotel-system/backend/pkg/logger"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		fmt.Fprintf(os.Stderr, "config: %v\n", err)
		os.Exit(1)
	}
	logger.Init(cfg.Log.Level, cfg.Log.Format)
	logger.Info("platform worker starting")

	ctx := context.Background()

	db, err := database.Init(ctx, cfg.Database)
	if err != nil {
		logger.Fatal(err, "database")
	}
	defer db.Close()

	_, err = cache.NewRedis(cfg.Redis.URL)
	if err != nil {
		logger.Fatal(err, "redis")
	}

	var broker events.Broker
	if cfg.Events.Broker == "rabbitmq" {
		broker, err = events.NewRabbitMQBroker(cfg.Rabbit.URL, cfg.Rabbit.Exchange)
		if err != nil {
			logger.Fatal(err, "rabbitmq")
		}
	} else {
		broker = events.NewInMemoryBroker()
	}
	defer broker.Close()

	notifRepo := notifications.NewPostgresRepository(db)
	notifSvc := notifications.NewService(notifRepo, []notifchannels.Channel{
		notifchannels.NewEmailChannel("localhost", "1025", "", "", "noreply@platform.local"),
		notifchannels.NewSMSChannel("", ""),
		notifchannels.NewPushChannel(""),
	})

	cancels := []func(){
		broker.Subscribe(events.TopicOrderConfirmed, evhandlers.OrderConfirmed(notifSvc)),
		broker.Subscribe(events.TopicPaymentCompleted, evhandlers.PaymentCompleted(notifSvc)),
		broker.Subscribe(events.TopicOrderPlaced, evhandlers.AuditLog()),
		broker.Subscribe(events.TopicOrderCancelled, evhandlers.AuditLog()),
		broker.Subscribe(events.TopicPaymentFailed, evhandlers.AuditLog()),
		broker.Subscribe(events.TopicUserRegistered, evhandlers.AuditLog()),
		broker.Subscribe(events.TopicItemCreated, evhandlers.AuditLog()),
	}

	logger.Info(fmt.Sprintf("worker ready — %d subscriptions active", len(cancels)))

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	logger.Info("worker shutting down")
	for _, c := range cancels {
		c()
	}
}
