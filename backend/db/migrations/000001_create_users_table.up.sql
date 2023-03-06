CREATE TABLE users (
    id VARCHAR(255) PRIMARY KEY UNIQUE,
    username VARCHAR(255) UNIQUE,
    password VARCHAR(255),
    created_at TIMESTAMP default NOW(),
    updated_at TIMESTAMP default NOW()
);
