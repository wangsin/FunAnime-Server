create table fa_collection
(
    id bigint(64) auto_increment comment '收藏视频主键ID',
    collection_id bigint(64) not null default 0 comment '收藏夹主ID',
    video_id bigint(64) not null default 0 comment '视频ID',
    video_cover_img varchar(1024) not null default '' comment '封面图，冗余查寻',
    video_name varchar(256) not null default '' comment '视频标题，冗余查寻',
    status tinyint(4) not null default '0' comment '收藏视频状态：-1已删除，1正常',
    create_time timestamp not null default CURRENT_TIMESTAMP comment '创建时间',
    modify_time timestamp not null default CURRENT_TIMESTAMP comment '修改时间',
    primary key (`id`),
    key (`collection_id`),
    key (`video_id`),
    key (`video_name`),
    key (`status`)
) ENGINE = 'InnoDB' DEFAULT CHARACTER SET = utf8 COMMENT = '收藏夹详情表';