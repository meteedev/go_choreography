package ports

import (
	"context"

	"github.com/meteedev/go_choreography/order/internal/application/core/domain"
	"github.com/meteedev/go_choreography/pkg/event"
)

type OrderServicePort interface {
	CreateOrder(ctx context.Context, order event.OrderCreateEvent) (*domain.Order, error)
	UpdateOrder(ctx context.Context, event event.OrderUpdateEvent) (*domain.Order, error)
}
