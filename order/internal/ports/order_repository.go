package ports

import (
	"context"
	"github.com/meteedev/go_choreography/order/internal/adapter/db"
)

type OrderRepositoryPort interface {
	InsertOrder(ctx context.Context, arg db.InsertOrderParams) (db.Order, error)
	InsertOrderItems(ctx context.Context, arg db.InsertOrderItemsParams) (db.OrderItem, error)
	UpdateOrders(ctx context.Context, arg db.UpdateOrdersParams) (db.Order, error)
	ExecTx(ctx context.Context, fn func(*db.Queries) error) error
}
