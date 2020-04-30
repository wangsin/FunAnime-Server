create table fa_config
(
    id bigint(64) auto_increment comment '配置ID',
    config_type tinyint(2) not null default '0' comment '配置类型 0:主配置',
    status tinyint(2) not null default '1' comment '配置状态 1:生效 -1:删除',
    config_data text not null comment '配置详情格式JSON',
    config_user bigint(64) not null default '0' comment '配置人信息',
    create_time timestamp not null default CURRENT_TIMESTAMP comment '创建时间',
    modify_time timestamp not null default CURRENT_TIMESTAMP comment '修改时间',
    primary key (`id`),
    key (`config_type`),
    key (`status`),
    key (`config_user`)
) ENGINE = 'InnoDB' DEFAULT CHARACTER SET = utf8 COMMENT = '配置详情表';