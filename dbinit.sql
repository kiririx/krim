create table user
(
    id         int auto_increment
        primary key,
    username   varchar(200)  not null,
    password   varchar(1000) not null,
    nickname   varchar(200)  not null,
    sex        int default 0 not null,
    created_at datetime null,
    updated_at datetime null
);


create table event
(
    id         int auto_increment
        primary key,
    created_at datetime null,
    updated_at datetime null,
    event_type int default 0 not null,
    source_id  int default 0 not null,
    target_id  int default 0 not null,
    progress   int default 0 not null,
)