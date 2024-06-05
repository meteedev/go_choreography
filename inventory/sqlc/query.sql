-- name: GetProduct :one
SELECT * FROM inventory
WHERE product_code = $1 LIMIT 1;

-- name: GetProductQuantity :one
SELECT count(quantity_in_stock) FROM inventory
WHERE product_code = $1 ;
