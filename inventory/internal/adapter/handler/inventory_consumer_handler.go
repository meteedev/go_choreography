package handler

import (
	"context"
	"encoding/json"
	"log"

	"github.com/meteedev/go_choreography/inventory/internal/application/core/service"
	"github.com/meteedev/go_choreography/inventory/internal/ports"
	"github.com/meteedev/go_choreography/pkg/event"
)

type InventoryConsumerHandler struct {
	InvService ports.InventoryServicePort
}

func NewInventoryConsumerHandler(s service.InventoryService) *InventoryConsumerHandler {
	return &InventoryConsumerHandler{InvService: s}
}

// Handlers for each event
func (c InventoryConsumerHandler) HandleOrderEvent(ctx context.Context, payload []byte) error {
	log.Printf("Handling order event with payload: %s", payload)

	var msg event.OrderCreateEvent

	err := json.Unmarshal(payload, &msg)

	if err != nil {
		log.Println(err.Error())
	}

	c.InvService.CheckInvBalance(ctx, msg)

	// Implement order event handling logic here
	return nil
}

func (c InventoryConsumerHandler) HandleOrderFailed(ctx context.Context, payload []byte) error {
	log.Printf("HandleOrderFailed with payload: %s", payload)
	// Implement order event handling logic here
	return nil
}

func (c InventoryConsumerHandler) HandlePaymentEvent(ctx context.Context, payload []byte) error {
	log.Printf("HandlePaymentEvent with payload: %s", payload)
	// Implement payment event handling logic here
	return nil
}

func (c InventoryConsumerHandler) HandleInventoryEvent(ctx context.Context, payload []byte) error {
	return nil
}
