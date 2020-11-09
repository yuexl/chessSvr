create table card_percents
(
    id            int auto_increment
        primary key,
    tech_level    int null,
    percent_one   int null,
    percent_two   int null,
    percent_three int null,
    percent_four  int null,
    percent_five  int null,
    call_number   int null
)
    comment '抽卡概率表';


create table chess_cards
(
    id           int auto_increment
        primary key,
    card_name    varchar(50) default '' not null,
    card_level   int         default 1  not null,
    cost         int                    null,
    basic_attack int                    null,
    basic_health int                    null,
    sale_price   int                    null,
    card_desc    varchar(200)           null
);

create table chess_rank
(
    user_id int           not null,
    point   int default 0 not null,
    `rank`  int default 0 not null,
    constraint chess_rank_user_id_uindex
        unique (user_id)
);

alter table chess_rank
    add primary key (user_id);


create table chess_users
(
    id        int auto_increment
        primary key,
    user_name varchar(50) default '' not null,
    is_vip    smallint    default 0  not null,
    password  varchar(32)            not null
)
    comment '用户表';


