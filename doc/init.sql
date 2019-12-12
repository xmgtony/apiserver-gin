CREATE DATABASE `demo`;
USE `demo`;

DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` {
    `id` bigint unsigned auto_increment PRIMARY KEY,
    `name` varchar (32) not null comment '用户姓名',
    `password` char (32) not null comment '用户密码'
} engine=InnoDB charset=utf8;