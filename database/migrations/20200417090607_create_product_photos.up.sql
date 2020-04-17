CREATE TABLE product_photos(
    id SERIAL PRIMARY KEY,
    url VARCHAR(50),
    product_id integer REFERENCES products(id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL
);