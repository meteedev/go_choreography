package service

import (
	"context"
	"encoding/json"

	"github.com/google/uuid"
	"github.com/meteedev/go_choreography/constant"
	"github.com/meteedev/go_choreography/pkg/event"
	"github.com/meteedev/go_choreography/pkg/messenger"
)

type PaymentService struct {
	MessageService messenger.MessengerService
}

func NewPaymentService(m messenger.MessengerService) PaymentService {
	return PaymentService{
		MessageService: m,
	}
}

func (p PaymentService) Pay(ctx context.Context, order event.OrderCreateEvent) (*event.OrderUpdateEvent, error) {

	newOrderID := uuid.New()

	event := event.OrderUpdateEvent{
		OrderID:     newOrderID,
		ProcessName: "Payment.Pay",
		Status:      true,
		Reason:      "",
	}

	msg, err := json.Marshal(event)

	if err != nil {
		return nil, err
	}

	p.MessageService.Messenger.Publish(ctx, msg, constant.Order_update, false)
	return nil, nil
}

// func (p PaymentService) Pay(ctx context.Context, order event.OrderCreateEvent) (*event.OrderUpdateEvent, error) {

// 	newOrderID := uuid.New()

// 	event := event.OrderUpdateEvent{
// 		OrderID:     newOrderID,
// 		ProcessName: "Payment.Pay",
// 		Status:      false,
// 		Reason:      "payment failed",
// 	}

// 	msg, err := json.Marshal(event)

// 	if err != nil {
// 		return nil, err
// 	}

// 	p.MessageService.Messenger.Publish(ctx, msg, constant.OrderFail,true)
// 	return nil, nil
// }
