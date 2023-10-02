CREATE DATABASE IF NOT EXISTS test_devs_crud;

CREATE TABLE IF NOT EXISTS devs 
(
    email varchar(255)  NOT NULL PRIMARY KEY,
    languages JSON,
	expertise int NOT NULL
);