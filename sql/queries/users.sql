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

-- name: UpdateUser :one
UPDATE users
SET email = $2, password_hash = $3, updated_at = $4
WHERE id = $1
RETURNING *; 

-- name: UpdateIsChirpyRed :exec
UPDATE users
SET is_chirpy_red = $2, updated_at = $3
WHERE id = $1;