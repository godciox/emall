package redisInit

import (
	"fmt"
	"strconv"
)

func PushToList(userID int64, productID int64, cost int, phone string, quantity int, address, consignee string) {
	val := strconv.Itoa(int(productID)) + "." + strconv.Itoa(int(userID)) + "." + strconv.Itoa(int(cost)) + "." + phone + "." + strconv.Itoa(int(quantity)) + "." + address + "." + consignee
	RDB.RPush("operate_mysql_1", val)
	return
}

func PushToListStr(val string) {
	RDB.RPush("operate_mysql_1", val)
	return
}

func PopToList() (string, error) {
	cmd, err := RDB.LPop("operate_mysql_1").Result()
	if err != nil {
		return "", err
	}
	return cmd, nil
}

var orderLimitKey = "order_limit_key"

func PushToSet(productID int64, userID int32) {
	val := strconv.Itoa(int(productID)) + "." + strconv.Itoa(int(userID))
	RDB.SAdd("operate_mysql", val)
	return
}

func PopToSet(productID int64, userID int32) error {
	val := strconv.Itoa(int(productID)) + "." + strconv.Itoa(int(userID))
	_, err := RDB.SRem("operate_mysql", val).Result()

	return err
}

func OrderLimit(userID int, productID int) (bool, error) {
	userOrder := fmt.Sprintf("%d-%d", userID, productID)
	fmt.Println(RDB)
	return RDB.SIsMember(orderLimitKey, userOrder).Result()
}
