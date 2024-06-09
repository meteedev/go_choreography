package handler

import (
	"context"
	"encoding/json"
	"log"

	"github.com/meteedev/go_choreography/payment/internal/application/core/service"
	"github.com/meteedev/go_choreography/payment/internal/ports"
	"github.com/meteedev/go_choreography/pkg/event"
)

type PaymentConsumerHandler struct {
	PaymentService ports.PaymentServicePort
}

func NewPaymentConsumerHandler(s service.PaymentService) *PaymentConsumerHandler {
	return &PaymentConsumerHandler{PaymentService: s}
}

// Handlers for each event
func (p PaymentConsumerHandler) HandleInventoryEvent(ctx context.Context, payload []byte) error {
	//log.Printf("Handling order event with payload: %s", payload)

	var msg event.OrderCreateEvent

	err := json.Unmarshal(payload, &msg)

	if err != nil {
		log.Println(err.Error())
	}

	p.PaymentService.Pay(ctx, msg)

	// Implement order event handling logic here
	return nil
}
