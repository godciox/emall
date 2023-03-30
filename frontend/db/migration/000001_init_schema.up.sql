CREATE TABLE users
(
    `id`                  bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键_id',
    `username`            varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '用户名',
    `password`            varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '密码',
    `name`                varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '姓名',
    `avatar`              varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '头像',
    `gender`              int(11) NULL DEFAULT NULL COMMENT '性别,0未知,1男,2女性别',
    `mobile`              varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '手机',
    `birth`               datetime NULL DEFAULT NULL COMMENT '出生日期',
    `zip_code`            varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '邮编',
    `address`             varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '地址',
    `email`               varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT 'E-mail',
    `login_date`          datetime NULL DEFAULT NULL COMMENT '最后登录日期',
    `login_failure_count` int(11) NOT NULL COMMENT '连续登录失败次数',
    `creation_date`       datetime NULL DEFAULT NULL COMMENT '创建日期',
    `last_updated_date`   datetime NULL DEFAULT NULL COMMENT '最后修改日期',
    `delete_flag`         tinyint(1) NOT NULL DEFAULT 0 COMMENT '删除标记',
    PRIMARY KEY (`id`) USING BTREE
);