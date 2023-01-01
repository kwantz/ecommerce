USE ecommerce_order;

CREATE TABLE IF NOT EXISTS orders (
    id               INT           NOT NULL AUTO_INCREMENT,
    account_id       INT           NOT NULL,
    status           VARCHAR(255)  NOT NULL,
    payment_status   VARCHAR(255)  NOT NULL,
    shipping_status  VARCHAR(255)  NOT NULL,
    PRIMARY KEY (id)
);