package handler

import (
	"context"
	"encoding/json"
	"log"

	"github.com/meteedev/go_choreography/order/internal/ports"
	"github.com/meteedev/go_choreography/pkg/event"
)

type OrderConsumerHandler struct {
	OrderService ports.OrderServicePort
}

func NewOrderConsumerHandler(s ports.OrderServicePort) *OrderConsumerHandler {
	return &OrderConsumerHandler{OrderService: s}
}

// Handlers for each event
func (o OrderConsumerHandler) HandleOrderUpdateEvent(ctx context.Context, payload []byte) error {
	//log.Printf("HandleOrderUpdateEvent with payload: %s", payload)
	var msg event.OrderUpdateEvent

	err := json.Unmarshal(payload, &msg)

	if err != nil {
		log.Println(err.Error())
	}

	o.OrderService.UpdateOrder(ctx, msg)

	return nil
}

// Handlers for each event
func (o OrderConsumerHandler) HandleOrderCompensate(ctx context.Context, payload []byte) error {
	//log.Printf("HandleOrderCompensate with payload: %s", payload)
	// Implement order event handling logic here
	return nil
}
