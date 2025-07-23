-- name: CreateUser :one
INSERT INTO users (id, created_at, updated_at, email, password_hash)
VALUES (    
    $1, 
    $2, 
    $3, 
    $4,
    $5
)
RETURNING *;



-- name: DeleteUsers :exec
DELETE FROM users; 


-- name: GetUser :one
SELECT * FROM users
WHERE id = $1;

-- name: GetUserByEmail :one
SELECT * FROM users
WHERE email = $1;