-- name: FetchUserByUsername :one
SELECT u.*
  FROM users u
  WHERE u.username = @email::text LIMIT 1;

-- name: CreateUser :one
INSERT INTO users (
  username, encrypted_password
) VALUES (
  @username::text, @encrypted_password::text
) RETURNING *;
