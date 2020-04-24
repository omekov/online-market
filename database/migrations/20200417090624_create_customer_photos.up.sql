CREATE TABLE customer_photos(
    id SERIAL PRIMARY KEY,
    URL VARCHAR(50),
    customer_id INTEGER REFERENCES customers(id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL
);