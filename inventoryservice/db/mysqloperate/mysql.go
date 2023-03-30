package mysqloperate

import (
	"context"
	"go-micro.dev/v4/logger"
	"inventoryservice/db/redisInit"
	db "inventoryservice/db/sqlc"
	"strconv"
	"strings"
)

func OperateInventory() {
	for {
		if num, _ := redisInit.RDB.LLen("operate_mysql").Result(); num == 0 {
			continue
		}
		str, err := redisInit.RDB.LPop("operate_mysql").Result()
		if err != nil {
			if err.Error() != "redis: nil" {
				logger.Infof("redis operate inventory failed of %s", err.Error())
			}
		}
		arr := strings.Split(str, ".")
		productID, _ := strconv.Atoi(arr[0])
		stock, _ := strconv.Atoi(arr[1])
		err = db.DBStore.UpdateInventory(context.Background(), db.UpdateInventoryParams{
			Stock:     int32(stock),
			ProductID: int64(productID),
			Sn:        "",
		})
		if err != nil {
			logger.Infof("mysql operate inventory failed of %s", err.Error())
		}
	}
}
