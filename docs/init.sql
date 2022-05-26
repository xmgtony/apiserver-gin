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

# 记账程序账目清单，eg.好友结婚随份子可以记一笔、同事结婚随礼记一笔、买手办记一笔等等。
# TIPS: 只是演示脚手架的使用，不考虑业务合理性。
CREATE TABLE IF NOT EXISTS `account_bill`
(
    `id`              bigint unsigned auto_increment PRIMARY KEY,
    `user_id`         bigint unsigned  not null default 0 comment '所属用户id',
    `bill_date`       date             not null comment '账单日期',
    `origin_incident` varchar(512)     not null default '' comment '账户产生的事由',
    `amount`          int unsigned     not null default 0 comment '账单金额（单位分）',
    `relation`        varchar(32)      not null default '' comment '与对方关系,如亲戚|同事|闺蜜',
    `to_name`         varchar(32)      not null default '' comment '对方姓名',
    `is_follow`       tinyint unsigned not null default 0 comment '是否关注或者跟进，0不关注、1关注',
    `remark`          tinytext comment '备注说明',
    `enabled_status`  tinyint          not null default 1 comment '有效状态，1正常、0无效',
    `created`         datetime         not null default CURRENT_TIMESTAMP comment '用户创建时间',
    `modified`        timestamp on update CURRENT_TIMESTAMP comment '用户修改时间',
    index idx_user_id (`user_id`)
) engine = InnoDB
  charset = utf8 comment '账目清单';