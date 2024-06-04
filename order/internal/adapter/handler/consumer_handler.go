package handler

import (
	"context"
	"log"

	"github.com/meteedev/go_choreography/order/internal/ports"
)

type OrderConsumerHandler struct {
	OrderService ports.OrderServicePort
}

func NewOrderConsumerHandler(s ports.OrderServicePort) *OrderConsumerHandler {
	return &OrderConsumerHandler{OrderService: s}
}

// Handlers for each event
func (o OrderConsumerHandler) HandleOrderUpdateEvent(ctx context.Context, payload []byte) error {
	log.Printf("HandleOrderUpdateEvent with payload: %s", payload)
	// Implement order event handling logic here
	return nil
}

// Handlers for each event
func (o OrderConsumerHandler) HandleOrderCompensate(ctx context.Context, payload []byte) error {
	log.Printf("HandleOrderCompensate with payload: %s", payload)
	// Implement order event handling logic here
	return nil
}
