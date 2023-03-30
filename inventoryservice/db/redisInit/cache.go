package redisInit

import "strconv"

var MapName = "emall_stock"

func DeleteCache(productID int64) error {
	cmd := RDB.HDel(MapName, strconv.Itoa(int(productID)))
	return cmd.Err()
}

func SetCache(productID int64, stock int64) (bool, error) {
	cmd := RDB.HSetNX(MapName, strconv.Itoa(int(productID)), stock)
	return cmd.Val(), cmd.Err()
}

func ChangeCache(productID int64, stock int64) error {
	cmd := RDB.HSet(MapName, strconv.Itoa(int(productID)), stock)
	return cmd.Err()
}

func CheckCache(productID int64) (int64, error) {
	rs, err := RDB.HGet(MapName, strconv.Itoa(int(productID))).Result()
	if err != nil {
		return -1, err
	}
	val, err := strconv.Atoi(rs)
	return int64(val), nil

}

func PushToList(productID int64, stock int32) {
	val := strconv.Itoa(int(productID)) + "." + strconv.Itoa(int(stock))
	RDB.RPush("operate_mysql", val)
	return
}

func PopToList() (string, error) {
	cmd, err := RDB.LPop("operate_mysql").Result()
	if err != nil {
		return "", err
	}
	return cmd, nil
}
