-- name: UpsertUserByEmail :one
INSERT INTO users (email, contact_email, phone_number, family_name, given_name, updated_at)
VALUES ($1, $2, $3, $4, $5, NOW())
ON CONFLICT (email) DO UPDATE SET
    contact_email = EXCLUDED.contact_email,
    phone_number = EXCLUDED.phone_number,
    family_name = EXCLUDED.family_name,
    given_name = EXCLUDED.given_name,
    updated_at = NOW()
RETURNING *;

-- name: GetUserByEmail :one
SELECT * FROM users WHERE email = $1;