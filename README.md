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
- [Endpoint](#endpoint)
    
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
1. Install Docker + Docker Compose (Docker Desktop)
2. Running Docker Desktop until the Engine Running (Green Highlight)
3. docker-compose up -d
```

# Endpoint

```
Account Service (http://localhost:8081)
GET /ping       - Service Health Check
GET /           - Get All Accounts
POST /          - Create New Account
POST /login     - Do Authentication
POST /authorize - Do Authorization

Product Service (http://localhost:8082)
GET /ping - Service Health Check
GET /     - Get All Products
POST /    - Create New Product

Order Service (http://localhost:8083)
GET /ping
```