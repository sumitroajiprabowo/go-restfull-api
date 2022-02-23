create database if not exists golang_restfull_api_test;
use golang_restfull_api_test;

create table category(
    id integer primary key auto_increment,
    name varchar(200) not null
    ) engine=InnoDB;
