# E-Commerce App

## Overview
This is a simple e-commerce application built with Go (Golang). It includes the following core features:

- User authentication and management
- Product catalog
- Checkout process
- Order management

The app uses PostgreSQL as the database, managed with Docker.

---

## Features

### 1. User Authentication
- Secure user registration and login.
- Password hashing for security.
- JWT-based authentication for protected routes.

### 2. Product Management
- Create, read, update, and delete products.
- Fetch product details and list all available products.

### 3. Checkout
- Add products to the cart.
- Calculate the total price, including taxes or discounts.
- Validate user inputs during the checkout process.

### 4. Order Management
- Place orders.
- Fetch order history for authenticated users.
- Manage order statuses (e.g., pending, completed, canceled).

---

## Tech Stack

- **Backend**: Go (Golang)
- **Database**: PostgreSQL
- **Containerization**: Docker

---

## Getting Started

### Prerequisites
Ensure you have the following installed:
- [Go](https://golang.org/doc/install)
- [Docker](https://docs.docker.com/get-docker/)
- [PostgreSQL](https://www.postgresql.org/download/) (optional, for local development without Docker)

### Setup Instructions

1. **Clone the Repository**
   ```bash
   git clone <repository-url>
   cd ecommerce-golang
   ```

2. **Set Up Environment Variables** Create a .env file in the project root with the following variables:
    ```bash
    DB_ADDR=
    DB_MAX_OPEN_CONNS=
    DB_MAX_IDLE_CONNS=
    DB_MAX_IDLE_TIME=
    JWT_SECRET=
    ```

3. **Run PostgreSQL with Docker**
    ```bash
    docker-compose up -d
    ```

4. **Run the Application**
    ```bash
    cd cmd
    air
    ```

5. **Access the App** The application will be available at **http://localhost:8081**
