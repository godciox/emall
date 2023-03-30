-- name: CreateBrand :execresult
INSERT INTO brand (name, logo, create_by, creation_date)
VALUES (?, ?, ?, ?);

-- name: CreateProduct :execresult
INSERT INTO product (sn, name, image, introduction, memo, keyword, seo_title, seo_keywords, seo_description, brand_id,
                     product_category_id, create_by, creation_date, price)
VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);

-- name: CreateProductCategory :execresult
INSERT INTO product_category (name, seo_title, seo_keywords, seo_description, tree_path, grade, image, parent_id,
                              create_by, creation_date)
VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?);

-- name: CreateProductCategoryBrand :execresult
INSERT INTO product_category_brand (product_categories, brands)
VALUES (?, ?);

-- name: InsertProductImage :execresult
INSERT INTO product_image (title, product_id, source, thumbnail, create_by, creation_date, orders)
VALUES (?, ?, ?, ?, ?, ?, ?);

-- name: SearchProducts :many
SELECT *
from product
where name like ? LIMIT ?
OFFSET ?;

-- name: GetProduct :many
SELECT *
from product
where id = ?;

-- name: GetProductInfoImg :many
SELECT *
from product_image
where product_id = ?;

-- name: ListProductsByBrand :many
SELECT *
from product
where brand_id = ?  LIMIT ?
OFFSET ?;

-- name: ListProductsByProductCategory :many
SELECT *
from product
where product_category_id = ? LIMIT ?
OFFSET ?;

-- name: CheckBrand :one
SELECT name, logo
from brand
where id = ?;

-- name: CheckCategory :many
SELECT *
from product_category
where grade = ?
  or parent_id = ?
  or id = ?
  or name = ?
LIMIT ?
OFFSET ?
;


-- name: CheckProductImage :many
SELECT source, thumbnail, title
from product_image
where product_id = ?
  or title = ?;

-- name: CheckBrandByProductCategory :many
SELECT *
from product_category_brand
where product_categories = ?;

-- name: CheckProductCategoryByBrand :many
SELECT *
from product_category_brand
where brands = ?;

-- name: GetProductByBrand :many
SELECT *
from product
where brand_id = ?;

-- name: GetProductByCategory :many
SELECT *
from product
where product_category_id = ?;

-- name: UpdateProduct :exec
UPDATE product
SET name = sqlc.arg(name),
    image = sqlc.arg(image),
    introduction = sqlc.arg(introduction),
    memo = sqlc.arg(memo),
    keyword = sqlc.arg(keyword),
    seo_title = sqlc.arg(seo_title),
    seo_description = sqlc.arg(seo_description),
    brand_id = sqlc.arg(brand_id),
    product_category_id = sqlc.arg(product_category_id),
    price = sqlc.arg(product_category_id)
where id = ?
   or sn = ?
   or name = ?;

-- name: UpdateProductCategory :exec
UPDATE product_category
SET name = sqlc.arg(name),
    image = sqlc.arg(image),
    seo_title = sqlc.arg(seo_title),
    seo_keywords = sqlc.arg(seo_keywords),
    seo_description = sqlc.arg(seo_description),
    tree_path = sqlc.arg(tree_path),
    grade = sqlc.arg(grade),
    parent_id = sqlc.arg(parent_id)
where id = ?
   or name = ?;

-- name: UpdateProductCategoryOfBrand :exec
UPDATE product_category_brand
SET product_categories = sqlc.arg(product_categories),
    brands = sqlc.arg(brands)
where product_categories = ?
   or brands = ?;

-- name: UpdateBrand :exec
UPDATE brand
SET name = sqlc.arg(name),
    logo = sqlc.arg(logo)
where id = ? or name = ?;


-- name: UpdateProductImage :exec
UPDATE product_image
SET source = sqlc.arg(source),
    thumbnail = sqlc.arg(thumbnail)
where id = ?
   or title = ?
   or product_id = ?
;

-- name: DeleteCategory :exec
DELETE FROM product_category
where id = ?
   or parent_id = ?
or name = ?
or grade = ?
;

-- name: DeleteProductImg :exec
DELETE FROM product_image
where id = ?
   or title = ?
   or product_id = ?
;

-- name: DeleteBrand :exec
DELETE FROM brand
where id = ?
   or name = ?
;