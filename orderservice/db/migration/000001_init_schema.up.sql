CREATE TABLE `orders`
(
    `id`                   bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键_id',
    `sn`                   varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '订单编号',
    `status`               int(11) NOT NULL DEFAULT 0 COMMENT '订单状态',
    `amount`               decimal(21, 6)                                          NOT NULL COMMENT '订单金额',
    `coupon_code_id`       varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '优惠卷id',
    `coupon_discount`      decimal(21, 6)                                          NOT NULL COMMENT '优惠券折扣',
    `user_id`              bigint(20) NOT NULL COMMENT '会员',
    `consignee`            varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '收货人',
    `address`              varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '地址',
    `phone`                varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '电话',
    `expire`               datetime NULL DEFAULT NULL COMMENT '到期时间',
    `type`                 int(11) NOT NULL COMMENT '类型',
    `refund_amount`        decimal(21, 9)                                          NOT NULL DEFAULT 0 COMMENT '退款金额',
    `shipping_method_name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '配送方式名称',
    `shipping_date`        datetime NULL DEFAULT NULL COMMENT '配送时间',
    `complete_date`        datetime NULL DEFAULT NULL COMMENT '完成日期',
    `creation_date`        datetime NULL DEFAULT NULL COMMENT '创建日期',
    `delete_flag`          tinyint(1) NOT NULL DEFAULT 0 COMMENT '删除标记',
    PRIMARY KEY (`id`) USING BTREE,
    UNIQUE INDEX `sn`(`sn`) USING BTREE,
    INDEX                  `FK25E6B94FC050045D`(`coupon_code_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 72 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '订单' ROW_FORMAT = DYNAMIC;


CREATE TABLE `order_item`
(
    `id`                bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键_id',
    `sn`                varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '商品编号',
    `name`              varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '商品名称',
    `price`             int(11) NOT NULL COMMENT '商品价格',
    `weight`            int(11) NULL DEFAULT NULL COMMENT '商品重量',
    `thumbnail`         varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '商品缩略图',
    `quantity`          int(11) NOT NULL COMMENT '数量',
    `shipped_quantity`  int(11) NOT NULL COMMENT '已发货数量',
    `return_quantity`   int(11) NOT NULL COMMENT '已退货数量',
    `order_id`          bigint(20) NOT NULL COMMENT '订单',
    `sku_id`            bigint(20) NULL DEFAULT NULL COMMENT 'SKU',
    `create_by`         varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '创建人',
    `creation_date`     datetime NULL DEFAULT NULL COMMENT '创建日期',
    `last_updated_by`   varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '最后修改人',
    `last_updated_date` datetime NULL DEFAULT NULL COMMENT '最后修改日期',
    `delete_flag`       tinyint(1) NOT NULL DEFAULT 0 COMMENT '删除标记',
    PRIMARY KEY (`id`) USING BTREE,
    INDEX               `FKD69FF403B992E8EF`(`order_id`) USING BTREE,
    UNIQUE INDEX `sn`(`sn`) USING BTREE,
    INDEX               `sku_id`(`sku_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 72 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '订单项' ROW_FORMAT = DYNAMIC;