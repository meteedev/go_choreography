package domain

import (
	"github.com/google/uuid"
)

type OrderItem struct {
	ProductCode string  `json:"product_code"`
	UnitPrice   float32 `json:"unit_price"`
	Quantity    int32   `json:"quantity"`
}

type Order struct {
	ID         uuid.UUID   `json:"id"`
	Status     string      `json:"status"`
	OrderItems []OrderItem `json:"order_items"`
	CreatedAt  int64       `json:"created_at"`
	CustomerID string      `json:"customer_id"`
}

const (
	OrderStatusProcessing = "processing"
	OrderStatusFailed     = "failed"
	OrderStatusSuccess    = "success"
)
