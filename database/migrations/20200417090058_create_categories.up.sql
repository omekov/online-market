CREATE TABLE categories(
    id SERIAL PRIMARY KEY,
    name VARCHAR(25),
    description TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP
);