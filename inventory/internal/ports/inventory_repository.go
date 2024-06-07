package ports

import (
	"context"

	"github.com/google/uuid"
	"github.com/meteedev/go_choreography/inventory/internal/adapter/db"
)

type InventoryRepoPort interface {
	GetProductQuantity(ctx context.Context, productCode string) (int32, error)
	UpdateProductQuantity(ctx context.Context, arg db.UpdateProductQuantityParams) (db.Inventory, error)
	InsertReservations(ctx context.Context, arg db.InsertReservationsParams) (db.Reservation, error)
	DeleteReservations(ctx context.Context, arg db.DeleteReservationsParams) (db.Reservation, error)
	GetReservationsByOrderId(ctx context.Context, orderID uuid.NullUUID) ([]db.Reservation, error)
	ExecTx(ctx context.Context, fn func(*db.Queries) error) error
}
