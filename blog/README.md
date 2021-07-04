## golang项目练习

> 目的是快速熟悉golang常用的开发框架，完成一个简单的数据系统，主要技术栈为 `gin + gorm + viper`

## 初始化

新建一个golang项目，并且安装需要的依赖

```shell
cd $yourpath
mkdir blog
go mod init blog 
go get -u github.com/gin-gonic/gin
go get github.com/spf13/viper
go get -u gorm.io/gorm
go get -u gorm.io/driver/mysql 
```

## 数据

```sql
CREATE TABLE `blog_article`
(
    `id`          int(10) unsigned NOT NULL AUTO_INCREMENT,
    `tag_id`      int(10) unsigned    DEFAULT '0' COMMENT '标签ID',
    `title`       varchar(100)        DEFAULT '' COMMENT '文章标题',
    `desc`        varchar(255)        DEFAULT '' COMMENT '简述',
    `content`     text,
    `create_time` datetime            DEFAULT NULL,
    `created_by`  varchar(100)        DEFAULT '' COMMENT '创建人',
    `modify_time` datetime            DEFAULT '0' COMMENT '修改时间',
    `modified_by` varchar(255)        DEFAULT '' COMMENT '修改人',
    `state`       tinyint(3) unsigned DEFAULT '1' COMMENT '状态 0为禁用1为启用',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8 COMMENT ='文章管理';

CREATE TABLE `blog_tag`
(
    `id`          int(10) unsigned NOT NULL AUTO_INCREMENT,
    `name`        varchar(100)        DEFAULT '' COMMENT '标签名称',
    `create_time` datetime            DEFAULT '0' COMMENT '创建时间',
    `created_by`  varchar(100)        DEFAULT '' COMMENT '创建人',
    `modify_time` datetime            DEFAULT '0' COMMENT '修改时间',
    `modified_by` varchar(100)        DEFAULT '' COMMENT '修改人',
    `state`       tinyint(3) unsigned DEFAULT '1' COMMENT '状态 0为禁用、1为启用',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8 COMMENT ='文章标签管理';
```






