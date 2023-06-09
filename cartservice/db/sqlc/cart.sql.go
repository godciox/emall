// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.14.0
// source: cart.sql

package db

import (
	"context"
	"database/sql"
)

const createCart = `-- name: CreateCart :execresult
INSERT INTO cart (user_id, cart_key, create_by, creation_date)
VALUES (?, ?, ?, ?)
`

type CreateCartParams struct {
	UserID       sql.NullInt64  `json:"user_id"`
	CartKey      string         `json:"cart_key"`
	CreateBy     sql.NullString `json:"create_by"`
	CreationDate sql.NullTime   `json:"creation_date"`
}

func (q *Queries) CreateCart(ctx context.Context, arg CreateCartParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, createCart,
		arg.UserID,
		arg.CartKey,
		arg.CreateBy,
		arg.CreationDate,
	)
}

const createCartItem = `-- name: CreateCartItem :execresult
INSERT INTO cart_item (cart_id, product_id, quantity, create_by, creation_date)
VALUES (?, ?, ?, ?, ?)
`

type CreateCartItemParams struct {
	CartID       int64          `json:"cart_id"`
	ProductID    int64          `json:"product_id"`
	Quantity     int32          `json:"quantity"`
	CreateBy     sql.NullString `json:"create_by"`
	CreationDate sql.NullTime   `json:"creation_date"`
}

func (q *Queries) CreateCartItem(ctx context.Context, arg CreateCartItemParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, createCartItem,
		arg.CartID,
		arg.ProductID,
		arg.Quantity,
		arg.CreateBy,
		arg.CreationDate,
	)
}

const deleteCart = `-- name: DeleteCart :exec
DELETE FROM cart
where user_id = ?
`

func (q *Queries) DeleteCart(ctx context.Context, userID sql.NullInt64) error {
	_, err := q.db.ExecContext(ctx, deleteCart, userID)
	return err
}

const deleteCartAllItem = `-- name: DeleteCartAllItem :exec
DELETE FROM cart_item
where cart_id = ?
`

func (q *Queries) DeleteCartAllItem(ctx context.Context, cartID int64) error {
	_, err := q.db.ExecContext(ctx, deleteCartAllItem, cartID)
	return err
}

const deleteCartItem = `-- name: DeleteCartItem :exec
DELETE FROM cart_item
where product_id = ?
   and cart_id = ?
`

type DeleteCartItemParams struct {
	ProductID int64 `json:"product_id"`
	CartID    int64 `json:"cart_id"`
}

func (q *Queries) DeleteCartItem(ctx context.Context, arg DeleteCartItemParams) error {
	_, err := q.db.ExecContext(ctx, deleteCartItem, arg.ProductID, arg.CartID)
	return err
}

const getCart = `-- name: GetCart :one
SELECT id, user_id, cart_key, create_by, creation_date, delete_flag
from cart
where user_id = ?
`

func (q *Queries) GetCart(ctx context.Context, userID sql.NullInt64) (Cart, error) {
	row := q.db.QueryRowContext(ctx, getCart, userID)
	var i Cart
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.CartKey,
		&i.CreateBy,
		&i.CreationDate,
		&i.DeleteFlag,
	)
	return i, err
}

const getCartItem = `-- name: GetCartItem :many
SELECT id, cart_id, product_id, quantity, create_by, creation_date, delete_flag
from cart_item
where cart_id = ?
`

func (q *Queries) GetCartItem(ctx context.Context, cartID int64) ([]CartItem, error) {
	rows, err := q.db.QueryContext(ctx, getCartItem, cartID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []CartItem{}
	for rows.Next() {
		var i CartItem
		if err := rows.Scan(
			&i.ID,
			&i.CartID,
			&i.ProductID,
			&i.Quantity,
			&i.CreateBy,
			&i.CreationDate,
			&i.DeleteFlag,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateCart = `-- name: UpdateCart :exec
UPDATE cart
SET user_id = ?
where id = ?
   or user_id = ?
`

type UpdateCartParams struct {
	UserID   sql.NullInt64 `json:"user_id"`
	ID       int64         `json:"id"`
	UserID_2 sql.NullInt64 `json:"user_id_2"`
}

func (q *Queries) UpdateCart(ctx context.Context, arg UpdateCartParams) error {
	_, err := q.db.ExecContext(ctx, updateCart, arg.UserID, arg.ID, arg.UserID_2)
	return err
}

const updateCartItem = `-- name: UpdateCartItem :exec
UPDATE cart_item
SET quantity = ?
where product_id = ?
and cart_id = ?
`

type UpdateCartItemParams struct {
	Quantity  int32 `json:"quantity"`
	ProductID int64 `json:"product_id"`
	CartID    int64 `json:"cart_id"`
}

func (q *Queries) UpdateCartItem(ctx context.Context, arg UpdateCartItemParams) error {
	_, err := q.db.ExecContext(ctx, updateCartItem, arg.Quantity, arg.ProductID, arg.CartID)
	return err
}
