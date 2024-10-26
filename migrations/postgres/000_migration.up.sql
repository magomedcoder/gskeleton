create database gskeleton;

create table users
(
    id         serial primary key,
    username   varchar(255) not null,
    password   varchar      not null,
    created_at timestamp    not null
);
