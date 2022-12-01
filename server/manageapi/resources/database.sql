--  后台用户表
CREATE TABLE `gam_user`
(
    `id`              INT(11) UNSIGNED NOT NULL AUTO_INCREMENT,
    `cellphone`       CHAR(11)     NOT NULL DEFAULT '' COMMENT '用户手机号',
    `password`        CHAR(32)     NOT NULL DEFAULT '' COMMENT '登录密码',
    `name`            VARCHAR(255) NOT NULL DEFAULT '' COMMENT '用户真实姓名',
    `icon`            VARCHAR(255) NOT NULL DEFAULT '' COMMENT '用户头像',
    `gender`          TINYINT(4) UNSIGNED NOT NULL DEFAULT '0' COMMENT '性别 0 未知 1 男 2 女',
    `birthday`        date         NOT NULL DEFAULT '0000-00-00' COMMENT '生日',
    `remark`          VARCHAR(255) NOT NULL DEFAULT '' COMMENT '备注',
    `created_user_id` INT(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建人',
    `updated_user_id` INT(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '更新人',
    `created`         DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated`         DATETIME     NOT NULL DEFAULT '0000-00-00 00:00:00' ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `last_login_time` DATETIME     NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '最后登录时间',
    `last_login_ip`   VARCHAR(15)  NOT NULL DEFAULT '' COMMENT '最后登录IP',
    `status`          TINYINT(1) UNSIGNED NOT NULL DEFAULT 1 COMMENT '状态：0禁用 1正常',
    PRIMARY KEY (`id`),
    KEY               `cellphone`(`cellphone`),
    KEY               `name`(`name`)
) ENGINE = INNODB DEFAULT CHARSET = utf8 COMMENT = '后台用户表';