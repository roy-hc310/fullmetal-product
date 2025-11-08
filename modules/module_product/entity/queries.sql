-- name: CreateProduct :one
INSERT INTO products (
  id, 
  brand_id, 
  category_id, 
  shop_id, 
  name, 
  sku, 
  description, 
  is_active, 
  status
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8, $9
)
RETURNING id;

-- name: GetDetailProduct :one
SELECT id, 
  created_at, 
  updated_at, 
  deleted_at, 
  brand_id, 
  category_id, 
  shop_id, 
  name, 
  sku, 
  description, 
  is_active, 
  status
FROM products
WHERE deleted_at IS NULL AND id = $1;

-- name: GetListProduct :many
SELECT id, 
  created_at, 
  updated_at, 
  deleted_at, 
  brand_id, 
  category_id, 
  shop_id, 
  name, 
  sku, 
  description, 
  is_active, 
  status 
FROM products
WHERE deleted_at IS NULL 
AND (sqlc.narg('shop_id')::uuid IS NULL OR shop_id = sqlc.narg('shop_id')::uuid)
AND (
  sqlc.narg('cursor')::text IS NULL OR
  (sqlc.narg('sort_by')::text = 'id ASC' AND id::text > sqlc.narg('cursor')::text) OR
  (sqlc.narg('sort_by')::text = 'id DESC' AND id::text < sqlc.narg('cursor')::text) OR
  (sqlc.narg('sort_by')::text = 'updated_at ASC' AND updated_at > sqlc.narg('cursor')::timestamp) OR
  (sqlc.narg('sort_by')::text = 'updated_at DESC' AND updated_at < sqlc.narg('cursor')::timestamp)
)
ORDER BY
  CASE WHEN sqlc.narg('sort_by')::text = 'id ASC' THEN id::text END ASC,
  CASE WHEN sqlc.narg('sort_by')::text = 'id DESC' THEN id::text END DESC,
  CASE WHEN sqlc.narg('sort_by')::text = 'updated_at ASC' THEN updated_at END ASC,
  CASE WHEN sqlc.narg('sort_by')::text = 'updated_at DESC' THEN updated_at END DESC,
  id ASC
LIMIT COALESCE(sqlc.narg('limit')::int, 10);

-- name: CountProduct :one
SELECT COUNT(*) FROM products
WHERE deleted_at IS NULL;

-- name: UpdateProduct :exec
UPDATE products
SET 
  updated_at = now(),
  name = COALESCE($2, name),
  sku = COALESCE($3, sku),
  description = COALESCE($4, description),
  is_active = COALESCE($5, is_active),
  status = COALESCE($6, status)
WHERE id = $1 AND deleted_at IS NULL;

-- name: DeleteProduct :exec
UPDATE products
SET deleted_at = now()
WHERE deleted_at IS NULL AND id = $1;