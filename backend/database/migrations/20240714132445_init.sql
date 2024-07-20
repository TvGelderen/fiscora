-- +goose Up
CREATE TABLE IF NOT EXISTS users (
    id VARCHAR(128) UNIQUE NOT NULL PRIMARY KEY,
    provider VARCHAR(16) NOT NULL,
    provider_id VARCHAR(64) NOT NULL,
    username VARCHAR(32) UNIQUE NOT NULL,
    password_hash VARCHAR(128) NOT NULL,
    email VARCHAR(64) UNIQUE,
    created TIMESTAMP,
    UNIQUE(provider, provider_id)
);

CREATE TABLE IF NOT EXISTS sessions (
    id VARCHAR(128) UNIQUE NOT NULL PRIMARY KEY,
    expires_at INT NOT NULL,
    user_id VARCHAR(128) NOT NULL,
    CONSTRAINT fk_user_id 
        FOREIGN KEY(user_id) 
        REFERENCES users(id)
        ON DELETE CASCADE
);

-- +goose Down
DROP TABLE sessions;
DROP TABLE users;
