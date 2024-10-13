-- +goose Up
CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY,
    provider VARCHAR(16) NOT NULL,
    provider_id VARCHAR(32) NOT NULL,
    username VARCHAR(32) UNIQUE NOT NULL,
    email VARCHAR(64) UNIQUE NOT NULL,
    avatar VARCHAR(128),
    password_hash BYTEA,
    created TIMESTAMP NOT NULL DEFAULT (now() at time zone 'utc'),
    updated TIMESTAMP NOT NULL DEFAULT (now() at time zone 'utc'),
    UNIQUE(provider, provider_id)
);

CREATE TABLE IF NOT EXISTS transactions (
    id SERIAL PRIMARY KEY,
    user_id UUID NOT NULL,
    amount DECIMAL(19, 4) NOT NULL,
    description VARCHAR(512) NOT NULL,
    type VARCHAR(32) NOT NULL,
    recurring BOOLEAN NOT NULL,
    start_date TIMESTAMP NOT NULL,
    end_date TIMESTAMP NOT NULL,
    interval VARCHAR(16),
    days_interval INTEGER,
    created TIMESTAMP NOT NULL DEFAULT (now() at time zone 'utc'),
    updated TIMESTAMP NOT NULL DEFAULT (now() at time zone 'utc'),
    FOREIGN KEY(user_id)
    REFERENCES users(id) 
    ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS budgets (
    id VARCHAR(16) PRIMARY KEY,
    user_id UUID NOT NULL,
    name VARCHAR(100) NOT NULL,
    description VARCHAR(256) NOT NULL,
    amount DECIMAL(19, 4) NOT NULL,
    start_date TIMESTAMP NOT NULL,
    end_date TIMESTAMP NOT NULL,
    created TIMESTAMP NOT NULL DEFAULT (now() at time zone 'utc'),
    updated TIMESTAMP NOT NULL DEFAULT (now() at time zone 'utc'),
    FOREIGN KEY(user_id)
    REFERENCES users(id) 
    ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS budget_expenses (
    id SERIAL PRIMARY KEY,
    budget_id VARCHAR(16) NOT NULL,
    name VARCHAR(64) NOT NULL,
    allocated_amount DECIMAL(19, 2) NOT NULL,
    current_amount DECIMAL(19, 2) NOT NULL DEFAULT 0,
    FOREIGN KEY (budget_id) 
    REFERENCES budgets(id)
    ON DELETE CASCADE
);

-- +goose Down
DROP TABLE budget_expenses;
DROP TABLE budgets;
DROP TABLE transactions;
DROP TABLE users;
