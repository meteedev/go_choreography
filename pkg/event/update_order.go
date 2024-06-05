package event

import (
	"github.com/google/uuid"
)

type OrderUpdateEvent struct {
	OrderID     uuid.UUID `json:"order_id"`
	ProcessName string    `json:"process_name"`
	Status      string    `json:"status"`
	Reason      string    `json:"reason"`
}
