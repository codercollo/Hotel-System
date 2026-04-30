// Package main is the entrypoint for the platform async worker process.
// It consumes events from the broker and dispatches them to domain handlers.
// Run separately from the API server to separate HTTP and async workloads.
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
	"github.com/codercollo/hotel-system/backend/pkg/logger"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		fmt.Fprintf(os.Stderr, "config error: %v\n", err)
		os.Exit(1)
	}

	logger.Init(cfg.Log.Level, cfg.Log.Format)
	logger.Info("starting platform worker")

	ctx := context.Background()

	pool, err := database.Init(ctx, cfg.Database)
	if err != nil {
		logger.Fatal(err, "database init failed")
	}
	defer pool.Close()

	var broker events.Broker
	if cfg.Events.Broker == "rabbitmq" {
		broker, err = events.NewRabbitMQBroker(cfg.Rabbit.URL, cfg.Rabbit.Exchange)
		if err != nil {
			logger.Fatal(err, "rabbitmq init failed")
		}
	} else {
		broker = events.NewInMemoryBroker()
		logger.Info("worker using in-memory broker (no async delivery in this mode)")
	}
	defer broker.Close()

	// ── Register event handlers (populated in Phase 6) ───────────────────────
	// Example:
	// broker.Subscribe(events.TopicOrderConfirmed, handlers.SendOrderConfirmationEmail(emailSvc))
	// broker.Subscribe(events.TopicPaymentCompleted, handlers.UpdateOrderOnPayment(orderSvc))

	logger.Info("worker ready — listening for events")

	// Block until signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	logger.Info("worker shutting down")
}
