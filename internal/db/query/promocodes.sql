-- name: CreatePromoCode :one
INSERT INTO promocodes (id, code, discount_type, discount_value, max_activations, activations_count, expires_at)
VALUES (gen_random_uuid(), $1, $2, $3, $4, 0, $5)
RETURNING id;

-- name: GetPromoCodeByCode :one
SELECT id, code, discount_type, discount_value, max_activations, activations_count, expires_at, created_at
FROM promocodes
WHERE code = $1;

-- name: IncrementActivationsCount :exec
UPDATE promocodes
SET activations_count = activations_count + 1
WHERE id = $1 AND activations_count < max_activations;
