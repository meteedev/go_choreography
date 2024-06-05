package ports

import (
	"context"
)

type InventoryRepoPort interface {
	GetProductQuantity(ctx context.Context, productCode string) (int32, error)
}
