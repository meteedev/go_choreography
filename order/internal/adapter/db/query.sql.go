// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: query.sql

package db

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

const getOrder = `-- name: GetOrder :one
SELECT id, customer_id, created_at, updated_at, update_reason, deleted_at, status FROM orders
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetOrder(ctx context.Context, id uuid.UUID) (Order, error) {
	row := q.db.QueryRowContext(ctx, getOrder, id)
	var i Order
	err := row.Scan(
		&i.ID,
		&i.CustomerID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.UpdateReason,
		&i.DeletedAt,
		&i.Status,
	)
	return i, err
}

const insertOrder = `-- name: InsertOrder :one
insert into 
  orders (
    created_at, 
    updated_at, 
    deleted_at, 
    customer_id, 
    status
  )
values
  (
    $1, 
    $2, 
    $3, 
    $4, 
    $5
  )
  RETURNING id, customer_id, created_at, updated_at, update_reason, deleted_at, status
`

type InsertOrderParams struct {
	CreatedAt  sql.NullTime
	UpdatedAt  sql.NullTime
	DeletedAt  sql.NullTime
	CustomerID sql.NullString
	Status     sql.NullString
}

func (q *Queries) InsertOrder(ctx context.Context, arg InsertOrderParams) (Order, error) {
	row := q.db.QueryRowContext(ctx, insertOrder,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.DeletedAt,
		arg.CustomerID,
		arg.Status,
	)
	var i Order
	err := row.Scan(
		&i.ID,
		&i.CustomerID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.UpdateReason,
		&i.DeletedAt,
		&i.Status,
	)
	return i, err
}

const insertOrderItems = `-- name: InsertOrderItems :one
  insert into 
    order_items (
      created_at, 
      updated_at, 
      deleted_at, 
      product_code, 
      unit_price, 
      quantity, 
      order_id
    )
  values
    (
      $1, 
      $2, 
      $3, 
      $4, 
      $5, 
      $6, 
      $7
    )
    RETURNING id, created_at, updated_at, deleted_at, product_code, unit_price, quantity, order_id
`

type InsertOrderItemsParams struct {
	CreatedAt   sql.NullTime
	UpdatedAt   sql.NullTime
	DeletedAt   sql.NullTime
	ProductCode sql.NullString
	UnitPrice   sql.NullString
	Quantity    sql.NullInt32
	OrderID     uuid.NullUUID
}

func (q *Queries) InsertOrderItems(ctx context.Context, arg InsertOrderItemsParams) (OrderItem, error) {
	row := q.db.QueryRowContext(ctx, insertOrderItems,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.DeletedAt,
		arg.ProductCode,
		arg.UnitPrice,
		arg.Quantity,
		arg.OrderID,
	)
	var i OrderItem
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
		&i.ProductCode,
		&i.UnitPrice,
		&i.Quantity,
		&i.OrderID,
	)
	return i, err
}

const listOrders = `-- name: ListOrders :many
SELECT id, customer_id, created_at, updated_at, update_reason, deleted_at, status FROM orders
ORDER BY created_at
`

func (q *Queries) ListOrders(ctx context.Context) ([]Order, error) {
	rows, err := q.db.QueryContext(ctx, listOrders)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Order
	for rows.Next() {
		var i Order
		if err := rows.Scan(
			&i.ID,
			&i.CustomerID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.UpdateReason,
			&i.DeletedAt,
			&i.Status,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateOrders = `-- name: UpdateOrders :one
update 
  orders 
set
  updated_at = $1,
  update_reason=$2,
  status = $3
where 
  id = $4
  RETURNING id, customer_id, created_at, updated_at, update_reason, deleted_at, status
`

type UpdateOrdersParams struct {
	UpdatedAt    sql.NullTime
	UpdateReason sql.NullString
	Status       sql.NullString
	ID           uuid.UUID
}

func (q *Queries) UpdateOrders(ctx context.Context, arg UpdateOrdersParams) (Order, error) {
	row := q.db.QueryRowContext(ctx, updateOrders,
		arg.UpdatedAt,
		arg.UpdateReason,
		arg.Status,
		arg.ID,
	)
	var i Order
	err := row.Scan(
		&i.ID,
		&i.CustomerID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.UpdateReason,
		&i.DeletedAt,
		&i.Status,
	)
	return i, err
}
