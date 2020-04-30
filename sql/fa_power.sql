create table fa_config
(
    id bigint(64) auto_increment comment '权限ID',
    pow_name varchar(128) not null default '' comment '权限名称',
    pow_rule varchar(64) not null default '' comment '权限点',
    inner_creator bigint(64) not null default '0' comment '添加用户',
    status tinyint(2) not null default 0 comment '状态：-1已删除 1生效',
    create_time timestamp not null default CURRENT_TIMESTAMP comment '创建时间',
    modify_time timestamp not null default CURRENT_TIMESTAMP comment '修改时间',
    primary key (`id`),
    key (`pow_rule`),
    key (`pow_name`),
    key (`inner_creator`)
) ENGINE = 'InnoDB' DEFAULT CHARACTER SET = utf8 COMMENT = '内部用户权限表';