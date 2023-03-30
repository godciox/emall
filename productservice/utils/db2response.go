package utils

import (
	db "productservice/db/sqlc"
	pb "productservice/proto"
	"strconv"
)

func ProductChange(item db.Product, cur *pb.Product) {

	cur.Id = strconv.Itoa(int(item.ID))
	cur.Name = item.Name
	cur.SeoTitle = item.SeoTitle.String
	cur.SeoKeywords = item.SeoKeywords.String
	cur.SeoDescription = item.SeoDescription.String
	cur.Sn = item.Sn
	cur.Introduction = item.Introduction.String
	cur.Picture = item.Image.String
	cur.ProductCategoryId = item.ProductCategoryID
	cur.Price = item.Price.(float32)
	cur.Score = item.Score.(float32)
}
