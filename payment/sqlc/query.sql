-- name: InsertPayments :one
INSERT INTO payments (order_id, amount, payment_method, payment_status)
VALUES ($1, $2, $3, $4)
RETURNING *;