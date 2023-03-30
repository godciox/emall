package mysqloperate

import (
	"context"
	"database/sql"
	"go-micro.dev/v4/logger"
	"orderservice/db/redisInit"
	db "orderservice/db/sqlc"
	"orderservice/utils"
	"strconv"
	"strings"
	"time"
)

func OperateInventory() {
	for {
		if num, _ := redisInit.RDB.LLen("operate_mysql_1").Result(); num == 0 {
			continue
		}
		str, err := redisInit.RDB.LPop("operate_mysql_1").Result()
		if err != nil {
			if err.Error() != "redis: nil" {
				logger.Infof("redis operate inventory failed of %s", err.Error())
			}
		}
		arr := strings.Split(str, ".")
		userID, _ := strconv.Atoi(arr[0])
		productID, _ := strconv.Atoi(arr[1])
		cost, _ := strconv.Atoi(arr[2])
		phone := arr[3]
		quantity, _ := strconv.Atoi(arr[4])
		address := arr[5]
		consignee := arr[6]
		rs, err := db.DBStore.CreateOrder(context.Background(), db.CreateOrderParams{
			Sn:             utils.Generate(time.Now()),
			Amount:         strconv.Itoa(quantity),
			CouponCodeID:   "",
			UserID:         int64(userID),
			CouponDiscount: "",
			Consignee:      consignee,
			Address:        address,
			Phone:          phone,
			Expire: sql.NullTime{
				Time:  time.Now().Add(30 * time.Minute),
				Valid: true,
			},
			Type: 0,
			CreationDate: sql.NullTime{
				Time:  time.Now(),
				Valid: true,
			},
		})
		if err != nil {
			redisInit.PushToListStr(str)
			continue
		}
		orderID, _ := rs.LastInsertId()
		_, err = db.DBStore.CreateOrderItem(context.Background(), db.CreateOrderItemParams{
			Sn:              utils.Generate(time.Now()),
			OrderID:         orderID,
			Name:            "",
			Price:           int32(cost / quantity),
			Weight:          sql.NullInt32{},
			Thumbnail:       sql.NullString{},
			Quantity:        int32(quantity),
			ShippedQuantity: 0,
			ReturnQuantity:  0,
			SkuID: sql.NullInt64{
				Int64: int64(productID),
				Valid: true,
			},
			CreateBy:     sql.NullString{},
			CreationDate: sql.NullTime{},
		})
		if err != nil {
			db.DBStore.DeleteOrder(context.Background(), orderID)
			redisInit.PushToListStr(str)
			continue
		}
	}
}
