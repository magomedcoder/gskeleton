CREATE TABLE users
(
    id         SERIAL PRIMARY KEY,
    username   VARCHAR(255) NOT NULL,
    password   VARCHAR      NOT NULL,
    name       VARCHAR,
    created_at TIMESTAMP DEFAULT now()
);
