// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.14.0

package db

import (
	"database/sql"
)

// 品牌
type Brand struct {
	// 主键_id
	ID int64 `json:"id"`
	// 名称
	Name string `json:"name"`
	// logo
	Logo sql.NullString `json:"logo"`
	// 排序
	Orders sql.NullInt32 `json:"orders"`
	// 创建人
	CreateBy sql.NullString `json:"create_by"`
	// 创建日期
	CreationDate sql.NullTime `json:"creation_date"`
	// 删除标记
	DeleteFlag bool `json:"delete_flag"`
}

// 商品
type Product struct {
	// 主键_id
	ID int64 `json:"id"`
	// 编号
	Sn string `json:"sn"`
	// 名称
	Name string `json:"name"`
	// 展示图片
	Image sql.NullString `json:"image"`
	// 介绍
	Introduction sql.NullString `json:"introduction"`
	// 备注
	Memo sql.NullString `json:"memo"`
	// 搜索关键词
	Keyword sql.NullString `json:"keyword"`
	// 页面标题
	SeoTitle sql.NullString `json:"seo_title"`
	// 页面关键词
	SeoKeywords sql.NullString `json:"seo_keywords"`
	// 页面描述
	SeoDescription sql.NullString `json:"seo_description"`
	// 评分
	Score interface{} `json:"score"`
	// 价格
	Price interface{} `json:"price"`
	// 总评分
	TotalScore int64 `json:"total_score"`
	// 评分数
	ScoreCount int64 `json:"score_count"`
	// 点击数
	Hits int64 `json:"hits"`
	// 周点击数
	WeekHits int64 `json:"week_hits"`
	// 月点击数
	MonthHits int64 `json:"month_hits"`
	// 销量
	Sales int64 `json:"sales"`
	// 周销量
	WeekSales int64 `json:"week_sales"`
	// 月销量
	MonthSales int64 `json:"month_sales"`
	// 品牌
	BrandID sql.NullInt64 `json:"brand_id"`
	// 商品分类
	ProductCategoryID int64 `json:"product_category_id"`
	// 创建人
	CreateBy sql.NullString `json:"create_by"`
	// 创建日期
	CreationDate sql.NullTime `json:"creation_date"`
	// 删除标记
	DeleteFlag bool `json:"delete_flag"`
}

// 商品分类
type ProductCategory struct {
	// 主键_id
	ID int64 `json:"id"`
	// 名称
	Name string `json:"name"`
	// 页面标题
	SeoTitle sql.NullString `json:"seo_title"`
	// 页面关键词
	SeoKeywords sql.NullString `json:"seo_keywords"`
	// 页面描述
	SeoDescription sql.NullString `json:"seo_description"`
	// 树路径
	TreePath string `json:"tree_path"`
	// 层级
	Grade int32 `json:"grade"`
	// 图片
	Image sql.NullString `json:"image"`
	// 上级分类
	ParentID sql.NullInt64 `json:"parent_id"`
	// 创建人
	CreateBy sql.NullString `json:"create_by"`
	// 创建日期
	CreationDate sql.NullTime `json:"creation_date"`
	// 删除标记
	DeleteFlag bool `json:"delete_flag"`
}

type ProductCategoryBrand struct {
	ProductCategories int64 `json:"product_categories"`
	Brands            int64 `json:"brands"`
}

// 商品图片
type ProductImage struct {
	// 主键_id
	ID int64 `json:"id"`
	// 标题
	Title sql.NullString `json:"title"`
	// 商品
	ProductID int64 `json:"product_id"`
	// 原图片
	Source sql.NullString `json:"source"`
	// 缩略图
	Thumbnail sql.NullString `json:"thumbnail"`
	// 排序
	Orders sql.NullInt32 `json:"orders"`
	// 创建人
	CreateBy sql.NullString `json:"create_by"`
	// 创建日期
	CreationDate sql.NullTime `json:"creation_date"`
	// 删除标记
	DeleteFlag bool `json:"delete_flag"`
}