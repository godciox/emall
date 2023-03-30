-- name: CreateOrder :execresult
INSERT INTO orders (sn, amount, coupon_code_id, user_id, coupon_discount, consignee, address, phone, expire, type,
                    creation_date)
VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);

-- name: CreateOrderItem :execresult
INSERT INTO order_item (sn, order_id, name, price, weight, thumbnail, quantity, shipped_quantity, return_quantity,
                        sku_id, create_by, creation_date)
VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);

-- name: GetAllOrder :many
SELECT id,
       sn,
       amount,
       status,
       coupon_code_id,
       coupon_discount,
       consignee,
       address
FROM orders
WHERE user_id = ?;

-- name: GetOrderAllItem :many
SELECT id,
       sn,
       name,
       thumbnail,
       quantity,
       shipped_quantity,
       return_quantity,
       price
FROM order_item
WHERE order_id = ?;

-- name: GetOrderByDate :many
SELECT id,
       sn,
       amount,
       status,
       coupon_code_id,
       coupon_discount,
       consignee,
       address
FROM orders
WHERE user_id = ?
  and creation_date between ? and ?;

-- name: GetOrderByStatus :many
SELECT id,
       sn,
       amount,
       status,
       coupon_code_id,
       coupon_discount,
       consignee,
       address
FROM orders
WHERE user_id = ?
  and status = ?;

-- name: UpdateOrderStatus :exec
UPDATE orders
SET status = sqlc.arg(status)
where sn = ?;

-- name: DeleteOrder :exec
DELETE FROM orders
where id = ?;