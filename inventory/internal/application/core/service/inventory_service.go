package service

import (
	"context"
	"encoding/json"

	"github.com/google/uuid"
	"github.com/meteedev/go_choreography/constant"
	"github.com/meteedev/go_choreography/pkg/event"
	"github.com/meteedev/go_choreography/pkg/messenger"
)

type InventoryService struct {
	MessageService messenger.MessengerService
}

func NewInventoryService(m messenger.MessengerService) InventoryService {
	return InventoryService{
		MessageService: m,
	}
}

func (i InventoryService) CheckInvBalance(ctx context.Context, order event.OrderCreateEvent) (*event.OrderUpdateEvent, error) {

	newOrderID := uuid.New()

	event := event.OrderUpdateEvent{
		OrderID:     newOrderID,
		ProcessName: "Inv.CheckInvBalance",
		Status:      true,
		Reason:      "",
	}

	msgEvent, err := json.Marshal(event)

	if err != nil {
		return nil, err
	}

	msgOrder, err := json.Marshal(order)

	if err != nil {
		return nil, err
	}

	i.MessageService.Messenger.Publish(ctx, msgOrder, constant.InventoryQueue, false)
	i.MessageService.Messenger.Publish(ctx, msgEvent, constant.OrderUpdateQueue, false)

	return &event, nil
}

func (i InventoryService) CompensateOrder(ctx context.Context, order event.OrderCreateEvent) (*event.OrderUpdateEvent, error) {

	newOrderID := uuid.New()

	event := event.OrderUpdateEvent{
		OrderID:     newOrderID,
		ProcessName: "Inv.CompensateOrder",
		Status:      true,
		Reason:      "",
	}

	msgEvent, err := json.Marshal(event)

	if err != nil {
		return nil, err
	}

	i.MessageService.Messenger.Publish(ctx, msgEvent, constant.OrderUpdateQueue, false)

	return &event, nil
}
