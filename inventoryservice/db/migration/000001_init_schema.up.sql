CREATE TABLE `stock_log`
(
    `id`                bigint(29) NOT NULL AUTO_INCREMENT COMMENT '主键_id',
    `in_quantity`       int(11) NOT NULL COMMENT '入库数量',
    `out_quantity`      int(11) NOT NULL COMMENT '出库数量',
    `stock`             int(11) NOT NULL COMMENT '当前库存',
    `product_id`        bigint(20) NOT NULL COMMENT 'SKU',
    `memo`              varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '备注',
    `create_by`         varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '创建人',
    `creation_date`     datetime NULL DEFAULT NULL COMMENT '创建日期',
    `last_updated_by`   varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '最后修改人',
    `last_updated_date` datetime NULL DEFAULT NULL COMMENT '最后修改日期',
    `delete_flag`       tinyint(1) DEFAULT 0 NOT NULL COMMENT '删除标记',
    PRIMARY KEY (`id`) USING BTREE,
    INDEX               `sku_id`(`sku_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = DYNAMIC;

CREATE TABLE `sku`
(
    `id`              bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键_id',
    `sn`              varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '编号',
    `stock`           int(11) NOT NULL COMMENT '库存',
    `allocated_stock` int(11) DEFAULT 0 NOT NULL COMMENT '已分配库存',
    `price`           decimal(22, 6)                                          NOT NULL COMMENT '售价',
    `product_id`      bigint(20) NOT NULL COMMENT '商品',
    `delete_flag`     tinyint(1) DEFAULT 0 NOT NULL COMMENT '删除标记',
    PRIMARY KEY (`id`) USING BTREE,
    INDEX             `ind_Sku_product_id`(`product_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 189 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = DYNAMIC;