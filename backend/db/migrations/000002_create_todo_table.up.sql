CREATE TABLE todos (
    id VARCHAR(255) DEFAULT gen_random_uuid() PRIMARY KEY,
    uid VARCHAR(255),
    todo VARCHAR(266),
    is_done BOOLEAN,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    CONSTRAINT fk_user FOREIGN KEY(uid) REFERENCES users(id) ON DELETE CASCADE
);
