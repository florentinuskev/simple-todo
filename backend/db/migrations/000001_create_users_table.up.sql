CREATE TABLE users (
    id VARCHAR(255) DEFAULT gen_random_uuid() PRIMARY KEY,
    username VARCHAR(255) UNIQUE,
    password VARCHAR(255),
    created_at TIMESTAMP default NOW(),
    updated_at TIMESTAMP default NOW()
);
