CREATE TABLE `admin`
(
    `admin_id`                      int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '序号',
    `admin_mobile`                  varchar(32) NOT NULL DEFAULT '' COMMENT '手机号',
    `admin_password`                varchar(32) NOT NULL DEFAULT '' COMMENT '密码',
    `admin_salt`                    varchar(32) NOT NULL DEFAULT '' COMMENT '加密盐',
    `admin_name`                    varchar(32) NOT NULL DEFAULT '' COMMENT '管理员名称',
    `admin_login_token`             varchar(32) NOT NULL DEFAULT '' COMMENT '用户登录token',
    `admin_login_token_expire_time` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '用户登录token过期时间',
    `admin_last_login_time`         datetime    NOT NULL DEFAULT '2000-01-01 00:00:00' COMMENT '管理员用户最后登录时间',
    `admin_last_login_ip`           varchar(15) NOT NULL DEFAULT '' COMMENT '管理员用户最后登录ip',
    `admin_role_id`                 int(10) unsigned NOT NULL DEFAULT '0' COMMENT '用户角色',
    `admin_status`                  tinyint(1) NOT NULL DEFAULT '0' COMMENT '管理员状态  1 启用  -1禁用',
    `operation_admin_id`            int(10) unsigned NOT NULL DEFAULT '0' COMMENT '操作人序号',
    `operation_admin_name`          varchar(20) NOT NULL DEFAULT '' COMMENT '管理员序号',
    `operation_admin_ip`            varchar(15) NOT NULL DEFAULT '' COMMENT '操作IP',
    `admin_create_time`             datetime    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `admin_update_time`             datetime    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
    `admin_email`                   varchar(60) NOT NULL DEFAULT '' COMMENT '管理员邮箱',
    PRIMARY KEY (`admin_id`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8 COMMENT='后台用户表'


CREATE TABLE `admin_role`
(
    `admin_role_id`          int(10) NOT NULL AUTO_INCREMENT COMMENT '角色ID',
    `admin_role_name`        varchar(30) NOT NULL COMMENT '角色名称',
    `admin_permission_id`    text        NOT NULL COMMENT '权限ID',
    `store_id`               text        NOT NULL COMMENT '门店ID',
    `admin_role_status`      tinyint(3) NOT NULL DEFAULT '0' COMMENT '角色状态 -1禁用 1正常',
    `operation_admin_id`     int(10) NOT NULL DEFAULT '0' COMMENT '操作人id',
    `operation_admin_name`   varchar(50) NOT NULL COMMENT '操作人姓名',
    `operation_admin_ip`     varchar(15) NOT NULL COMMENT '操作人ip',
    `admin_role_create_time` datetime    NOT NULL COMMENT '创建时间',
    PRIMARY KEY (`admin_role_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='后台角色表'

CREATE TABLE `admin_permission`
(
    `admin_permission_id`          int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '管理员权限表',
    `admin_permission_name`        varchar(20) NOT NULL DEFAULT '' COMMENT '管理员权限名称',
    `admin_permission_path`        varchar(50) NOT NULL DEFAULT '' COMMENT '管理员权限地址',
    `admin_permission_parent_id`   int(10) unsigned NOT NULL DEFAULT '0' COMMENT '父级权限序号',
    `admin_permission_status`      tinyint(3) NOT NULL DEFAULT '0' COMMENT '权限状态 -1禁用 1正常',
    `admin_permission_menu`        tinyint(1) unsigned NOT NULL DEFAULT '1' COMMENT '是否在菜单栏显示0不显示，1显示',
    `admin_permission_sort`        tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '权限排序',
    `operation_admin_id`           int(10) unsigned NOT NULL DEFAULT '0' COMMENT '操作人序号',
    `operation_admin_name`         varchar(20) NOT NULL DEFAULT '' COMMENT '管理员序号',
    `operation_admin_ip`           varchar(15) NOT NULL DEFAULT '' COMMENT '操作IP',
    `admin_permission_create_time` timestamp   NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '创建时间',
    `admin_permission_update_time` datetime    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`admin_permission_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='后台权限表'