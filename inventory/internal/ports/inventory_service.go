package ports

import (
	"context"

	"github.com/meteedev/go_choreography/pkg/event"
)

type InventoryServicePort interface {
	CheckInvBalance(ctx context.Context, order event.OrderCreateEvent) (*event.OrderUpdateEvent, error)
	CompensateOrder(ctx context.Context, order event.OrderCreateEvent) (*event.OrderUpdateEvent, error)
}
