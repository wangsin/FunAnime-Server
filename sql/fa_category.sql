create table fa_category
(
    id bigint(64) auto_increment comment '分类主键ID',
    category_name varchar(32) not null default '' comment '分类名称',
    parent_category_id bigint(64) not null default '0' comment '上级分类ID/一级分类时为0',
    creator bigint(64) not null default '0' comment '分类配置用户ID',
    create_time timestamp not null default CURRENT_TIMESTAMP comment '创建时间',
    modify_time timestamp not null default CURRENT_TIMESTAMP comment '修改时间',
    primary key (`id`),
    key (`creator`),
    key (`parent_category_id`)
) ENGINE = 'InnoDB' DEFAULT CHARACTER SET = utf8 COMMENT = '视频分类表';