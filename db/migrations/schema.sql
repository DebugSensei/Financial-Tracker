-- db/migrations/schema.sql
CREATE TABLE IF NOT EXISTS categories (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS transactions (
    id SERIAL PRIMARY KEY,
    date TIMESTAMP NOT NULL,
    amount NUMERIC(10, 2) NOT NULL,
    currency VARCHAR(3) NOT NULL,
    type VARCHAR(10) NOT NULL,
    category_id INTEGER NOT NULL REFERENCES categories(id)
);

INSERT INTO categories (name) VALUES ('Shopping'), ('Groceries'), ('Utilities'), ('Entertainment') ON CONFLICT (name) DO NOTHING;

