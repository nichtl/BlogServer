create table if not exists article
(
    id           bigint auto_increment
    primary key,
    title        varchar(256) default ''                not null comment '标题',
    content      longtext                               null comment 'text内容',
    intro        varchar(512) default ''                not null comment '介绍',
    file_url     varchar(512) default ''                not null comment 'url',
    user_id      int                                    null,
    user_account varchar(128) default ''                not null,
    del_flag     int          default 0                 not null,
    add_time     datetime     default CURRENT_TIMESTAMP not null comment '新增时间',
    update_time  datetime     default CURRENT_TIMESTAMP not null comment '更新时间'
    )
    engine = InnoDB;

create table if not exists category
(
    id          bigint auto_increment
    primary key,
    father_id   bigint       default 0                 not null,
    name        varchar(128) default ''                not null comment '类别名称',
    del_flag    int          default 0                 not null,
    update_time datetime     default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP,
    add_time    datetime     default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP,
    user_id     int          default 0                 null
    )
    comment '分类表' engine = InnoDB
    collate = utf8mb4_general_ci;

create table if not exists tag
(
    id          bigint auto_increment comment '标签表'
    primary key,
    name        varchar(128) default ''                not null,
    father_id   bigint       default 0                 not null,
    path        varchar(256) default ''                not null,
    del_flag    datetime     default CURRENT_TIMESTAMP not null,
    update_time datetime     default CURRENT_TIMESTAMP not null,
    add_time    datetime     default CURRENT_TIMESTAMP not null
    )
    comment '标签表' engine = InnoDB;

create table if not exists tag_mapper
(
    id          bigint auto_increment
    primary key,
    mapper_type int      default -1                not null,
    article_id  bigint   default -1                not null comment '文章表 id',
    tag_id      bigint   default -1                not null comment '标签表 id',
    del_flag    int      default 0                 not null comment '0 有效 1无效',
    add_time    datetime default CURRENT_TIMESTAMP not null,
    update_time datetime default CURRENT_TIMESTAMP not null
)
    comment '标签映射表' engine = InnoDB;

create index tag_mapping_article_id_index
    on tag_mapper (article_id);

create index tag_mapping_tag_id_index
    on tag_mapper (tag_id);

create table if not exists user
(
    id                  bigint auto_increment
    primary key,
    real_name           varchar(128) default ''                not null comment '真实姓名',
    user_type           int          default 0                 null comment '用户类型',
    nick_name           varchar(256) default ''                not null comment '昵称',
    account             varchar(128)                           not null comment '账户名',
    password            varchar(256)                           not null comment '密码',
    email               varchar(128) default ''                not null comment '邮箱',
    name_indentity_code varchar(128)                           not null comment '用户身份码  userId 有可能变动 这个做备用',
    extend_info         varchar(1024)                          null comment '扩展信息',
    link_account        varchar(256)                           null comment '关联账户',
    gender              tinyint                                null comment '性别:0女,1男,其余为未知 ',
    signature           varchar(256)                           null comment '个性签名',
    we_chat             varchar(128) default ''                not null comment '微信',
    register_time       datetime     default CURRENT_TIMESTAMP not null,
    qq                  bigint       default 0                 not null,
    birthday            date         default '1997-01-01'      not null comment '生日年月日',
    add_time            datetime     default CURRENT_TIMESTAMP not null,
    update_time         datetime     default CURRENT_TIMESTAMP null on update CURRENT_TIMESTAMP,
    del_flag            int          default 0                 not null,
    constraint user_pk_2
    unique (account),
    constraint user_pk_code
    unique (name_indentity_code)
    )
    engine = InnoDB;

