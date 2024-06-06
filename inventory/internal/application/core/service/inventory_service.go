package service

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
	"github.com/meteedev/go_choreography/constant"
	"github.com/meteedev/go_choreography/inventory/internal/adapter/db"
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

// Helper function to publish a message
func (i InventoryService) publishMessage(ctx context.Context, msg interface{}, queue string) error {
	msgData, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	return i.MessageService.Messenger.Publish(ctx, msgData, queue, false)
}

// Helper function to create order update event
func newOrderUpdateEvent(orderID uuid.UUID, processName string, status string, reason string) event.OrderUpdateEvent {
	return event.OrderUpdateEvent{
		OrderID:     orderID,
		ProcessName: processName,
		Status:      status,
		Reason:      reason,
	}
}

func (i InventoryService) CheckInvBalance(ctx context.Context, order event.OrderCreateEvent) (*event.OrderUpdateEvent, error) {
	// Check inventory balance
	updateEvent, sufficient, err := i.checkInventory(ctx, order)
	if err != nil {
		return nil, err
	}
	if !sufficient {
		return updateEvent, nil
	}

	// Create order update event
	orderUpdateEvent := newOrderUpdateEvent(order.ID, "Inv.CheckInvBalance", constant.Order_status_processing, "Inventory check passed")

	// Publish the messages
	if err := i.publishMessage(ctx, order, constant.Inventory_reserved); err != nil {
		return nil, err
	}
	if err := i.publishMessage(ctx, orderUpdateEvent, constant.Order_update); err != nil {
		return nil, err
	}

	return &orderUpdateEvent, nil
}

func (i InventoryService) ReservedProduct(ctx context.Context, order event.OrderCreateEvent) (*event.OrderUpdateEvent, error) {
	// Check inventory balance
	updateEvent, sufficient, err := i.checkInventory(ctx, order)
	if err != nil {
		return nil, err
	}
	if !sufficient {
		return updateEvent, nil
	}

	// Execute the transaction
	err = i.InventoryRepo.ExecTx(ctx, func(q *db.Queries) error {
		for _, item := range order.OrderItems {
			// Update the product quantity
			_, err := q.UpdateProductQuantity(ctx, db.UpdateProductQuantityParams{
				QuantityInStock: int32(item.Quantity),
				ProductCode:     item.ProductCode,
			})
			if err != nil {
				return err
			}

			// Insert the reservation
			_, err = q.InsertReservations(ctx, db.InsertReservationsParams{
				OrderID:     uuid.NullUUID{UUID: order.ID, Valid: true},
				ProductCode: item.ProductCode,
				Quantity:    int32(item.Quantity),
			})
			if err != nil {
				// Publish error message to order_update queue
				errUpdateEvent := newOrderUpdateEvent(order.ID, "Inv.CheckInvBalance", constant.Order_status_fail, fmt.Sprintf("Failed to reserve product: %v", err))
				if publishErr := i.publishMessage(ctx, errUpdateEvent, constant.Order_update); publishErr != nil {
					return fmt.Errorf("error publishing reservation failure event: %v", publishErr)
				}
				return err
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	// Create order update event after successful reservation
	orderUpdateEvent := newOrderUpdateEvent(order.ID, "Inv.CheckInvBalance", constant.Order_status_processing, "Inventory check and reservation passed")

	// Publish the messages
	if err := i.publishMessage(ctx, order, constant.Inventory_reserved); err != nil {
		return nil, err
	}
	if err := i.publishMessage(ctx, orderUpdateEvent, constant.Order_update); err != nil {
		return nil, err
	}

	return &orderUpdateEvent, nil
}

func (i InventoryService) CompensateOrder(ctx context.Context, order event.OrderCreateEvent) (*event.OrderUpdateEvent, error) {

	// Execute the transaction
	err := i.InventoryRepo.ExecTx(ctx, func(q *db.Queries) error {
		for _, item := range order.OrderItems {
			// Compensate by adding back the reserved quantity to the inventory
			_, err := q.UpdateProductQuantity(ctx, db.UpdateProductQuantityParams{
				QuantityInStock: -int32(item.Quantity), // Subtracting the quantity from the inventory
				ProductCode:     item.ProductCode,
			})
			if err != nil {
				return err
			}

			// Delete the reservation
			_, err = q.DeleteReservations(ctx, db.DeleteReservationsParams{
				OrderID:     uuid.NullUUID{UUID: order.ID, Valid: true},
				ProductCode: item.ProductCode,
			})
			if err != nil {
				// If delete reservation failed, publish an error message to order_update
				errUpdateEvent := newOrderUpdateEvent(order.ID, "Inv.CompensateOrder", constant.Order_status_fail, fmt.Sprintf("Failed to delete reservation: %v", err))
				if publishErr := i.publishMessage(ctx, errUpdateEvent, constant.Order_update); publishErr != nil {
					return fmt.Errorf("error publishing compensation failure event: %v", publishErr)
				}
				return err
			}
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	// Create the order update event after successful compensation
	orderUpdateEvent := newOrderUpdateEvent(order.ID, "Inv.CompensateOrder", constant.Order_status_fail, "")

	return &orderUpdateEvent, nil
}

// Helper function to check inventory balance
func (i InventoryService) checkInventory(ctx context.Context, order event.OrderCreateEvent) (*event.OrderUpdateEvent, bool, error) {
	for _, item := range order.OrderItems {
		count, err := i.InventoryRepo.GetProductQuantity(ctx, item.ProductCode)
		if err != nil {
			return nil, false, err
		}

		if count < int32(item.Quantity) {
			updateEvent := newOrderUpdateEvent(order.ID, "Inv.CheckInvBalance", constant.Order_status_fail, fmt.Sprintf("Insufficient inventory for product code %s", item.ProductCode))
			if err := i.publishMessage(ctx, updateEvent, constant.Order_update); err != nil {
				return nil, false, err
			}
			return &updateEvent, false, nil
		}
	}
	return nil, true, nil
}
