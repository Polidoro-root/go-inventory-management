-- name: FindUserByID :one
SELECT *
FROM users
WHERE id = $1;
-- name: FindUserByEmail :one
SELECT *
FROM users
WHERE email = $1;
-- name: UserEmailExist :one
SELECT EXISTS (
  SELECT 1 FROM users
  WHERE email = $1
);
-- name: UserPhoneNumberExist :one
SELECT EXISTS (
  SELECT 1 FROM users
  WHERE phone_number = $1
);
-- name: SaveUser :exec
INSERT INTO users (
    id,
    name,
    role,
    email,
    phone_number,
    password,
    created_at,
    updated_at
  )
VALUES ($1, $2, $3, $4, $5, $6, $7, $8);