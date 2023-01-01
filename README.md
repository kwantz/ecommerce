# Table of Contents

- [Usecases](#usecases)
- [Architecture](#architecture)
    - [Monolithic](#monolithic)
    - [Microservice (simple)](#microservice-simple)
    - [Microservice (ideal)](#microservice-ideal)
    - [Data Model](#data-model)
    - [Docker Services](#docker-services)
- [What's Next?](#whats-next)
    - [Authentication Service](#authentication-service)
    - [Account Service](#account-service)
    - [Product Service](#product-service)
    - [Cart Service](#cart-service)
    - [Order Service](#order-service)
    - [Payment Service](#payment-service)
    - [Shipping Service](#shipping-service)
- [Setup](#setup)
- [Flows Ping](#flows-ping)
- [Flows Create Account and Products](#flows-create-account-and-products)
- [Flows Add Product to Cart and Move to Order #1](#flows-add-product-to-cart-and-move-to-order-1)
- [Flows Add Product to Cart and Move to Order #2](#flows-add-product-to-cart-and-move-to-order-2)
    
# Usecases

Reference: [http://www.itk.ilstu.edu/faculty/bllim/wwwdev/sample1.htm](http://www.itk.ilstu.edu/faculty/bllim/wwwdev/sample1.htm)

1. Visit a site
2. Login
3. Search for Products
4. Check Availability of Products
5. Add Products to Shopping Cart
6. Edit Products in the Shopping Cart
7. remove Products from the Shopping Cart
8. heck out the Shopping Cart
9. Check Order Status
10. Browse Order History
11. Display Product Details
12. Logout
13. Leave a site

# Architecture

### Monolithic

![architecture_1](./documentations/images/architecture_1.jpg?raw=true)


### Microservice (simple)

![architecture_2](./documentations/images/architecture_2.jpg?raw=true)


### Microservice (ideal)

![architecture_3](./documentations/images/architecture_3.jpg?raw=true)

### Data Model

![data_model](./documentations/images/data_model.jpg?raw=true)

### Docker Services

![docker_services](./documentations/images/docker_services.jpg?raw=true)

# What's Next?

### Authentication Service
- Split Authentication from Account Service
- Handle Basic Auth
- Handle Oauth like Google, FB, Twitter, etc.
- Handle Token Auth
- Handle Limit Failed Auth Attempts

### Account Service
- Edit (Change, Delete) Account Info
- Reporting Order, Shipping, Payment via Email
- Secure Personally Identifiable Information (PII) like email, phone, address, etc.
- Use synchronous encryption for internal service purposes
- Use asynchronous encryption for external purposes (client, user, partner, etc.)

### Product Service
- Maintain (Add, Change, Delete) Inventory
- Create Product Review
- Create Product Discussion

### Cart Service
- Split Cart from Product Service
- Focus on Carting and Wishlist

### Order Service
- View Order Summary & Detail
- Change Order ID into Invoice Number
- View History Date of Status Changes (Order, Payment, Shipping)

### Payment Service
- Split Payment from Order Service
- Focus Handle Third-party Payment API like BCA, BRI, BNI, Gopay, Ovo, etc.

### Shipping Service
- Split Shipping from Order Service
- Focus Handle Third-party Shipping API like JNE, JNT, Gojek, Grab, etc.

# Setup

```
1. Install Go Language
2. Install Docker + Docker Compose (Docker Desktop)
3. Running Docker Desktop until the Engine Running (Green Highlight)
4. Open CLI, change to project directory, and run :
   $ docker-compose up -d --build
```

# Flows Ping

## Ping Account

### Request

`GET /ping`

```curl
curl --location --request GET 'localhost:8081/ping'
```

### Response

```json
{
    "message": "Hello World from Account Service"
}
```

## Ping Product

### Request

`GET /ping`

```curl
curl --location --request GET 'localhost:8082/ping'
```

### Response

```json
{
    "message": "Hello World from Product Service"
}
```

## Ping Order

### Request

`GET /ping`

```curl
curl --location --request GET 'localhost:8083/ping'
```

### Response

```json
{
    "message": "Hello World from Order Service"
}
```

# Flows Create Account and Products

## Create Account

### Request

`POST /`

```curl
curl --location --request POST 'localhost:8081/' \
--header 'Content-Type: application/json' \
--data-raw '{
    "email": "user@test.com",
    "password": "password",
    "phone": "08123456678",
    "address": "JL. Thamrin"
}'
```

### Response

```json
{
    "id": 1,
    "email": "user@test.com",
    "phone": "08123456678",
    "address": "JL. Thamrin"
}
```

##  Create Product #1

### Request

`POST /`

```curl
curl --location --request POST 'localhost:8082/' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name": "Product Test #1",
    "stock": 100,
    "price": 1000
}'
```

### Response

```json
{
    "id": 1,
    "name": "Product Test #1",
    "stock": 100,
    "price": 1000
}
```

##  Create Product #2

### Request

`POST /`

```curl
curl --location --request POST 'localhost:8082/' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name": "Product Test #2",
    "stock": 200,
    "price": 2000
}'
```

### Response

```json
{
    "id": 2,
    "name": "Product Test #2",
    "stock": 200,
    "price": 2000
}
```

##  Create Product #3

### Request

`POST /`

```curl
curl --location --request POST 'localhost:8082/' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name": "Product Test #3",
    "stock": 300,
    "price": 3000
}'
```

### Response

```json
{
    "id": 3,
    "name": "Product Test #3",
    "stock": 300,
    "price": 3000
}
```

##  Create Product #4

### Request

`POST /`

```curl
curl --location --request POST 'localhost:8082/' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name": "Product Test #4",
    "stock": 400,
    "price": 4000
}'
```

### Response

```json
{
    "id": 4,
    "name": "Product Test #4",
    "stock": 400,
    "price": 4000
}
```

##  Create Product #5

### Request

`POST /`

```curl
curl --location --request POST 'localhost:8082/' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name": "Product Test #5",
    "stock": 500,
    "price": 5000
}'
```

### Response

```json
{
    "id": 5,
    "name": "Product Test #5",
    "stock": 500,
    "price": 5000
}
```

## Get All Products

### Request

`GET /`

```curl
curl --location --request GET 'localhost:8082/'
```

### Response

```json
[
    {
        "id": 1,
        "name": "Product Test #1",
        "stock": 100,
        "price": 1000
    },
    {
        "id": 2,
        "name": "Product Test #2",
        "stock": 200,
        "price": 2000
    },
    {
        "id": 3,
        "name": "Product Test #3",
        "stock": 300,
        "price": 3000
    },
    {
        "id": 4,
        "name": "Product Test #4",
        "stock": 400,
        "price": 4000
    },
    {
        "id": 5,
        "name": "Product Test #5",
        "stock": 500,
        "price": 5000
    }
]
```

# Flows Add Product to Cart and Move to Order #1

## Authentication

### Request

`POST /login`

```curl
curl --location --request POST 'localhost:8081/login' \
--header 'Content-Type: application/json' \
--data-raw '{
    "email": "user@test.com",
    "password": "password"
}'
```

### Response

```json
{
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJBQ0NPVU5UX1NFUlZJQ0VfSVNTVUVSIiwiZXhwIjoxNjcyNTcyMTYwLCJpZCI6MSwiZW1haWwiOiJ1c2VyQHRlc3QuY29tIiwicGhvbmUiOiIwODEyMzQ1NjY3OCIsImFkZHJlc3MiOiJKTC4gVGhhbXJpbiJ9.MV4j709mXainNlGf4shbfmt07tGBsKisJ3TQtgzjlTg"
}
```

## Add Product to Cart #1

### Request

`POST /cart`

```curl
curl --location --request POST 'localhost:8082/cart' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer <token>' \
--data-raw '{
    "product_id": 1,
    "quantity": 10
}'
```

### Response

```json
{
    "id": 1,
    "account_id": 1,
    "product_id": 1,
    "quantity": 10
}
```

## Add Product to Cart #2

### Request

`POST /cart`

```curl
curl --location --request POST 'localhost:8082/cart' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer <token>' \
--data-raw '{
    "product_id": 2,
    "quantity": 20
}'
```

### Response

```json
{
    "id": 2,
    "account_id": 1,
    "product_id": 2,
    "quantity": 20
}
```

## Add Product to Cart #3

### Request

`POST /cart`

```curl
curl --location --request POST 'localhost:8082/cart' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer <token>' \
--data-raw '{
    "product_id": 3,
    "quantity": 30
}'
```

### Response

```json
{
    "id": 3,
    "account_id": 1,
    "product_id": 3,
    "quantity": 30
}
```

## Move Cart to Order

### Request

`POST /`

```curl
curl --location --request POST 'localhost:8083/' \
--header 'Authorization: Bearer <token>'
```

### Response

```json
{
    "invoice": "7c338f43-6030-4e4f-a934-29592cba7f7d",
    "status": "PENDING",
    "payment_status": "AWAITING PAYMENT",
    "shipping_status": "AWAITING SHIPPING"
}
```

## Get Detail Order

### Request

`GET /<income>`

```curl
curl --location --request GET 'localhost:8083/<invoice>' \
--header 'Authorization: Bearer <token>'
```

### Response

```json
{
    "invoice": "7c338f43-6030-4e4f-a934-29592cba7f7d",
    "status": "PENDING",
    "payment_status": "AWAITING PAYMENT",
    "shipping_status": "AWAITING SHIPPING",
    "list_order_product": [
        {
            "id": 1,
            "product_id": 1,
            "product_name": "Product Test #1",
            "quantity": 10,
            "price": 1000
        },
        {
            "id": 2,
            "product_id": 2,
            "product_name": "Product Test #2",
            "quantity": 20,
            "price": 2000
        },
        {
            "id": 3,
            "product_id": 3,
            "product_name": "Product Test #3",
            "quantity": 30,
            "price": 3000
        }
    ]
}
```

## Get All Products

### Request

`GET /`

```curl
curl --location --request GET 'localhost:8082/'
```

### Response

```json
[
    {
        "id": 1,
        "name": "Product Test #1",
        "stock": 90,
        "price": 1000
    },
    {
        "id": 2,
        "name": "Product Test #2",
        "stock": 180,
        "price": 2000
    },
    {
        "id": 3,
        "name": "Product Test #3",
        "stock": 270,
        "price": 3000
    },
    {
        "id": 4,
        "name": "Product Test #4",
        "stock": 400,
        "price": 4000
    },
    {
        "id": 5,
        "name": "Product Test #5",
        "stock": 500,
        "price": 5000
    }
]
```

# Flows Add Product to Cart and Move to Order #2

## Add Product to Cart #4

### Request

`POST /cart`

```curl
curl --location --request POST 'localhost:8082/cart' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer <token>' \
--data-raw '{
    "product_id": 4,
    "quantity": 40
}'
```

### Response

```json
{
    "id": 4,
    "account_id": 1,
    "product_id": 4,
    "quantity": 40
}
```

## Add Product to Cart #5

### Request

`POST /cart`

```curl
curl --location --request POST 'localhost:8082/cart' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer <token>' \
--data-raw '{
    "product_id": 5,
    "quantity": 50
}'
```

### Response

```json
{
    "id": 5,
    "account_id": 1,
    "product_id": 5,
    "quantity": 50
}
```

## Move Cart to Order

### Request

`POST /`

```curl
curl --location --request POST 'localhost:8083/' \
--header 'Authorization: Bearer <token>'
```

### Response

```json
{
    "invoice": "277a73ac-d8c8-486d-ba1e-8e46562c6e00",
    "status": "PENDING",
    "payment_status": "AWAITING PAYMENT",
    "shipping_status": "AWAITING SHIPPING"
}
```

## Get Detail Order

### Request

`GET /<income>`

```curl
curl --location --request GET 'localhost:8083/<invoice>' \
--header 'Authorization: Bearer <token>'
```

### Response

```json
{
    "invoice": "277a73ac-d8c8-486d-ba1e-8e46562c6e00",
    "status": "PENDING",
    "payment_status": "AWAITING PAYMENT",
    "shipping_status": "AWAITING SHIPPING",
    "list_order_product": [
        {
            "id": 4,
            "product_id": 4,
            "product_name": "Product Test #4",
            "quantity": 40,
            "price": 4000
        },
        {
            "id": 5,
            "product_id": 5,
            "product_name": "Product Test #5",
            "quantity": 50,
            "price": 5000
        }
    ]
}
```

## Get All Order

### Request

`GET /`

```curl
curl --location --request GET 'localhost:8083/' \
--header 'Authorization: Bearer <token>'
```

### Response

```json
[
    {
        "invoice": "7c338f43-6030-4e4f-a934-29592cba7f7d",
        "status": "PENDING",
        "payment_status": "AWAITING PAYMENT",
        "shipping_status": "AWAITING SHIPPING"
    },
    {
        "invoice": "277a73ac-d8c8-486d-ba1e-8e46562c6e00",
        "status": "PENDING",
        "payment_status": "AWAITING PAYMENT",
        "shipping_status": "AWAITING SHIPPING"
    }
]
```

## Get All Products

### Request

`GET /`

```curl
curl --location --request GET 'localhost:8082/'
```

### Response

```json
[
    {
        "id": 1,
        "name": "Product Test #1",
        "stock": 90,
        "price": 1000
    },
    {
        "id": 2,
        "name": "Product Test #2",
        "stock": 180,
        "price": 2000
    },
    {
        "id": 3,
        "name": "Product Test #3",
        "stock": 270,
        "price": 3000
    },
    {
        "id": 4,
        "name": "Product Test #4",
        "stock": 360,
        "price": 4000
    },
    {
        "id": 5,
        "name": "Product Test #5",
        "stock": 450,
        "price": 5000
    }
]
```
