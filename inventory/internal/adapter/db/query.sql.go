// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: query.sql

package db

import (
	"context"
)

const getProduct = `-- name: GetProduct :one
SELECT id, product_code, product_name, description, quantity_in_stock, unit_price, reorder_level, created_at, updated_at, deleted_at FROM inventory
WHERE product_code = $1 LIMIT 1
`

func (q *Queries) GetProduct(ctx context.Context, productCode string) (Inventory, error) {
	row := q.db.QueryRowContext(ctx, getProduct, productCode)
	var i Inventory
	err := row.Scan(
		&i.ID,
		&i.ProductCode,
		&i.ProductName,
		&i.Description,
		&i.QuantityInStock,
		&i.UnitPrice,
		&i.ReorderLevel,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const getProductQuantity = `-- name: GetProductQuantity :one
SELECT count(quantity_in_stock) FROM inventory
WHERE product_code = $1
`

func (q *Queries) GetProductQuantity(ctx context.Context, productCode string) (int64, error) {
	row := q.db.QueryRowContext(ctx, getProductQuantity, productCode)
	var count int64
	err := row.Scan(&count)
	return count, err
}
