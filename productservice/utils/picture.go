package utils

import "strconv"

func GetName(productId, order int64) string {
	return "id" + strconv.Itoa(int(productId)) + "order" + strconv.Itoa(int(order)) + ".png"
}
