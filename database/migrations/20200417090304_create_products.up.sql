CREATE TABLE products(
    id SERIAL PRIMARY KEY,
    name VARCHAR(50),
    description TEXT,
    price DECIMAL(5,2),
    category_id INTEGER REFERENCES categories(id),
    stock_id INTEGER REFERENCES stocks(id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP
);