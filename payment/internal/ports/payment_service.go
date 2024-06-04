package ports

import (
	"context"

	"github.com/meteedev/go_choreography/pkg/event"
)

type PaymentServicePort interface {
	Pay(ctx context.Context, order event.OrderCreateEvent) (*event.OrderUpdateEvent, error)
}
