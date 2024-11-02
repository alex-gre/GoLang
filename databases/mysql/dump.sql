create database productdb;
use productdb;
create table products (
    id int auto_increment primary key,
    model varchar(30) not null,
    company varchar(30) not null,
    price int not null
)
