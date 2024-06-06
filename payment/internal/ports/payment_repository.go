package ports

import (
	"context"

	"github.com/meteedev/go_choreography/payment/internal/adapter/db"
)

type PaymentRepositoryPort interface {
	InsertPayments(ctx context.Context, arg db.InsertPaymentsParams) (db.Payment, error)
}
