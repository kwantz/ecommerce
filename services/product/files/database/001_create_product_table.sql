USE ecommerce_product;

CREATE TABLE IF NOT EXISTS product (
    id     INT           NOT NULL AUTO_INCREMENT,
    name   VARCHAR(255)  NOT NULL,
    stock  INT           NOT NULL,
    price  INT           NOT NULL,
    PRIMARY KEY (id)
);