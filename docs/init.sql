# sql脚本
CREATE DATABASE `go_test`;
USE `go_test`;
# 创建表也可以使用gorm提供的自动方式
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`
(
    `id`             bigint unsigned auto_increment PRIMARY KEY,
    `name`           varchar(32) not null unique key comment '用户姓名',
    `password`       char(64)    not null comment '用户密码',
    `enabled_status` tinyint   default 1 comment '用户账户有效状态,1正常0无效',
    `birthday`       date comment '用户出生日期,一般年月日即可',
    `created`        datetime  default CURRENT_TIMESTAMP comment '用户创建时间',
    `modified`       timestamp default CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP comment '用户修改时间'
) engine = InnoDB
  charset = utf8 comment '用户信息表';
# 初始化数据
INSERT INTO `user` (`name`, `password`, `enabled_status`)
VALUES ('xmgtony', '$2a$10$vTNEx1HYfIUkMAYUs0Wz/uLrmZjb.WirLJF0ONU5/roHkX0O/6VyO', 1);