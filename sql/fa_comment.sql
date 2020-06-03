create table fa_comment
(
    id bigint(64) auto_increment comment '评论主键ID',
    video_id bigint(64) not null default '0' comment '视频ID',
    creator bigint(64) not null default '0' comment '发评论用户ID',
    comment_text text not null comment '评论文本',
    status tinyint(4) not null default '0' comment '评论状态：-1已删除，1正常',
    create_time timestamp not null default CURRENT_TIMESTAMP comment '创建时间',
    modify_time timestamp not null default CURRENT_TIMESTAMP comment '修改时间',
    primary key (`id`),
    key (`video_id`),
    key (`creator`)
) ENGINE = 'InnoDB' DEFAULT CHARACTER SET = utf8 COMMENT = '视频评论表';