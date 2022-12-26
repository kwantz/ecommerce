USE ecommerce_product;

CREATE TABLE IF NOT EXISTS cart (
    id          INT       NOT NULL AUTO_INCREMENT,
    account_id  INT       NOT NULL,
    product_id  INT       NOT NULL,
    quantity    INT       NOT NULL,
    deleted_at  DATETIME  DEFAULT NULL,
    PRIMARY KEY (id)
);