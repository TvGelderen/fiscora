-- +goose Up
ALTER TABLE users ADD COLUMN avatar VARCHAR(128);

CREATE TABLE IF NOT EXISTS budgets (
    id SERIAL PRIMARY KEY,
    user_id UUID NOT NULL,
    type VARCHAR(32) NOT NULL,
    name VARCHAR(100) NOT NULL,
    description VARCHAR(256) NOT NULL,
    amount DECIMAL(19, 4) NOT NULL,
    start_date TIMESTAMP NOT NULL,
    end_date TIMESTAMP NOT NULL,
    created TIMESTAMP NOT NULL,
    updated TIMESTAMP NOT NULL,
    CONSTRAINT fk_user_id
        FOREIGN KEY(user_id)
        REFERENCES users(id) 
        ON DELETE CASCADE
);

CREATE TABLE budget_categories (
    id SERIAL PRIMARY KEY,
    budget_id INT NOT NULL,
    name VARCHAR(64) NOT NULL,
    allocated_amount DECIMAL(10, 2) NOT NULL,
    CONSTRAINT fk_budget_id
        FOREIGN KEY (budget_id) 
        REFERENCES budgets(budget_id)
        ON DELETE CASCADE
);

-- +goose Down
ALTER TABLE users DROP COLUMN avatar;

DROP TABLE budgets;
