CREATE TABLE cart_product(
    cart_id INTEGER REFERENCES cart(id),
    product_id INTEGER REFERENCES products(id)
);