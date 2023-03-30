-- name: CreateCart :execresult
INSERT INTO cart (user_id, cart_key, create_by, creation_date)
VALUES (?, ?, ?, ?);

-- name: CreateCartItem :execresult
INSERT INTO cart_item (cart_id, product_id, quantity, create_by, creation_date)
VALUES (?, ?, ?, ?, ?);

-- name: GetCart :one
SELECT *
from cart
where user_id = ?;

-- name: GetCartItem :many
SELECT *
from cart_item
where cart_id = ?;

-- name: UpdateCart :exec
UPDATE cart
SET user_id = sqlc.arg(user_id)
where id = ?
   or user_id = ?;

-- name: UpdateCartItem :exec
UPDATE cart_item
SET quantity = sqlc.arg(quantity)
where product_id = ?
and cart_id = ?;

-- name: DeleteCartItem :exec
DELETE FROM cart_item
where product_id = ?
   and cart_id = ?
;

-- name: DeleteCartAllItem :exec
DELETE FROM cart_item
where cart_id = ?
;

-- name: DeleteCart :exec
DELETE FROM cart
where user_id = ?
;