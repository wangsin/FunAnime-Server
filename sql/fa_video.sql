create table fa_video
(
    id bigint(64) auto_increment comment '视频ID',
    video_name varchar(256) not null default '' comment '视频标题',
    video_remote_id varchar(128) not null default '' comment '腾讯云视频ID',
    video_desc text not null comment '视频描述',
    category_top_level bigint(64) not null default '0' comment '一级分类',
    category_top_level_desc varchar(64) not null default '' comment '一级分类描述',
    category_next_level bigint(64) not null default '0' comment '二级分类',
    category_next_level_desc varchar(64) not null default '' comment '二级分类描述',
    cover_img varchar(1024) not null default '' comment '封面图连接，可为空，视频上传后需更新',
    creator bigint(64) not null default '0' comment 'UP主用户ID',
    pv bigint(64) not null default '0' comment '单条视频PV',
    uv bigint(64) not null default '0' comment '单条视频UV',
    status tinyint(4) not null default '0' comment '视频状态：-1已删除，1审核中，2审核失败，3隐藏，4正常',
    pass_time timestamp not null default CURRENT_TIMESTAMP comment '过审时间',
    create_time timestamp not null default CURRENT_TIMESTAMP comment '创建时间',
    modify_time timestamp not null default CURRENT_TIMESTAMP comment '修改时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY (`video_remote_id`, `status`),
    KEY (`creator`),
    KEY (`video_name`)
) ENGINE = 'InnoDB' DEFAULT CHARACTER SET = utf8 COMMENT = '视频表';