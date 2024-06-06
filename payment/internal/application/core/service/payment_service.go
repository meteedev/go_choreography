package service

import (
	"context"
	"encoding/json"

	"github.com/google/uuid"
	"github.com/meteedev/go_choreography/constant"
	"github.com/meteedev/go_choreography/payment/internal/adapter/db"
	"github.com/meteedev/go_choreography/payment/internal/ports"
	"github.com/meteedev/go_choreography/pkg/event"
	"github.com/meteedev/go_choreography/pkg/messenger"
)

type PaymentService struct {
	PaymentRepo    ports.PaymentRepositoryPort
	MessageService messenger.MessengerService
}

func NewPaymentService(r ports.PaymentRepositoryPort, m messenger.MessengerService) PaymentService {
	return PaymentService{
		PaymentRepo:    r,
		MessageService: m,
	}
}

func (p PaymentService) Pay(ctx context.Context, order event.OrderCreateEvent) (*event.OrderUpdateEvent, error) {
	// Prepare payment data
	payment := db.InsertPaymentsParams{
		OrderID:       uuid.NullUUID{UUID: order.ID, Valid: true},
		Amount:        calculateTotalOrderAmount(order), // You need to implement this function
		PaymentMethod: "Credit Card",                    // Example payment method
		PaymentStatus: "Pending",                        // Initial payment status
	}

	// Insert payment record into the database
	_, err := p.PaymentRepo.InsertPayments(ctx, payment)

	var msg event.OrderUpdateEvent

	if err != nil {

		msg = event.OrderUpdateEvent{
			OrderID:     order.ID,
			ProcessName: "PaymentService.Pay",
			Status:      constant.Payment_failed,
			Reason:      err.Error(),
		}

		p.publishMessage(ctx, msg, constant.Payment_failed, true)
		return nil, err

	}
	msg = event.OrderUpdateEvent{
		OrderID:     order.ID,
		ProcessName: "PaymentService.Pay",
		Status:      constant.Payment_processed,
		Reason:      "",
	}
	p.publishMessage(ctx, msg, constant.Order_update, false)

	// Return a success message or event indicating the payment initiation
	return &msg, nil
}

func calculateTotalOrderAmount(order event.OrderCreateEvent) float64 {
	total := 0.0
	for _, item := range order.OrderItems {
		total += float64(item.UnitPrice) * float64(item.Quantity)
	}
	return total
}

func (p PaymentService) publishMessage(ctx context.Context, msg interface{}, queue string, isExchange bool) error {
	msgData, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	return p.MessageService.Messenger.Publish(ctx, msgData, queue, isExchange)
}
