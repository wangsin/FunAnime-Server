create table fa_barrage
(
    id bigint(64) auto_increment comment '弹幕主键ID',
    video_id bigint(64) not null default '0' comment '视频ID',
    creator bigint(64) not null default '0' comment '发弹幕用户ID',
    barrage_text varchar(256) not null default '' comment '弹幕文本',
    barrage_color varchar(32) not null default '#FFF' comment '弹幕颜色',
    status tinyint(4) not null default '0' comment '弹幕状态：-1已删除，1正常',
    create_time timestamp not null default CURRENT_TIMESTAMP comment '创建时间',
    modify_time timestamp not null default CURRENT_TIMESTAMP comment '修改时间',
    primary key (`id`),
    key (`video_id`),
    key (`status`),
    key (`creator`)
) ENGINE = 'InnoDB' DEFAULT CHARACTER SET = utf8 COMMENT = '视频弹幕表';