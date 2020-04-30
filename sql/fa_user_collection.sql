create table fa_user_collection
(
    id bigint(64) auto_increment comment '收藏夹主ID',
    user_id bigint(64) not null default 0 comment '用户ID',
    collection_name varchar(64) not null default '收藏夹名称',
    cover_img varchar(1024) not null default '' comment '封面图，冗余查寻',
    is_share tinyint(2) not null default '0' comment '是否对外开放',
    status tinyint(4) not null default '0' comment '视频状态：-1已删除，1正常',
    primary key (`id`),
    key (`user_id`),
    key (`status`)
) ENGINE = 'InnoDB' DEFAULT CHARACTER SET = utf8 COMMENT = '用户-收藏夹对应表';