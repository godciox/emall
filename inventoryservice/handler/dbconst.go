package handler

import "github.com/go-redis/redis/v7"

const getSku = `-- name: GetSku :one
SELECT stock FROM sku
WHERE sn = ? or id = ? or product_id = ?
`

const adjustInventory = `-- name: AdjustInventory :exec
UPDATE sku SET stock = ? + stock  where id = ? or sn = ? or product_id = ?
`

const createInventoryLog = `-- name: CreateInventoryLog :execresult
INSERT INTO stock_log (
    in_quantity,out_quantity,stock,product_id,creation_date
) VALUES (
     ?, ? , ? , ? , ?
)
`

const createInventory = `-- name: CreateInventory :execresult
INSERT INTO sku (
    sn,stock,price,product_id
) VALUES (
     ?, ? , ? , ?
)
`

var script = redis.NewScript(`
local num = redis.call("HGET", KEYS[1], KEYS[2])
if num == 0 then
    return -1
end

local current = redis.call("HINCRBY", KEYS[1], KEYS[2], -KEYS[3])
if current < 0 then
	redis.call("HINCRBY", KEYS[1], KEYS[2], KEYS[3])
    return -1
else
    return current
end`)
