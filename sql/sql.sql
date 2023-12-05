drop table if exists user;
drop table if exists product;

create table user (
    id int primary key auto_increment,
    name varchar(50) not null,
    email varchar(50) not null unique,
    password varchar(100) not null,
    roletype varchar(20) not null,
    created_in timestamp default current_timestamp()
)ENGINE=INNODB;

create table product (
    id int primary key auto_increment,
    name varchar(100) not null,
    description varchar(500) not null,
    price float(8,2) not null,
    stock int not null,
    added_in timestamp default current_timestamp()
)ENGINE=INNODB;