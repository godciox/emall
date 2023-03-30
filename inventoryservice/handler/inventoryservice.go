package handler

import (
	"context"
	"database/sql"
	cache "inventoryservice/db/redisInit"
	db "inventoryservice/db/sqlc"
	pb "inventoryservice/proto"
	"inventoryservice/utils"
	"strconv"
	"time"
)

type InventoryService struct {
}

func (i InventoryService) MakeProductHot(ctx context.Context, request *pb.MakeProductHotRequest, response *pb.MakeProductHotResponse) error {
	sku, err := db.DBStore.GetSku(ctx, db.GetSkuParams{
		Sn:        "",
		ProductID: request.ProductId,
	})
	if err != nil {
		response.Status = "500"
		response.Description = err.Error()
		return nil
	}
	ok, err := cache.SetCache(request.GetProductId(), int64(sku))
	if !ok {
		response.Status = "500"
		response.Description = "已经存在该key"
		return nil
	}
	if err != nil {
		response.Status = "500"
		response.Description = err.Error()
		return nil
	}
	response.Status = "100"
	response.Description = "热点商品预热完成"
	return nil
}

func (i InventoryService) DecreaseInventoryToSpike(ctx context.Context, req *pb.DecreaseInventoryToSpikeReq, rsp *pb.DecreaseInventoryToSpikeRsp) error {
	cmd := script.Run(cache.RDB, []string{cache.MapName, strconv.Itoa(int(req.ProductId)), strconv.Itoa(int(req.Quantity))})
	rs, err := cmd.Val(), cmd.Err()
	if err != nil || rs == nil {
		rsp.Status = "500"
		rsp.Description = err.Error()
		return nil
	}
	if rs.(int64) == -1 {
		script.Run(cache.RDB, []string{cache.MapName, strconv.Itoa(int(req.ProductId)), strconv.Itoa(int(-req.Quantity))})
		rsp.Status = "500"
		rsp.Description = "库存不足"
		return nil
	}
	rsp.Status = "100"
	rsp.Description = "扣库存成功"
	cache.PushToList(req.ProductId, int32(rs.(int64)))
	return nil
}

func (i InventoryService) IncreaseInventoryToSpike(ctx context.Context, req *pb.IncreaseInventoryToSpikeReq, rsp *pb.IncreaseInventoryToSpikeRsp) error {
	cmd := script.Run(cache.RDB, []string{cache.MapName, strconv.Itoa(int(req.ProductId)), strconv.Itoa(int(-req.Quantity))})
	rs, err := cmd.Val(), cmd.Err()
	if err != nil || rs == nil {
		rsp.Status = "500"
		rsp.Description = err.Error()
		return nil
	}
	if rs.(int64) != -1 {
		rsp.Status = "100"
		rsp.Description = "增加库存成功"
	}
	cache.PushToList(req.ProductId, int32(rs.(int64)))
	return nil
}

func (i InventoryService) AdjustInventory(ctx context.Context, request *pb.InventoryRequest, response *pb.InventoryResponse) error {
	tx, err := db.DB.Begin()
	if err != nil {
		response.Status = 1
		response.Description = err.Error()
		return nil
	}
	for _, val := range request.ProductList {
		if request.Tag {
			val.Quantity = -val.Quantity
		}
		skuNum := int64(0)
		if val.Quantity < 0 {
			rs, err := tx.Query(getSku, "", 0, val.ProductId)
			if err != nil {
				tx.Rollback()
				response.Status = 1
				response.Description = err.Error()
				return nil
			}

			for rs.Next() {
				rs.Scan(&skuNum)
			}
			if skuNum < (-val.Quantity) {
				tx.Rollback()
				response.Status = 1
				response.Description = strconv.Itoa(int(val.ProductId)) + "库存不够了"
				return nil
			}
			_, err = tx.Exec(createInventoryLog, 0, -val.Quantity, skuNum+val.Quantity, val.ProductId, sql.NullTime{Time: time.Now()})
			if err != nil {
				tx.Rollback()
				response.Status = 1
				response.Description = "stock—log出错：" + err.Error()
				return nil
			}
		} else {
			_, err = tx.Exec(createInventoryLog, val.Quantity, 0, skuNum+val.Quantity, val.ProductId, sql.NullTime{Time: time.Now()})
			if err != nil {
				tx.Rollback()
				response.Status = 1
				response.Description = "stock—log出错：" + err.Error()
				return nil
			}
		}

		_, err := tx.Exec(adjustInventory, val.Quantity, 0, "", val.ProductId)
		if err != nil {
			tx.Rollback()
			response.Status = 1
			response.Description = err.Error()
			return nil
		}
	}
	tx.Commit()
	response.Status = 0
	response.Description = "操作库存成功"
	return nil
}

func (i InventoryService) InsertInventory(ctx context.Context, request *pb.InsertInventoryRequest, response *pb.InsertInventoryResponse) error {
	tx, err := db.DB.Begin()
	if err != nil {
		response.Status = "500"
		response.Description = err.Error()
		return nil
	}

	_, err = tx.Exec(createInventoryLog, request.Stock, 0, request.Stock, request.ProductId, sql.NullTime{Time: time.Now()})
	if err != nil {
		tx.Rollback()
		response.Status = "500"
		response.Description = "stock—log出错：" + err.Error()
		return nil
	}

	_, err = tx.Exec(createInventory, utils.Generate(time.Now()), request.Stock, request.Price, request.ProductId)
	if err != nil {
		tx.Rollback()
		response.Status = "500"
		response.Description = err.Error()
		return nil
	}
	if err != nil {
		if err.Error() != "redis: nil" {
			tx.Rollback()
			response.Status = "500"
			response.Description = err.Error()
			return nil
		}
	}
	tx.Commit()
	response.Status = "100"
	response.Description = "操作库存成功"
	return nil
}
