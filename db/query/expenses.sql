-- name: CreateExpense :one 
INSERT INTO expenses (user_id, category_id, amount, description, date) 
VALUES ($1, $2, $3, $4, $5) 
RETURNING *; 

-- name: GetExpense :one 
SELECT * FROM expenses
WHERE id = $1 LIMIT 1; 

