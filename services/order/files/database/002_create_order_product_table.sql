USE ecommerce_order;

CREATE TABLE IF NOT EXISTS order_product (
    id          INT  NOT NULL AUTO_INCREMENT,
    order_id    INT  NOT NULL,
    product_id  INT  NOT NULL,
    quantity    INT  NOT NULL,
    price       INT  NOT NULL,
    PRIMARY KEY (id)
);