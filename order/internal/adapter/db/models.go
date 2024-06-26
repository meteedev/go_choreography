// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package db

import (
	"database/sql"

	"github.com/google/uuid"
)

type Order struct {
	ID           uuid.UUID
	CustomerID   sql.NullString
	CreatedAt    sql.NullTime
	UpdatedAt    sql.NullTime
	UpdateReason sql.NullString
	DeletedAt    sql.NullTime
	Status       sql.NullString
}

type OrderItem struct {
	ID          int32
	CreatedAt   sql.NullTime
	UpdatedAt   sql.NullTime
	DeletedAt   sql.NullTime
	ProductCode sql.NullString
	UnitPrice   sql.NullString
	Quantity    sql.NullInt32
	OrderID     uuid.NullUUID
}
