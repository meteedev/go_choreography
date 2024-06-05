package service

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
	"github.com/meteedev/go_choreography/constant"
	"github.com/meteedev/go_choreography/inventory/internal/ports"
	"github.com/meteedev/go_choreography/pkg/event"
	"github.com/meteedev/go_choreography/pkg/messenger"
)

type InventoryService struct {
	MessageService messenger.MessengerService
	InventoryRepo  ports.InventoryRepoPort
}

func NewInventoryService(m messenger.MessengerService, r ports.InventoryRepoPort) InventoryService {
	return InventoryService{
		InventoryRepo:  r,
		MessageService: m,
	}
}

func (i InventoryService) CheckInvBalance(ctx context.Context, order event.OrderCreateEvent) (*event.OrderUpdateEvent, error) {
	// Helper function to publish a message
	publishMessage := func(ctx context.Context, msg interface{}, queue string) error {
		msgData, err := json.Marshal(msg)
		if err != nil {
			return err
		}
		return i.MessageService.Messenger.Publish(ctx, msgData, queue, false)
	}

	// Loop through the order items
	for _, item := range order.OrderItems {
		count, err := i.InventoryRepo.GetProductQuantity(ctx, item.ProductCode)
		if err != nil {
			return nil, err
		}

		if count < int64(item.Quantity) {
			updateEvent := event.OrderUpdateEvent{
				OrderID:     order.ID,
				ProcessName: "Inv.CheckInvBalance",
				Status:      false,
				Reason:      fmt.Sprintf("Insufficient inventory for product code %s", item.ProductCode),
			}
			if err := publishMessage(ctx, updateEvent, constant.OrderUpdateQueue); err != nil {
				return nil, err
			}
			return &updateEvent, nil
		}
	}

	// If all items have sufficient inventory
	orderUpdateEvent := event.OrderUpdateEvent{
		OrderID:     order.ID,
		ProcessName: "Inv.CheckInvBalance",
		Status:      true,
		Reason:      "Inventory check passed",
	}

	// Publish the messages
	if err := publishMessage(ctx, order, constant.InventoryQueue); err != nil {
		return nil, err
	}

	if err := publishMessage(ctx, orderUpdateEvent, constant.OrderUpdateQueue); err != nil {
		return nil, err
	}

	return &orderUpdateEvent, nil
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
