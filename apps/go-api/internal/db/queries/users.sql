-- name: GetUserByEmail :one
SELECT id, email, contact_email, is_verified, phone_number, family_name, given_name, is_active, created_at, updated_at
FROM users
WHERE email = $1 AND is_active = true;

-- name: CreateUser :one
INSERT INTO users (email, contact_email, phone_number, family_name, given_name, is_verified, is_active)
VALUES ($1, $2, $3, $4, $5, false, true)
RETURNING id, email, contact_email, is_verified, phone_number, family_name, given_name, is_active, created_at, updated_at;

-- name: UpdateUser :one
UPDATE users 
SET contact_email = $2, phone_number = $3, family_name = $4, given_name = $5, updated_at = NOW()
WHERE email = $1 AND is_active = true
RETURNING id, email, contact_email, is_verified, phone_number, family_name, given_name, is_active, created_at, updated_at;

-- name: CheckUserExists :one
SELECT EXISTS(SELECT 1 FROM users WHERE email = $1 AND is_active = true);