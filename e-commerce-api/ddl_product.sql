CREATE DATABASE sbfcase_products;

CREATE TABLE sbfcase_products.products
(
    product_id   UUID PRIMARY KEY,
    name String NOT NULL,
    price DECIMAL NOT NULL
);