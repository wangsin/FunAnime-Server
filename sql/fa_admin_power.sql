create table fa_admin_power
(
    id bigint(64) auto_increment comment '规则ID',
    user_name varchar(128) not null default '' comment '内部用户名 冗余查询',
    user_id bigint(64) not null default '0' comment '内部用户ID',
    status tinyint(2) not null default 0 comment '状态：-1已删除 1生效',
    power bigint(64) not null default '0' comment '权限',
    creator varchar(128) not null default '' comment '创建人',
    create_time timestamp not null default CURRENT_TIMESTAMP comment '创建时间',
    modify_time timestamp not null default CURRENT_TIMESTAMP comment '修改时间',
    PRIMARY KEY (`id`),
    key (`user_id`),
    key (`creator`)
) ENGINE = 'InnoDB' DEFAULT CHARACTER SET = utf8 COMMENT = '用户权限规则配置表';