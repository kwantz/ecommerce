USE ecommerce_account;

CREATE TABLE IF NOT EXISTS account (
    id        INT           NOT NULL AUTO_INCREMENT,
    password  VARCHAR(255)  NOT NULL,
    email     VARCHAR(255)  NOT NULL,
    phone     VARCHAR(255)  NOT NULL,
    address   VARCHAR(255)  NOT NULL,
    PRIMARY KEY (id)
);