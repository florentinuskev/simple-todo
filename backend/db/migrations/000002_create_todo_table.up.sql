CREATE TABLE todos (
    id VARCHAR(255) DEFAULT gen_random_uuid() PRIMARY KEY,
    uid VARCHAR(255),
    todo VARCHAR(266),
    CONSTRAINT fk_user FOREIGN KEY(uid) REFERENCES users(id) ON DELETE CASCADE
);
