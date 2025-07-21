-- name: CreateUser :one
INSERT INTO chirpy (id, created_at, updated_at, email)
VALUES (    
    $1, 
    $2, 
    $3, 
    $4
)
RETURNING *;

