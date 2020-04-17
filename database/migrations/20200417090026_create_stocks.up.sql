CREATE TABLE stocks(
    id SERIAL PRIMARY KEY,
    name VARCHAR(25),
    description TEXT,
    precent NUMERIC(10,2),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP
);