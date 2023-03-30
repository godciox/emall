package handler

const createProduct = `-- name: CreateProduct :execresult
INSERT INTO product (
    sn, name, image, introduction, memo, keyword, seo_title, seo_keywords, seo_description, brand_id, product_category_id, create_by, creation_date, price
) VALUES (
    ?, ? , ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?
)
`

const updateProduct = `-- name: UpdateProduct :exec
UPDATE product
SET name = ?,
    image = ?,
    introduction = ?,
    memo = ?,
    keyword = ?,
    seo_title = ?,
    seo_description = ?,
    brand_id = ?,
    product_category_id = ?,
	price = ?
where id = ?
   or sn = ?
   or name = ?
`

const createProductCategory = `-- name: CreateProductCategory :execresult
INSERT INTO product_category (name, seo_title, seo_keywords, seo_description, tree_path, grade, image, parent_id,
                              create_by, creation_date)
VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`

const updateProductCategory = `-- name: UpdateProductCategory :exec
UPDATE product_category
SET name = ?,
image = ?,
seo_title = ?,
seo_keywords = ?,
seo_description = ?,
tree_path = ?,
grade = ?,
parent_id = ?
where id = ?
or name = ?
`

const insertProductImage = `-- name: InsertProductImage :execresult
INSERT INTO product_image (title, product_id, source, thumbnail, create_by, creation_date)
VALUES (?, ?, ?, ?, ?, ?)
`

const updateProductImage = `-- name: UpdateProductImage :exec
UPDATE product_image
SET source = ?,
    thumbnail = ?
where id = ?
   or title = ?
   or product_id = ?
`
const createBrand = `-- name: CreateBrand :execresult
INSERT INTO brand (name, logo, create_by, creation_date)
VALUES (?, ?, ?, ?)
`

const createProductCategoryBrand = `-- name: CreateProductCategoryBrand :execresult
INSERT INTO product_category_brand (product_categories, brands)
VALUES (?, ?)
`

const updateBrand = `-- name: UpdateBrand :exec
UPDATE brand
SET name = ?,
    logo = ?
where id = ?
`
