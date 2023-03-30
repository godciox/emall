CREATE TABLE `cart`
(
    `id`            bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键_id',
    `user_id`       bigint(20) NULL DEFAULT NULL COMMENT '会员',
    `cart_key`      varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '密钥',
    `create_by`     varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '创建人',
    `creation_date` datetime NULL DEFAULT NULL COMMENT '创建日期',
    `delete_flag`   tinyint(1) NOT NULL DEFAULT 0 COMMENT '删除标记',
    PRIMARY KEY (`id`) USING BTREE,
    UNIQUE(`user_id`)
) ENGINE = InnoDB AUTO_INCREMENT = 985 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '购物车' ROW_FORMAT = DYNAMIC;

CREATE TABLE `cart_item`
(
    `id`            bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键_id',
    `cart_id`       bigint(20) NOT NULL COMMENT '购物车',
    `product_id`    bigint(20) NOT NULL COMMENT 'SKU',
    `quantity`      int(11) NOT NULL COMMENT '数量',
    `create_by`     varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '创建人',
    `creation_date` datetime NULL DEFAULT NULL COMMENT '创建日期',
    `delete_flag`   tinyint(1) NOT NULL DEFAULT 0 COMMENT '删除标记',
    PRIMARY KEY (`id`) USING BTREE,
    INDEX           `FK1A67F65339A23004`(`cart_id`) USING BTREE,
    INDEX           `FK1A67F65379F8D99A`(`product_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1786 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '购物车项' ROW_FORMAT = DYNAMIC;