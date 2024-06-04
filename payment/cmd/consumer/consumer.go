package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/meteedev/go_choreography/config"
	"github.com/meteedev/go_choreography/constant"
	"github.com/meteedev/go_choreography/payment/internal/adapter/handler"
	"github.com/meteedev/go_choreography/payment/internal/application/core/service"
	"github.com/meteedev/go_choreography/pkg/event"
	"github.com/meteedev/go_choreography/pkg/messenger"
)

func main() {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	m, err := messenger.NewRabbitMQ(ctx, config.GetAmqpURL())
	if err != nil {
		log.Fatalf("Failed to initialize Messenger: %v", err)
	}
	defer m.Close()

	ms := messenger.NewMessengerService(m)

	s := service.NewPaymentService(ms)

	h := handler.NewPaymentConsumerHandler(s)

	events := []event.EventConfig{
		{Name: constant.InventoryQueue, Handler: h.HandleInventoryEvent},
	}

	var wg sync.WaitGroup

	// Handle OS signals for graceful shutdown
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigChan
		log.Println("Received shutdown signal")
		cancel()
	}()

	// Start consumers for each event
	for _, event := range events {
		wg.Add(1)
		go startConsumer(ctx, &wg, event, m)
	}

	// Wait for all consumers to finish before shutting down
	wg.Wait()
	log.Println("All consumers shut down gracefully")

}

// startConsumer sets up and starts a RabbitMQ consumer for the given event.
func startConsumer(ctx context.Context, wg *sync.WaitGroup, event event.EventConfig, m messenger.Messenger) {
	defer wg.Done()

	consumerCtx, consumerCancel := context.WithCancel(ctx)
	defer consumerCancel()

	// Start consuming messages for this event
	go func() {
		err := m.Consume(consumerCtx, event.Name, event.Handler)
		if err != nil {
			log.Fatalf("Message service error for event %s: %v", event.Name, err)
		}
	}()

	// Wait until the context is cancelled
	<-ctx.Done()
	log.Printf("Shutting down consumer for event %s", event.Name)
}
