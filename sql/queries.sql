-- name: CreateUser :one
INSERT INTO users (email, password)
VALUES ($1, $2)
RETURNING *;

-- name: GetUser :one
SELECT *
FROM users
WHERE id = $1
LIMIT 1;

-- name: UpdateUser :one
UPDATE users
SET email = $2,
    password  = $3
WHERE id = $1
RETURNING *;

-- name: PartialUpdateUser :one
UPDATE users
SET email = CASE WHEN @update_email::boolean THEN @email::VARCHAR(255) ELSE email END,
    password  = CASE WHEN @update_password::boolean THEN @password::VARCHAR(255) ELSE password END
WHERE id = @id
RETURNING *;

-- name: DeleteUser :exec
DELETE
FROM users
WHERE id = $1;

-- name: ListUser :many
SELECT *
FROM users
ORDER BY id;

-- name: CreateProfile :one
INSERT INTO profile (user_id, name, bio)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetProfileByUserId :one
SELECT *
FROM profile
WHERE user_id = $1
LIMIT 1;

-- name: UpdateProfileByUserId :one
UPDATE profile
SET name = $2,
    bio = $3
WHERE user_id = $1
RETURNING *;

-- name: PartialUpdateProfileByUserId :one
UPDATE profile
SET name = CASE WHEN @update_name::boolean THEN @name::VARCHAR(255) ELSE name END,
    bio = CASE WHEN @update_bio::boolean THEN @bio::TEXT ELSE bio END
WHERE user_id = @user_id
RETURNING *;

-- name: DeleteProfileByUserId :exec
DELETE
FROM profile
WHERE user_id = $1;

-- name: ListProfiles :many
SELECT *
FROM profile;
