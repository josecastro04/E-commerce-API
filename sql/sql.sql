drop table produtos;
drop table carrinho;
drop table user;

create table user (
    id int primary key auto_increment,
    name varchar(50) not null,
    email varchar(50) not null unique,
    password varchar(50) not null,
    created_in timestamp default current_timestamp()
)ENGINE=INNODB;