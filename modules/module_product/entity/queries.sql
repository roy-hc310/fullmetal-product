-- name: CreateProduct :one
INSERT INTO products (
  product_id, name, sku, description
) VALUES (
  $1, $2, $3, $4
)
RETURNING product_id;

-- name: GetDetailProduct :one
SELECT * FROM products
WHERE product_id = $1 LIMIT 1;

-- name: GetListProduct :many
SELECT * FROM products
LIMIT $1 OFFSET $2;

-- name: UpdateProduct :exec
UPDATE products
  set name = $2
WHERE product_id = $1;

-- name: DeleteProduct :exec
DELETE FROM products
WHERE product_id = $1;