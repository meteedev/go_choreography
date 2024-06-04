package service

import (
	"database/sql"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/meteedev/go_choreography/order/internal/adapter/db"
	"github.com/meteedev/go_choreography/order/internal/application/core/domain"
	"github.com/meteedev/go_choreography/pkg/event"
	"github.com/meteedev/go_choreography/pkg/util"
)

// Convert original Order to InsertOrderParams struct
func ConvertOrderEventToParams(original event.OrderCreateEvent) db.InsertOrderParams {
	return db.InsertOrderParams{
		CreatedAt: sql.NullTime{Time: time.Unix(original.CreatedAt, 0), Valid: true},
		UpdatedAt: sql.NullTime{Valid: false}, // Initialize as null
		DeletedAt: sql.NullTime{Valid: false}, // Initialize as null
		Status:    sql.NullString{String: original.Status, Valid: true},
	}
}

// Convert original OrderItem to InsertOrderItemsParams struct
func ConvertOrderEventItemToParams(original event.OrderItem, orderID uuid.UUID) db.InsertOrderItemsParams {
	return db.InsertOrderItemsParams{
		ProductCode: sql.NullString{String: original.ProductCode, Valid: true},
		UnitPrice:   sql.NullString{String: strconv.FormatFloat(float64(original.UnitPrice), 'f', 6, 32), Valid: true},
		Quantity:    sql.NullInt32{Int32: original.Quantity, Valid: true},
		OrderID:     uuid.NullUUID{UUID: orderID, Valid: true},
		CreatedAt:   sql.NullTime{Valid: false}, // Initialize as null
		UpdatedAt:   sql.NullTime{Valid: false}, // Initialize as null
		DeletedAt:   sql.NullTime{Valid: false}, // Initialize as null
	}
}

// Convert a slice of original OrderItem to a slice of InsertOrderItemsParams structs
func ConvertOrderEventItemsToParams(originalItems []event.OrderItem, orderID uuid.UUID) []db.InsertOrderItemsParams {
	var newItems []db.InsertOrderItemsParams
	for _, item := range originalItems {
		newItems = append(newItems, ConvertOrderEventItemToParams(item, orderID))
	}
	return newItems
}

// Conversion function for OrderItem
func convertOrderItem(item db.OrderItem) domain.OrderItem {

	return domain.OrderItem{
		ProductCode: util.ConvertNullString(item.ProductCode),
		UnitPrice:   util.ConvertNullFloat(item.UnitPrice),
		Quantity:    util.ConvertNullInt32(item.Quantity),
	}
}

// Conversion function for Order
func convertOrder(order db.Order, items []db.OrderItem) domain.Order {
	newItems := make([]domain.OrderItem, len(items))
	for i, item := range items {
		newItems[i] = convertOrderItem(item)
	}
	return domain.Order{
		ID:         order.ID,                                 // Assuming order.ID is uuid.UUID
		CustomerID: util.ConvertNullString(order.CustomerID), // Assuming CustomerID is originally a string representing an int
		Status:     util.ConvertNullString(order.Status),
		OrderItems: newItems,
		CreatedAt:  util.ConvertNullTime(order.CreatedAt),
	}
}

// func ConvertUpdateMsgToParams(msg invmsg.UpdateInventoryMessage, status string) db.UpdateOrdersParams {

// 	u := db.UpdateOrdersParams{
// 		UpdatedAt: sql.NullTime{Time: time.Now(), Valid: true},
// 		Status:    sql.NullString{String: status, Valid: true},
// 		ID:        msg.OrderId,
// 	}

// 	return u
// }
