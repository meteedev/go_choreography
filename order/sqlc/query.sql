-- name: GetOrder :one
SELECT * FROM orders
WHERE id = $1 LIMIT 1;

-- name: ListOrders :many
SELECT * FROM orders
ORDER BY created_at;


-- name: InsertOrder :one
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
  RETURNING * ;

-- name: InsertOrderItems :one
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
    RETURNING * ;

-- name: UpdateOrders :one
update 
  orders 
set
  updated_at = $1,
  update_reason=$2,
  status = $3
where 
  id = $4
  RETURNING * ;