


drop table if exists order_item;
drop table if exists orders;
drop table if exists product_reviews;
drop table if exists user;
drop table if exists product;
drop table if exists images;

create table user (
    id int primary key auto_increment,
    username varchar(50) not null unique,
    email varchar(50) not null unique,
    password varchar(50) not null,
    name varchar(255) not null,
    phone varchar(20),
    roletype varchar(20) not null,
    created_in timestamp default current_timestamp()
)ENGINE=INNODB;

create table images(
                       image_id int primary key auto_increment,
                       filename varchar(255),
                       path varchar(255)
)ENGINE=INNODB;

create table product (
                         id int primary key auto_increment,
                         name varchar(100) not null,
                         description text not null,
                         price float not null,
                         stock int not null,
                         product_image_id int not null,
                         foreign key(product_image_id)
                             references images(image_id),
                         added_in timestamp default current_timestamp()
)ENGINE=INNODB;

create table orders (
    order_id int primary key not null,
    user_id int not null,
    date timestamp default current_timestamp(),
    status varchar(50),
    foreign key (user_id)
                    references user(id)
                    on delete cascade
)ENGINE=INNODB;

create table order_item(
    product_id int primary key not null,
    foreign key (product_id)
                       references product(id)
                       on delete cascade,
    order_id int not null,
    foreign key (order_id )
                       references orders(order_id)
                       on delete cascade,
    amount int not null,
    price float(8,2) not null
)ENGINE=INNODB;

create table product_reviews (
    review_id int primary key auto_increment,
    product_id int not null,
    foreign key (product_id)
    references product(id)
    on delete cascade,
    user_id int not null,
    foreign key (user_id)
    references user(id)
    on delete cascade,
    review text,
    stars int not null,
    date timestamp default current_timestamp()
)ENGINE=INNODB;

