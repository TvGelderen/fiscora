-- +goose Up
CREATE TABLE IF NOT EXISTS users (
    id UUID UNIQUE NOT NULL PRIMARY KEY,
    provider VARCHAR(16) NOT NULL,
    provider_id VARCHAR(32) NOT NULL,
    username VARCHAR(32) UNIQUE NOT NULL,
    email VARCHAR(64) UNIQUE NOT NULL,
    password_hash BYTEA,
    created TIMESTAMP NOT NULL,
    updated TIMESTAMP NOT NULL,
    UNIQUE(provider, provider_id)
);

CREATE TABLE IF NOT EXISTS transactions (
    id BIGSERIAL PRIMARY KEY,
    user_id UUID NOT NULL,
    amount DECIMAL(19, 4) NOT NULL,
    description VARCHAR(512) NOT NULL,
    incoming BOOLEAN NOT NULL,
    type VARCHAR(32) NOT NULL,
    recurring BOOLEAN NOT NULL,
    start_date TIMESTAMP NOT NULL,
    end_date TIMESTAMP NOT NULL,
    interval VARCHAR(16),
    days_interval INTEGER,
    created TIMESTAMP NOT NULL,
    updated TIMESTAMP NOT NULL,
    CONSTRAINT fk_user_id
        FOREIGN KEY(user_id)
        REFERENCES users(id) 
        ON DELETE CASCADE
);

-- +goose Down
DROP TABLE transactions;
DROP TABLE users;
