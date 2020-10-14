drop database if exists `blog`;
create database `blog` default character set utf8mb4 collate utf8mb4_general_ci;
use `blog`;
drop table if exists `blog_article`;
CREATE TABLE `blog_article`
(
    `id`          int          NOT NULL AUTO_INCREMENT,
    `title`       varchar(64)  NOT NULL COMMENT '文章标题',
    `desc`        varchar(256) NOT NULL COMMENT '文章概述',
    `label_id` int          NOT NULL COMMENT '文章分类',
    `context`     text         NOT NULL COMMENT '文章内容',
    `state`       tinyint      NOT NULL DEFAULT '0' COMMENT '状态0：正常1：禁用2：删除',
    `createTime`  int          NOT NULL DEFAULT '0' COMMENT '创建时间',
    `updateTime`  int          NOT NULL DEFAULT '0' COMMENT '修改时间',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_0900_ai_ci COMMENT ='文章表';
drop table if exists `blog_label`;
CREATE TABLE `blog_label`
(
    `id`         int         NOT NULL AUTO_INCREMENT,
    `name`       varchar(64) NOT NULL COMMENT '类目名称',
    `type`       tinyint     NOT NULL COMMENT '1:分类2：标签',
    `state`      tinyint     NOT NULL DEFAULT '0' COMMENT '状态0：正常1：禁用2：删除',
    `createTime` int         NOT NULL DEFAULT '0' COMMENT '创建时间',
    `updateTime` int         NOT NULL DEFAULT '0' COMMENT '修改时间',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_0900_ai_ci COMMENT ='文章类目表';

drop table if exists `blog_article_label`;
CREATE TABLE `blog_article_label`
(
    `id`         int NOT NULL AUTO_INCREMENT,
    `article_id` int NOT NULL COMMENT '文章id',
    `label_id`   int NOT NULL COMMENT '标签id',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_0900_ai_ci COMMENT ='文章标签关联表';

drop table if exists `blog_user`;
create table `blog_user`
(
    `id`         int(11)     not null auto_increment,
    `username`   varchar(64) not null comment '用户名',
    `password`   varchar(64) not null comment '密码',
    `tel`        varchar(11) not null comment '手机号',
    `email`      varchar(64) not null comment '邮箱',
    `state`      tinyint     NOT NULL DEFAULT '0' COMMENT '状态0：正常1：禁用2：删除',
    `createTime` int         NOT NULL DEFAULT '0' COMMENT '创建时间',
    `updateTime` int         NOT NULL DEFAULT '0' COMMENT '修改时间',
    primary key (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_0900_ai_ci COMMENT ='文章标签关联表'