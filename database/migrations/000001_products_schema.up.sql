-- Create table --
CREATE TABLE IF NOT EXISTS products(
    id              BIGSERIAL NOT NULL,
    name            VARCHAR(50) NOT NULL,
    supplier_id     INT NOT NULL,
    category_id     INT NOT NULL,
    stock           INT NOT NULL,
    price           DECIMAL(6, 2),
    discontinued    BOOLEAN,
    CONSTRAINT products_id_pk PRIMARY KEY (id)
);