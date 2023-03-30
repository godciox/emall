-- name: CreateInventory :execresult
INSERT INTO sku (sn, stock, price, product_id)
VALUES (?, ?, ?, ?);

-- name: CreateInventoryLog :execresult
INSERT INTO stock_log (in_quantity, out_quantity, stock, product_id, creation_date)
VALUES (?, ?, ?, ?, ?);

-- name: UpdateInventory :exec
UPDATE sku
SET stock = sqlc.arg(stock)
where product_id = ?
   or sn = ?;

-- name: AdjustInventory :exec
UPDATE sku
SET stock = sqlc.arg(adujust_num) + stock
where id = ?
   or sn = ?
   or product_id = ?;

-- name: GetSku :one
SELECT stock
FROM sku
WHERE sn = ?
   or product_id = ?;