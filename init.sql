CREATE TABLE IF NOT EXISTS products (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    price DECIMAL(10,2) NOT NULL
);

INSERT INTO products (id, name, price) VALUES
(1, 'Lipstick', 15.99),
(2, 'Mascara', 12.50),
(3, 'Foundation', 23.75),
(4, 'Eyeliner', 9.99),
(5, 'Blush', 14.25),
(6, 'Eyeshadow Palette', 29.99),
(7, 'Concealer', 11.49),
(8, 'Highlighter', 18.99),
(9, 'Bronzer', 17.50),
(10, 'Setting Spray', 13.99);