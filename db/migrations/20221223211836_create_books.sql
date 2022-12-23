-- +goose Up
CREATE TABLE books (
    id INT UNSIGNED AUTO_INCREMENT,
    uuid VARCHAR(100),
    name VARCHAR(100),
    created_at TIMESTAMP,
    updated_at TIMESTAMP null,
    deleted_at TIMESTAMP null,
    PRIMARY KEY(id)
);
