-- name: GetProduct :one
SELECT * FROM inventory
WHERE product_code = $1 LIMIT 1;

-- name: GetProductQuantity :one
SELECT COALESCE(quantity_in_stock, 0) as count FROM inventory
WHERE product_code = $1 ;


-- name: UpdateProductQuantity :one
UPDATE inventory 
SET quantity_in_stock = quantity_in_stock - $1, updated_at = NOW() 
WHERE product_code = $2
RETURNING * ;

-- name: InsertReservations :one
INSERT INTO reservations (order_id, product_code, quantity) 
VALUES ($1, $2, $3)
RETURNING * ;

-- name: DeleteReservations :one
DELETE FROM reservations where order_id = $1 and  product_code = $2 
RETURNING * ;

-- name: GetReservationsByOrderId :many
SELECT * FROM reservations WHERE order_id = $1 ;