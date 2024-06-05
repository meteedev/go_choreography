-- name: GetProduct :one
SELECT * FROM inventory
WHERE product_code = $1 LIMIT 1;

-- name: GetProductQuantity :one
SELECT quantity_in_stock as count FROM inventory
WHERE product_code = $1 ;


-- name: UpdateProductQuantity :one
UPDATE inventory 
SET quantity_in_stock = quantity_in_stock - $1, updated_at = NOW() 
WHERE product_code = $2
RETURNING * ;

-- name: insertReservations :one
INSERT INTO reservations (order_id, product_code, quantity) 
VALUES ($1, $2, $3)
RETURNING * ;