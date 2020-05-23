create table fa_user
(
	id bigint(64) auto_increment comment '用户ID',
	user_name varchar(128) not null default '' comment '用户名',
	nick_name varchar(128) not null default '' comment '昵称',
	password varchar(512) not null default '' comment '密码',
	phone varchar(64) not null default '' comment '手机号',
	sex tinyint(2) not null default 0 comment '性别',
    mail varchar(128) not null default '' comment '邮箱',
    birthday timestamp not null default CURRENT_TIMESTAMP comment '生日',
    avatar varchar(1024) not null default '' comment '头像',
    status tinyint(2) not null default 0 comment '状态',
    create_time timestamp not null default CURRENT_TIMESTAMP comment '创建时间',
    modify_time timestamp not null default CURRENT_TIMESTAMP comment '修改时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY (`user_name`, `status`),
    KEY (`phone`),
    KEY (`mail`)
) ENGINE = 'InnoDB' DEFAULT CHARACTER SET = utf8 COMMENT = '外部用户表';