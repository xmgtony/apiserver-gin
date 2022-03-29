# sql脚本
CREATE DATABASE `apiserver_gin`;
USE `apiserver_gin`;
# 创建表也可以使用gorm提供的自动方式
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`
(
    `id`             bigint unsigned auto_increment PRIMARY KEY,
    `name`           varchar(32) not null comment '用户名称',
    `password`       char(64)    not null comment '用户密码',
    `mobile`         char(11) comment '用户手机号',
    `email`          varchar(128) comment '电子邮箱',
    `sex`            tinyint     not null default 0 comment '0未知，1男，2女',
    `age`            tinyint     not null default 0 comment '年纪，0表示未知',
    `enabled_status` tinyint     not null default 1 comment '用户账户有效状态，1正常0无效',
    `birthday`       date comment '用户出生日期，一般年月日即可',
    `created`        datetime             default CURRENT_TIMESTAMP comment '用户创建时间',
    `modified`       timestamp            default CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP comment '用户修改时间',
    index idx_account_mobile_email (`mobile`, `email`)
) engine = InnoDB
  charset = utf8 comment '用户信息表';
# 初始化数据
INSERT INTO `user` (`name`, `password`, `mobile`, `email`, `enabled_status`)
VALUES ('测试账户', '$2a$10$vTNEx1HYfIUkMAYUs0Wz/uLrmZjb.WirLJF0ONU5/roHkX0O/6VyO', '10100000000',
        'xmgtony@gmail.com', 1);