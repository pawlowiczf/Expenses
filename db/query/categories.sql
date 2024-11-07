-- name: CreateCategory :one 
INSERT INTO categories (name, description) 
VALUES ($1, $2) 
RETURNING *; 

-- name: GetCategory :one 
SELECT * FROM categories
WHERE id = $1 LIMIT 1; 

-- name: GetCategoryName :one 
SELECT name from categories 
WHERE id = $1 LIMIT 1;