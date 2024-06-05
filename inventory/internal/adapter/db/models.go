// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package db

import (
	"database/sql"
)

type Inventory struct {
	ID              int32
	ProductCode     string
	ProductName     string
	Description     sql.NullString
	QuantityInStock int32
	UnitPrice       string
	ReorderLevel    int32
	CreatedAt       sql.NullTime
	UpdatedAt       sql.NullTime
	DeletedAt       sql.NullTime
}
