package handler

import (
	"context"
	"database/sql"
	"fmt"
	cache "orderservice/db/redisInit"
	db "orderservice/db/sqlc"
	pb "orderservice/proto"
	"orderservice/utils"
	"strconv"
	"time"
)

var InventoryService pb.InventoryService

const createOrder = `-- name: CreateOrder :execresult
INSERT INTO orders (
    sn,amount,coupon_code_id,user_id,coupon_discount,consignee,address,phone,expire,type,creation_date
) VALUES (
     ?, ? , ? , ? , ?, ?, ?, ?, ?, ?, ?
);
`
const createOrderItem = `-- name: CreateOrderItem :execresult
INSERT INTO order_item (
    sn,order_id,name,price,weight,thumbnail,quantity,shipped_quantity,return_quantity,sku_id,create_by,creation_date
) VALUES (
    ?, ? , ? , ? , ?, ?, ?, ?, ?, ?, ?, ?
)
`

type OrderService struct {
}

func (o OrderService) SpikePlaceOrder(ctx context.Context, request *pb.SpikePlaceOrderRequest, response *pb.SpikePlaceOrderResponse) error {
	rsp, _ := InventoryService.DecreaseInventoryToSpike(ctx, &pb.DecreaseInventoryToSpikeReq{
		ProductId: request.ProductId,
		Quantity:  request.Quantity,
	})
	if rsp.Status == "500" {
		response.State = "500"
		response.Description = "预扣库存失败" + rsp.Description
		return nil
	}
	if exit, _ := cache.OrderLimit(int(request.UserId), int(request.ProductId)); !exit {
		cache.PushToSet(request.ProductId, int32(request.UserId))
	} else {
		InventoryService.IncreaseInventoryToSpike(ctx, &pb.IncreaseInventoryToSpikeReq{
			ProductId: request.ProductId,
			Quantity:  request.Quantity,
		})
		response.State = "500"
		response.Description = "该秒杀已存在"
		return nil
	}
	cache.PushToList(request.UserId, request.ProductId, int(request.Cost), request.Phone, int(request.Quantity), request.Address, request.Consignee)
	response.State = "100"
	response.Description = "秒杀订单下单成功"
	return nil
}

func (o OrderService) ChangeStateOfOrder(ctx context.Context, request *pb.ChangeOrderStateRequest, response *pb.ChangeOrderStateResponse) error {
	err := db.DBStore.UpdateOrderStatus(ctx, db.UpdateOrderStatusParams{
		Sn:     request.Sn,
		Status: int32(request.OrderState),
	})
	if err != nil {
		response.State = "500"
		response.Description = err.Error()
		return nil
	}
	response.State = "100"
	response.Description = "修改成功"
	return nil
}

func (o OrderService) CheckOrderToUserByDate(ctx context.Context, request *pb.CheckOrderRequest, result *pb.CheckOrderResponse) error {
	timeStartStr, timeEndStr := request.TimeStart.Year+"-"+request.TimeStart.Month+"-"+request.TimeStart.Day, request.TimeEnd.Year+"-"+request.TimeEnd.Month+"-"+request.TimeEnd.Day
	startTime, endTime := utils.Parse_timestr_to_datetime(timeStartStr), utils.Parse_timestr_to_datetime(timeEndStr)
	orderRes, err := db.DBStore.GetOrderByDate(ctx, db.GetOrderByDateParams{CreationDate: sql.NullTime{Time: startTime, Valid: true}, CreationDate_2: sql.NullTime{Time: endTime, Valid: true}, UserID: request.UserId})
	result.ResList = []*pb.OrderResult{}
	if err != nil {
		result.ResList = append(result.ResList, new(pb.OrderResult))
		result.ResList[0].Status = "500"
		result.ResList[0].Descripiton = err.Error()
		return nil
	}
	for _, val := range orderRes {
		curOrderItems, _ := db.DBStore.GetOrderAllItem(ctx, val.ID)
		curRes := new(pb.OrderResult)
		curRes.OrderId = val.ID
		curRes.Sn = val.Sn
		curRes.Status = "100"
		curRes.ShippingAddress = val.Address
		curRes.ShippingCost = 0
		curRes.Items = []*pb.OrderItem{}
		for _, valOfItem := range curOrderItems {
			curItem := new(pb.OrderItem)
			curItem.Item = new(pb.CartItem)
			curItem.Item.ProductId = valOfItem.ID
			curItem.Item.Quantity = int64(valOfItem.Quantity)
			curItem.Cost = int64(valOfItem.Price * valOfItem.Quantity)
			curRes.Items = append(curRes.Items, curItem)
		}
		result.ResList = append(result.ResList, curRes)
	}
	return nil
}

func (o OrderService) CheckOrderToUserByStatus(ctx context.Context, request *pb.CheckOrderRequest, result *pb.CheckOrderResponse) error {
	orderRes, err := db.DBStore.GetOrderByStatus(ctx, db.GetOrderByStatusParams{Status: int32(request.State), UserID: request.UserId})
	result.ResList = []*pb.OrderResult{}
	if err != nil {
		result.ResList = append(result.ResList, new(pb.OrderResult))
		result.ResList[0].Status = "500"
		result.ResList[0].Descripiton = err.Error()
		return nil
	}
	for _, val := range orderRes {
		curOrderItems, _ := db.DBStore.GetOrderAllItem(ctx, val.ID)
		curRes := new(pb.OrderResult)
		curRes.OrderId = val.ID
		curRes.Sn = val.Sn
		curRes.Status = "100"
		curRes.ShippingAddress = val.Address
		curRes.ShippingCost = 0
		curRes.Items = []*pb.OrderItem{}
		for _, valOfItem := range curOrderItems {
			curItem := new(pb.OrderItem)
			curItem.Item = new(pb.CartItem)
			curItem.Item.ProductId = valOfItem.ID
			curItem.Item.Quantity = int64(valOfItem.Quantity)
			curItem.Cost = int64(valOfItem.Price * valOfItem.Quantity)
			curRes.Items = append(curRes.Items, curItem)
		}
		result.ResList = append(result.ResList, curRes)
	}
	return nil
}

func (o OrderService) PlaceOrder(ctx context.Context, request *pb.PlaceOrderRequest, response *pb.PlaceOrderResponse) error {
	var req = new(pb.InventoryRequest)
	req.ProductList = []*pb.InventoryRequestItem{}
	cost := int64(0)
	for _, val := range request.OrderItemId {
		cur := new(pb.InventoryRequestItem)
		cur.Quantity = val.Item.Quantity
		cur.ProductId = val.Item.ProductId
		cost += val.Cost
		req.ProductList = append(req.ProductList, cur)
	}

	req.Tag = true
	// 参数初始化

	var params db.CreateOrderParams
	params.Sn = utils.Generate(time.Now())
	params.Amount = strconv.Itoa(int(cost))
	params.CouponCodeID = request.CouponCodeID
	params.CouponDiscount = fmt.Sprintf("%f", request.CouponDiscount)
	params.Consignee = request.Consignee
	params.Address = request.Address
	params.Phone = request.Phone
	params.UserID = request.UserId
	params.Expire = sql.NullTime{Time: time.Now().Add(30 * time.Minute), Valid: true}
	params.CreationDate = sql.NullTime{Time: time.Now().Add(30 * time.Minute), Valid: true}
	//预扣库存
	rsp, _ := InventoryService.AdjustInventory(ctx, req)
	if rsp != nil && rsp.Status == 0 {
		response.Order = new(pb.OrderResult)
		response.Order.Status = "500"
		response.Order.Descripiton = "预扣库存失败,因为" + rsp.Description
		return nil
	}
	tx, err := db.DB.BeginTx(context.Background(), &sql.TxOptions{
		Isolation: sql.LevelReadCommitted,
		ReadOnly:  false,
	})
	if err != nil {
		response.Order = new(pb.OrderResult)
		response.Order.Status = "500"
		response.Order.Descripiton = "开展事务失败"
		return nil
	}
	//	const createOrder = `-- name: CreateOrder :execresult
	//INSERT INTO orders (
	//    sn,amount,coupon_code_id,user_id,coupon_discount,consignee,address,phone,expire,type,creation_date
	//) VALUES (
	//     ?, ? , ? , ? , ?, ?, ?, ?, ?, ?, ?
	//);
	// 创建order
	res, err := tx.ExecContext(ctx, createOrder,
		params.Sn,
		params.Amount,
		params.CouponCodeID,
		params.UserID,
		params.CouponDiscount,
		params.Consignee,
		params.Address,
		params.Phone,
		params.Expire,
		params.Type,
		params.CreationDate,
	)
	if err != nil {
		tx.Rollback()
		req.Tag = false
		InventoryService.AdjustInventory(ctx, req)
		response.Order = new(pb.OrderResult)
		response.Order.Status = "500"
		response.Order.Descripiton = err.Error()
		return nil
	}
	orderId, _ := res.LastInsertId()
	response.Order = new(pb.OrderResult)
	response.Order.OrderId = orderId
	response.Order.Sn = params.Sn
	response.Order.Status = "0"
	response.Order.ShippingAddress = request.Address
	response.Order.Items = []*pb.OrderItem{}
	for i := range request.OrderItemId {
		cur := new(pb.OrderItem)
		//if item.IsOperateSucceed {
		cur.Cost = request.OrderItemId[i].Cost
		cur.Item = request.OrderItemId[i].Item
		//}
		response.Order.Items = append(response.Order.Items, cur)
		var p db.CreateOrderItemParams
		p.Sn = utils.Generate(time.Now())
		//v := "202303112035181761120003"
		//p.Sn = v
		p.Price = int32(request.OrderItemId[i].Cost / request.OrderItemId[i].Item.Quantity)
		p.OrderID = orderId
		//p.
		p.Quantity = int32(request.OrderItemId[i].Item.Quantity)
		p.CreationDate = sql.NullTime{Time: time.Now(), Valid: true}

		_, err := tx.ExecContext(context.Background(), createOrderItem,
			p.Sn,
			p.OrderID,
			p.Name,
			p.Price,
			p.Weight,
			p.Thumbnail,
			p.Quantity,
			p.ShippedQuantity,
			p.ReturnQuantity,
			p.SkuID,
			p.CreateBy,
			p.CreationDate,
		)
		if err != nil {
			tx.Rollback()
			req.Tag = false
			InventoryService.AdjustInventory(ctx, req)
			response.Order = new(pb.OrderResult)
			response.Order.Status = "500"
			response.Order.Descripiton = err.Error()
			return nil
		}
	}
	cache.RDB.HSet("order_state", strconv.Itoa(int(orderId)), time.Now().Add(30*time.Minute).UnixNano(), 0)
	tx.Commit()
	return nil
}

func (o OrderService) CheckAllOrderToUser(ctx context.Context, request *pb.CheckOrderRequest, result *pb.CheckOrderResponse) error {
	orderRes, err := db.DBStore.GetAllOrder(ctx, request.UserId)
	result.ResList = []*pb.OrderResult{}
	if err != nil {
		result.ResList = append(result.ResList, new(pb.OrderResult))
		result.ResList[0].Status = "500"
		result.ResList[0].Descripiton = err.Error()
		return nil
	}
	for _, val := range orderRes {
		curOrderItems, _ := db.DBStore.GetOrderAllItem(ctx, val.ID)
		curRes := new(pb.OrderResult)
		curRes.OrderId = val.ID
		curRes.Sn = val.Sn
		curRes.Status = "100"
		curRes.ShippingAddress = val.Address
		curRes.ShippingCost = 0
		curRes.Items = []*pb.OrderItem{}
		for _, valOfItem := range curOrderItems {
			curItem := new(pb.OrderItem)
			curItem.Item = new(pb.CartItem)
			curItem.Item.ProductId = valOfItem.ID
			curItem.Item.Quantity = int64(valOfItem.Quantity)
			curItem.Cost = int64(valOfItem.Price * valOfItem.Quantity)
			curRes.Items = append(curRes.Items, curItem)
		}
		result.ResList = append(result.ResList, curRes)
	}
	return nil
}
