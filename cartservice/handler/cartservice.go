package handler

import (
	db "cartservice/db/sqlc"
	pb "cartservice/proto"
	"context"
	"database/sql"
	"strconv"
	"time"
)

type Cartservice struct {
}

func (c Cartservice) AddItem(ctx context.Context, request *pb.AddItemRequest, response *pb.CartResponse) error {
	rs, err := db.DBStore.GetCart(ctx, sql.NullInt64{
		Int64: request.UserId,
		Valid: true,
	})
	if err != nil {
		if err == sql.ErrNoRows {
			rs, err := db.DBStore.CreateCart(ctx, db.CreateCartParams{
				UserID: sql.NullInt64{
					Int64: request.GetUserId(),
					Valid: true,
				},
				CartKey:  "",
				CreateBy: sql.NullString{},
				CreationDate: sql.NullTime{
					Time:  time.Now(),
					Valid: true,
				},
			})
			if err != nil {
				response.Status = "500"
				response.Response = err.Error()
				return nil
			} else {
				response.CartId, _ = rs.LastInsertId()
			}
		} else {
			response.Status = "500"
			response.Response = err.Error()
			return nil
		}
	} else {
		response.CartId = rs.ID
	}

	_, err = db.DBStore.CreateCartItem(ctx, db.CreateCartItemParams{
		CartID:    response.CartId,
		ProductID: request.Item.ProductId,
		Quantity:  int32(request.Item.Quantity),
		CreateBy: sql.NullString{
			String: strconv.Itoa(int(request.UserId)),
			Valid:  true,
		},
		CreationDate: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
	})
	if err != nil {
		response.Status = "500"
		response.Response = err.Error()
		return nil
	}
	response.Status = "100"
	response.Response = "新增购物车项成功"
	return nil
}

func (c Cartservice) AdjustItem(ctx context.Context, request *pb.AddItemRequest, cart *pb.Cart) error {
	rs, err := db.DBStore.GetCart(ctx, sql.NullInt64{
		Int64: request.UserId,
		Valid: true,
	})
	cart.CartRsp = new(pb.CartResponse)
	if err != nil {
		cart.CartRsp.Status = "500"
		cart.CartRsp.Response = err.Error()
		return nil
	}
	if request.Item.Quantity != 0 {
		err = db.DBStore.UpdateCartItem(ctx, db.UpdateCartItemParams{
			CartID:    rs.ID,
			ProductID: request.Item.ProductId,
			Quantity:  int32(request.Item.Quantity),
		})
	} else {
		err = db.DBStore.DeleteCartItem(ctx, db.DeleteCartItemParams{
			ProductID: request.Item.ProductId,
			CartID:    rs.ID,
		})
	}
	if err != nil {
		cart.CartRsp.Status = "500"
		cart.CartRsp.Response = err.Error()
		return nil
	}
	cart.CartRsp.Status = "100"
	cart.CartRsp.Response = "调整购物车项成功"
	return nil
}

func (c Cartservice) EmptyCart(ctx context.Context, request *pb.EmptyCartRequest, response *pb.CartResponse) error {
	rs, err := db.DBStore.GetCart(ctx, sql.NullInt64{
		Int64: request.UserId,
		Valid: true,
	})
	if err != nil {
		response.Status = "500"
		response.Response = err.Error()
		return nil
	}
	err = db.DBStore.DeleteCart(ctx, sql.NullInt64{
		Int64: request.UserId,
		Valid: true,
	})
	if err != nil {
		response.Status = "500"
		response.Response = err.Error()
		return nil
	}
	err = db.DBStore.DeleteCartAllItem(ctx, rs.ID)
	if err != nil {
		response.Status = "500"
		response.Response = err.Error()
		return nil
	}
	response.Status = "100"
	response.Response = "删除购物车成功"
	return nil
}

func (c Cartservice) GetCart(ctx context.Context, request *pb.GetCartRequest, cart *pb.Cart) error {
	rs, err := db.DBStore.GetCart(ctx, sql.NullInt64{
		Int64: request.UserId,
		Valid: true,
	})
	cart.CartRsp = new(pb.CartResponse)
	if err != nil {
		cart.CartRsp.Status = "500"
		cart.CartRsp.Response = err.Error()
		return nil
	}
	itemRes, err := db.DBStore.GetCartItem(ctx, rs.ID)
	if err != nil {
		cart.CartRsp.Status = "500"
		cart.CartRsp.Response = err.Error()
		return nil
	}
	for _, val := range itemRes {
		cur := new(pb.CartItem)
		cur.ProductId = val.ProductID
		cur.Quantity = int64(val.Quantity)
		cart.Items = append(cart.Items, cur)
	}
	cart.CartRsp.Status = "100"
	cart.CartRsp.Response = "获取购物车项成功"
	return nil
}
