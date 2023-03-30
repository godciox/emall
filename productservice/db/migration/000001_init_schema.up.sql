CREATE TABLE `brand`
(
    `id`            bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键_id',
    `name`          varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '名称',
    `logo`          varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT 'logo',
    `orders`        int(11) NULL DEFAULT NULL COMMENT '排序',
    `create_by`     varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '创建人',
    `creation_date` datetime NULL DEFAULT NULL COMMENT '创建日期',
    `delete_flag`   tinyint(1) NOT NULL DEFAULT 0 COMMENT '删除标记',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 26 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '品牌' ROW_FORMAT = DYNAMIC;


CREATE TABLE `product`
(
    `id`                  bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键_id',
    `sn`                  varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '编号',
    `name`                varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '名称',
    `image`               varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '展示图片',
    `introduction`        longtext CHARACTER SET utf8 COLLATE utf8_general_ci NULL COMMENT '介绍',
    `memo`                varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '备注',
    `keyword`             varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '搜索关键词',
    `seo_title`           varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '页面标题',
    `seo_keywords`        varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '页面关键词',
    `seo_description`     varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '页面描述',
    `score`               float         DEFAULT 0                                         NOT NULL COMMENT '评分',
    `price`               float NOT NULL COMMENT '价格',
    `total_score`         bigint(20) NOT NULL DEFAULT 0 COMMENT '总评分',
    `score_count`         bigint(20) NOT NULL DEFAULT 0 COMMENT '评分数',
    `hits`                bigint(20) NOT NULL DEFAULT 0 COMMENT '点击数',
    `week_hits`           bigint(20) NOT NULL DEFAULT 0 COMMENT '周点击数',
    `month_hits`          bigint(20) NOT NULL DEFAULT 0 COMMENT '月点击数',
    `sales`               bigint(20) NOT NULL DEFAULT 0 COMMENT '销量',
    `week_sales`          bigint(20) NOT NULL DEFAULT 0 COMMENT '周销量',
    `month_sales`         bigint(20) NOT NULL DEFAULT 0 COMMENT '月销量',
    `brand_id`            bigint(20) NULL DEFAULT NULL COMMENT '品牌',
    `product_category_id` bigint(20) NOT NULL COMMENT '商品分类',
    `create_by`           varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '创建人',
    `creation_date`       datetime NULL DEFAULT NULL COMMENT '创建日期',
    `delete_flag`         tinyint(1) NOT NULL DEFAULT 0 COMMENT '删除标记',
    PRIMARY KEY (`id`) USING BTREE,
    UNIQUE INDEX `sn`(`sn`) USING BTREE,
    INDEX                 `FK7C9E82B0D7629117`(`product_category_id`) USING BTREE,
    INDEX                 `FK7C9E82B0FA9695CA`(`brand_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 433 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '商品' ROW_FORMAT = DYNAMIC;


CREATE TABLE `product_category`
(
    `id`              bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键_id',
    `name`            varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '名称',
    `seo_title`       varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '页面标题',
    `seo_keywords`    varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '页面关键词',
    `seo_description` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '页面描述',
    `tree_path`       varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '树路径',
    `grade`           int(11) NOT NULL COMMENT '层级',
    `image`           varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '图片',
    `parent_id`       bigint(20) NULL DEFAULT NULL COMMENT '上级分类',
    `create_by`       varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '创建人',
    `creation_date`   datetime NULL DEFAULT NULL COMMENT '创建日期',
    `delete_flag`     tinyint(1) NOT NULL DEFAULT 0 COMMENT '删除标记',
    PRIMARY KEY (`id`) USING BTREE,
    INDEX             `FK1B7971ADFBDD5B73`(`parent_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 243 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '商品分类' ROW_FORMAT = DYNAMIC;


CREATE TABLE `product_category_brand`
(
    `product_categories` bigint(20) NOT NULL,
    `brands`             bigint(20) NOT NULL,
    PRIMARY KEY (`product_categories`, `brands`) USING BTREE,
    INDEX                `FKE42D6A75A2AB700F`(`brands`) USING BTREE,
    INDEX                `FKE42D6A758C4C0635`(`product_categories`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = DYNAMIC;

CREATE TABLE `product_image`
(
    `id`            bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键_id',
    `title`         varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '标题',
    `product_id`    bigint(20) NOT NULL COMMENT '商品',
    `source`        varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '原图片',
    `thumbnail`     varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '缩略图',
    `orders`        int(11) NULL DEFAULT NULL COMMENT '排序',
    `create_by`     varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '创建人',
    `creation_date` datetime NULL DEFAULT NULL COMMENT '创建日期',
    `delete_flag`   tinyint(1) NOT NULL DEFAULT 0 COMMENT '删除标记',
    PRIMARY KEY (`id`) USING BTREE,
    INDEX           `FK66470ABC79F8D99A`(`product_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 287 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '商品图片' ROW_FORMAT = DYNAMIC;