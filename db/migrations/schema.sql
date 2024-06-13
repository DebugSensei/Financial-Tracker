DO
$$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_database WHERE datname = 'financial_tracker') THEN
        CREATE DATABASE financial_tracker;
    END IF;
END
$$;

\c financial_tracker;

CREATE TABLE IF NOT EXISTS categories (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL
);

CREATE TABLE IF NOT EXISTS transactions (
    id SERIAL PRIMARY KEY,
    date TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    amount DECIMAL(10, 2) NOT NULL,
    currency VARCHAR(3) NOT NULL,
    type VARCHAR(10) NOT NULL,
    category_id INT REFERENCES categories(id)
);

INSERT INTO categories (name) VALUES
('Food'),
('Transport'),
('Entertainment'),
('Utilities');
