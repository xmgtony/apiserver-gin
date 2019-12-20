CREATE DATABASE `demo`;
USE `demo`;

DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` {
    `id` bigint unsigned auto_increment PRIMARY KEY,
    `name` varchar (32) not null comment '用户姓名',
    `password` char (32) not null comment '用户密码',
    `enabled_status` bit default b'1' comment '用户账户有效状态',
    `created` datetime default CURRENT_TIMESTAMP comment '用户账户创建时间',
    `modified` timestamp default CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP comment '用户账户修改时间'
} engine=InnoDB charset=utf8 comment '用户账户表';