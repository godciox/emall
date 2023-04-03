package handler

import (
	"context"
	"database/sql"
	"io/ioutil"
	db "productservice/db/sqlc"
	pb "productservice/proto"
	"productservice/utils"
	"time"
)

type ProductService struct {
}

func (p ProductService) OperateProducts(ctx context.Context, request *pb.OperateProductsRequest, response *pb.OperateProductsResponse) error {
	tx, err := db.DB.Begin()
	if err != nil {
		response.Status = "500"
		response.Description = err.Error() + "OperateProducts"
		return nil
	}
	for _, val := range request.Products {
		_, err := tx.Exec(updateProduct, val.Name, val.Picture, val.Introduction, "", val.SeoKeywords, val.SeoTitle, val.SeoDescription, val.BrandId, val.ProductCategoryId, val.Price, val.Id, val.Sn, val.Name)
		if err != nil {
			tx.Rollback()
			response.Status = "500"
			response.Description = err.Error() + "OperateProducts"
			return nil
		}
	}
	tx.Commit()
	response.Status = "100"
	response.Description = "批量修改商品成功"
	return nil
}

func (p ProductService) InsertProducts(ctx context.Context, request *pb.InsertProductsRequest, response *pb.InsertProductsResponse) error {
	tx, err := db.DB.Begin()
	if err != nil {
		response.Status = "500"
		response.Description = err.Error() + "InsertProducts"
		return nil
	}
	for _, val := range request.Products {
		_, err := tx.Exec(createProduct, utils.Generate(time.Now()), val.Name, val.Picture, val.Introduction, "", "", val.SeoTitle, val.SeoKeywords, val.SeoDescription, val.BrandId, val.ProductCategoryId, request.Creator, sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		}, val.Price)
		if err != nil {
			tx.Rollback()
			response.Status = "500"
			response.Description = err.Error() + "InsertProducts"
			return nil
		}
	}
	tx.Commit()
	response.Status = "100"
	response.Description = "批量插入商品成功"
	return nil
}

func (p ProductService) InsertCategory(ctx context.Context, request *pb.InsertCategoryRequest, response *pb.InsertCategoryResponse) error {
	tx, err := db.DB.Begin()
	if err != nil {
		response.Status = "500" + "InsertCategory"
		response.Description = err.Error()
		return nil
	}
	for _, val := range request.Categorys {
		_, err := tx.Exec(createProductCategory, val.Name, val.SeoTitle, val.SeoKeywords, val.SeoDescription, val.TreePath, val.Grade, val.Image, val.ParentId, request.Creator, sql.NullTime{Time: time.Now(), Valid: true})
		if err != nil {
			tx.Rollback()
			response.Status = "500"
			response.Description = err.Error() + "InsertCategory"
			return nil
		}
	}
	tx.Commit()
	response.Status = "100"
	response.Description = "批量插入分类成功"
	return nil
}

func (p ProductService) GetCategory(ctx context.Context, request *pb.CategoryRequest, response *pb.CategoryResponse) error {
	categorys, err := db.DBStore.CheckCategory(ctx, db.CheckCategoryParams{Name: request.Name, ID: request.Id, ParentID: sql.NullInt64{Int64: request.ParentId, Valid: true}, Grade: request.Grade, Limit: 20, Offset: 20 * (request.Page - 1)})
	if err != nil {
		response.Status = "500"
		response.Description = err.Error() + "GetCategory"
		return nil
	}
	if len(categorys) == 0 {
		response.Status = "500"
		response.Description = "无该分类"
		return nil
	}
	response.Categorys = []*pb.Category{}
	for _, item := range categorys {
		cur := new(pb.Category)
		cur.Name = item.Name
		cur.ParentId = item.ParentID.Int64
		cur.Grade = item.Grade
		cur.SeoTitle = item.SeoTitle.String
		cur.TreePath = item.TreePath
		cur.Image = item.Image.String
		cur.SeoDescription = item.SeoDescription.String
		cur.SeoKeywords = item.SeoKeywords.String
		response.Categorys = append(response.GetCategorys(), cur)
	}
	response.Status = "100"
	response.Description = "GetCategory成功"
	return nil
}

func (p ProductService) DeleteCategory(ctx context.Context, request *pb.CategoryRequest, response *pb.CategoryResponse) error {
	err := db.DBStore.DeleteCategory(ctx, db.DeleteCategoryParams{ID: request.Id, ParentID: sql.NullInt64{Int64: request.ParentId, Valid: true}, Name: request.Name, Grade: request.Grade})
	if err != nil {
		response.Status = "500"
		response.Description = err.Error() + "DeleteCategory出错"
		return nil
	}
	response.Status = "100"
	response.Description = "删除分类成功"
	return nil
}

func (p ProductService) ChangeCategory(ctx context.Context, request *pb.InsertCategoryRequest, response *pb.InsertCategoryResponse) error {
	tx, err := db.DB.Begin()
	if err != nil {
		response.Status = "500"
		response.Description = err.Error() + "OperateProducts"
		return nil
	}
	for _, val := range request.Categorys {
		_, err := tx.Exec(updateProductCategory, val.Name, val.Image, val.SeoTitle, val.SeoKeywords, val.SeoDescription, val.TreePath, val.Grade, val.ParentId, val.Id, "")
		if err != nil {
			tx.Rollback()
			response.Status = "500"
			response.Description = err.Error() + "OperateProducts"
			return nil
		}
	}
	tx.Commit()
	response.Status = "100"
	response.Description = "批量修改分类成功"
	return nil
}

func (p ProductService) GetProductImage(ctx context.Context, request *pb.GetImageRequest, response *pb.GetImageResponse) error {
	imgs, err := db.DBStore.CheckProductImage(ctx, db.CheckProductImageParams{
		ProductID: request.ProductId,
	})
	response.Rsp = new(pb.ImageResponse)
	if err != nil {
		response.Rsp.Status = "500"
		response.Rsp.Description = err.Error() + "GetProductImage"
		return nil
	}
	if len(imgs) == 0 {
		response.Rsp.Status = "500"
		response.Rsp.Description = "无该图片"
		return nil
	}
	response.Images = []*pb.ProductImage{}
	for _, item := range imgs {
		cur := new(pb.ProductImage)
		cur.ProductId = request.ProductId
		cur.SourceImg = item.Source.String
		cur.Title = item.Title.String
		response.Images = append(response.Images, cur)
	}
	response.Rsp.Status = "100"
	response.Rsp.Description = "获取商品图片成功"
	return nil
}

func (p ProductService) InsertProductImage(ctx context.Context, request *pb.ImageRequest, response *pb.ImageResponse) error {
	tx, err := db.DB.Begin()
	if err != nil {
		response.Status = "500" + "InsertProductImage"
		response.Description = err.Error()
		return nil
	}
	for _, val := range request.Images {
		_, err := tx.Exec(insertProductImage, val.Title, val.ProductId, val.SourceImg, val.Thumbnail, request.Creator, sql.NullTime{Time: time.Now(), Valid: true})
		if err != nil {
			tx.Rollback()
			response.Status = "500"
			response.Description = err.Error() + "InsertProductImage"
			return nil
		}
	}
	tx.Commit()
	response.Status = "100"
	response.Description = "批量插入商品图片成功"
	return nil
}

func (p ProductService) ChangeProductImage(ctx context.Context, request *pb.ImageRequest, response *pb.ImageResponse) error {
	tx, err := db.DB.Begin()
	if err != nil {
		response.Status = "500"
		response.Description = err.Error() + "ChangeProductImage"
		return nil
	}
	for _, val := range request.Images {
		_, err := tx.Exec(updateProductImage, val.SourceImg, val.Thumbnail, val.Id, val.Title, val.ProductId)
		if err != nil {
			tx.Rollback()
			response.Status = "500"
			response.Description = err.Error() + "ChangeProductImage"
			return nil
		}
	}
	tx.Commit()
	response.Status = "100"
	response.Description = "批量修改商品图片成功"
	return nil
}

func (p ProductService) DeleteProductImage(ctx context.Context, request *pb.DeleteImageRequest, response *pb.ImageResponse) error {
	err := db.DBStore.DeleteProductImg(ctx, db.DeleteProductImgParams{ID: 0, ProductID: request.ProductId, Title: sql.NullString{
		String: request.Title,
		Valid:  true,
	}})
	if err != nil {
		response.Status = "500"
		response.Description = err.Error() + "DeleteProductImage出错"
		return nil
	}
	response.Status = "100"
	response.Description = "删除商品图片成功"
	return nil
}

func (p ProductService) GetBrand(ctx context.Context, request *pb.GetBrandRequest, response *pb.GetBrandResponse) error {
	brand, err := db.DBStore.CheckBrand(ctx, request.BrandId)
	response.Rsp = new(pb.BrandResponse)
	if err != nil {
		response.Rsp.Status = "500"
		response.Rsp.Description = err.Error() + "GetBrand"
		return nil
	}
	response.BrandInfo = []*pb.Brand{}
	cur := new(pb.Brand)
	cur.LogoImgPath = brand.Logo.String
	cur.Name = brand.Name
	cur.Id = request.BrandId
	response.BrandInfo = append(response.BrandInfo, cur)
	response.Rsp.Status = "100"
	response.Rsp.Description = "获取商品图片成功"
	return nil
}

func (p ProductService) InsertBrand(ctx context.Context, request *pb.BrandRequest, response *pb.BrandResponse) error {
	tx, err := db.DB.Begin()
	if err != nil {
		response.Status = "500" + "InsertBrand"
		response.Description = err.Error()
		return nil
	}
	for _, val := range request.Brands {
		rs, err := tx.Exec(createBrand, val.Name, val.LogoImgPath, request.Creator, sql.NullTime{Time: time.Now(), Valid: true})
		if err != nil {
			tx.Rollback()
			response.Status = "500"
			response.Description = err.Error() + "InsertBrand"
			return nil
		}
		id, err := rs.LastInsertId()
		if err != nil {
			tx.Rollback()
			response.Status = "500"
			response.Description = err.Error() + "InsertBrand"
			return nil
		}
		_, err = tx.Exec(createProductCategoryBrand, val.ProductCategoriesId, id)
		if err != nil {
			tx.Rollback()
			response.Status = "500"
			response.Description = err.Error() + "InsertBrand"
			return nil
		}
	}
	tx.Commit()
	response.Status = "100"
	response.Description = "批量插入品牌成功"
	return nil
}

func (p ProductService) ChangeBrand(ctx context.Context, request *pb.BrandRequest, response *pb.BrandResponse) error {
	tx, err := db.DB.Begin()
	if err != nil {
		response.Status = "500"
		response.Description = err.Error() + "ChangeBrand"
		return nil
	}
	for _, val := range request.Brands {
		_, err := tx.Exec(updateBrand, val.Name, val.LogoImgPath, val.Id)
		if err != nil {
			tx.Rollback()
			response.Status = "500"
			response.Description = err.Error() + "ChangeBrand"
			return nil
		}
	}
	tx.Commit()
	response.Status = "100"
	response.Description = "批量修改品牌成功"
	return nil
}

func (p ProductService) DeleteBrand(ctx context.Context, request *pb.DeleteBrandRequest, response *pb.BrandResponse) error {
	err := db.DBStore.DeleteBrand(ctx, db.DeleteBrandParams{ID: request.Id, Name: request.Name})
	if err != nil {
		response.Status = "500"
		response.Description = err.Error() + "DeleteBrand出错"
		return nil
	}
	response.Status = "100"
	response.Description = "删除品牌成功"
	return nil
}

func (p ProductService) GetProduct(ctx context.Context, request *pb.GetProductRequest, response *pb.GetProductResponse) error {
	rs, err := db.DBStore.GetProduct(ctx, request.GetId())
	if err != nil {
		response.Status = "500"
		response.Description = err.Error()
		return nil
	}
	if len(rs) == 0 {
		response.Status = "500"
		response.Description = "无该商品"
		return nil
	}
	response.Product = new(pb.Product)
	utils.ProductChange(rs[0], response.Product)
	response.Status = "100"
	response.Description = "操作成功"
	return nil
}

func (p ProductService) SearchProducts(ctx context.Context, request *pb.SearchProductsRequest, response *pb.SearchProductsResponse) error {
	rs, err := db.DBStore.SearchProducts(ctx, db.SearchProductsParams{
		Name:   "%" + request.Query + "%",
		Limit:  20,
		Offset: 20 * (request.Page - 1),
	})
	if err != nil {
		response.Status = "500"
		response.Description = err.Error() + "SearchProducts"
		return nil
	}

	if len(rs) == 0 {
		response.Status = "500"
		response.Description = "无该商品"
		return nil
	}
	response.Results = []*pb.Product{}
	for _, item := range rs {
		cur := new(pb.Product)
		utils.ProductChange(item, cur)
		fileBytes, err := ioutil.ReadFile("./img/" + item.Image.String)
		response.Results = append(response.Results, cur)
		if err != nil {
			response.Imgs = append(response.Imgs, fileBytes)
		} else {
			response.Imgs = append(response.Imgs, []byte{})
		}
	}
	response.Status = "100"
	response.Description = "SearchProducts成功"
	return nil
}

func (p ProductService) ListProductsByBrand(ctx context.Context, request *pb.ListProductsRequest, response *pb.ListProductsResponse) error {
	products, err := db.DBStore.ListProductsByBrand(ctx, db.ListProductsByBrandParams{
		BrandID: sql.NullInt64{Int64: request.GetBrandId(), Valid: true},
		Limit:   20,
		Offset:  20 * (request.Page - 1),
	})
	if err != nil {
		response.Status = "500"
		response.Description = err.Error() + "ListProductsByBrand"
		return nil
	}

	if len(products) == 0 {
		response.Status = "500"
		response.Description = "无该商品"
		return nil
	}
	response.Products = []*pb.Product{}
	for _, item := range products {
		cur := new(pb.Product)
		cur.BrandId = request.BrandId
		utils.ProductChange(item, cur)
		response.Products = append(response.Products, cur)
	}
	response.Status = "100"
	response.Description = "ListProductsByBrand成功"
	return nil
}

func (p ProductService) ListProductsByProductCategory(ctx context.Context, request *pb.ListProductsRequest, response *pb.ListProductsResponse) error {
	products, err := db.DBStore.ListProductsByProductCategory(ctx, db.ListProductsByProductCategoryParams{
		ProductCategoryID: request.CategoryId,
		Limit:             20,
		Offset:            20 * (request.Page - 1),
	})
	if err != nil {
		response.Status = "500"
		response.Description = err.Error() + "ListProductsByProductCategory"
		return nil
	}
	if len(products) == 0 {
		response.Status = "500"
		response.Description = "无该商品"
		return nil
	}
	response.Products = []*pb.Product{}
	for _, item := range products {
		cur := new(pb.Product)
		cur.BrandId = request.BrandId
		utils.ProductChange(item, cur)
		response.Products = append(response.Products, cur)
	}
	response.Status = "100"
	response.Description = "ListProductsByProductCategory成功"
	return nil
}
