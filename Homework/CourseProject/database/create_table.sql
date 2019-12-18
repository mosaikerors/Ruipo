drop database if exists shortlink;
create database shortlink;
use shortlink;
drop table if exists link;
create table link
(
    link_index int auto_increment,
    long_link varchar(100) not null,
    constraint link_pk
        primary key (link_index)
);