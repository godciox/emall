// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.14.0
// source: inventory.sql

package db

import (
	"context"
	"database/sql"
)

const adjustInventory = `-- name: AdjustInventory :exec
UPDATE sku
SET stock = ? + stock
where id = ?
   or sn = ?
   or product_id = ?
`

type AdjustInventoryParams struct {
	AdujustNum interface{} `json:"adujust_num"`
	ID         int64       `json:"id"`
	Sn         string      `json:"sn"`
	ProductID  int64       `json:"product_id"`
}

func (q *Queries) AdjustInventory(ctx context.Context, arg AdjustInventoryParams) error {
	_, err := q.db.ExecContext(ctx, adjustInventory,
		arg.AdujustNum,
		arg.ID,
		arg.Sn,
		arg.ProductID,
	)
	return err
}

const createInventory = `-- name: CreateInventory :execresult
INSERT INTO sku (sn, stock, price, product_id)
VALUES (?, ?, ?, ?)
`

type CreateInventoryParams struct {
	Sn        string `json:"sn"`
	Stock     int32  `json:"stock"`
	Price     string `json:"price"`
	ProductID int64  `json:"product_id"`
}

func (q *Queries) CreateInventory(ctx context.Context, arg CreateInventoryParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, createInventory,
		arg.Sn,
		arg.Stock,
		arg.Price,
		arg.ProductID,
	)
}

const createInventoryLog = `-- name: CreateInventoryLog :execresult
INSERT INTO stock_log (in_quantity, out_quantity, stock, product_id, creation_date)
VALUES (?, ?, ?, ?, ?)
`

type CreateInventoryLogParams struct {
	InQuantity   int32        `json:"in_quantity"`
	OutQuantity  int32        `json:"out_quantity"`
	Stock        int32        `json:"stock"`
	ProductID    int64        `json:"product_id"`
	CreationDate sql.NullTime `json:"creation_date"`
}

func (q *Queries) CreateInventoryLog(ctx context.Context, arg CreateInventoryLogParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, createInventoryLog,
		arg.InQuantity,
		arg.OutQuantity,
		arg.Stock,
		arg.ProductID,
		arg.CreationDate,
	)
}

const getSku = `-- name: GetSku :one
SELECT stock
FROM sku
WHERE sn = ?
   or product_id = ?
`

type GetSkuParams struct {
	Sn        string `json:"sn"`
	ProductID int64  `json:"product_id"`
}

func (q *Queries) GetSku(ctx context.Context, arg GetSkuParams) (int32, error) {
	row := q.db.QueryRowContext(ctx, getSku, arg.Sn, arg.ProductID)
	var stock int32
	err := row.Scan(&stock)
	return stock, err
}

const updateInventory = `-- name: UpdateInventory :exec
UPDATE sku
SET stock = ?
where product_id = ?
   or sn = ?
`

type UpdateInventoryParams struct {
	Stock     int32  `json:"stock"`
	ProductID int64  `json:"product_id"`
	Sn        string `json:"sn"`
}

func (q *Queries) UpdateInventory(ctx context.Context, arg UpdateInventoryParams) error {
	_, err := q.db.ExecContext(ctx, updateInventory, arg.Stock, arg.ProductID, arg.Sn)
	return err
}