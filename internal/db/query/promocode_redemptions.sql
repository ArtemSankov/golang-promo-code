-- name: CreateRedemption :one
INSERT INTO promocode_redemptions (id, promocode_id, user_id, redeemed_at)
VALUES (gen_random_uuid(), $1, $2, $3)
RETURNING *;

-- name: GetRedemptionByID :one
SELECT id, promocode_id, user_id, redeemed_at
FROM promocode_redemptions
WHERE id = $1;


