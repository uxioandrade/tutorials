CREATE DATABASE todos_db;

\connect todos_db;

CREATE TABLE todos(
    id SERIAL NOT NULL PRIMARY KEY,
    description VARCHAR(100),
    priority INT,
    status VARCHAR(100)
);