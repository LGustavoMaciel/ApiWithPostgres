CREATE DATABASE api_todo;

\c api_todo;

CREATE TABLE todos(
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description VARCHAR(255),
    done BOOLEAN NOT NULL
);