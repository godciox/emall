// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.14.0

package db

import (
	"context"
	"database/sql"
)

type Querier interface {
	GetUserByID(ctx context.Context, id int64) (User, error)
	GetUserByName(ctx context.Context, username string) (User, error)
	GetUserByPhone(ctx context.Context, mobile sql.NullString) (User, error)
	GetUserPassword(ctx context.Context, id int64) (string, error)
	GetUserPhoneByID(ctx context.Context, id int64) (sql.NullString, error)
	RegisterUser(ctx context.Context, arg RegisterUserParams) (sql.Result, error)
	UpdateUserInfo(ctx context.Context, arg UpdateUserInfoParams) error
	UpdateUserLoginDate(ctx context.Context, arg UpdateUserLoginDateParams) error
	UpdateUserPassword(ctx context.Context, arg UpdateUserPasswordParams) error
	UpdateUserUpdateDate(ctx context.Context, arg UpdateUserUpdateDateParams) error
}

var _ Querier = (*Queries)(nil)