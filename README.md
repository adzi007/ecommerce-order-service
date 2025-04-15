# 📦 E-Commerce Order Service
A backend microservice that handles order creation and updates in an e-commerce system. Built with Go and Fiber using Clean Architecture. It communicates with the Cart Service via gRPC and sends notifications via RabbitMQ.

---
## 🚀 Features
- Create new customer orders
- Update order status (e.g. paid, shipped, canceled)
- Validates cart contents through gRPC with Cart Service
- Publishes order events to RabbitMQ for Notification Service
- Follows Clean Architecture for maintainability and testability
---
## 🧰 Tech Stack
- **Language**: Golang (Go 1.21+)
- **Framework**: [Fiber](https://gofiber.io/)
- **Architecture**: Clean Architecture (Handler → Service → Repository)
- **Database**: MySQL
- **Messaging**: RabbitMQ (as publisher)
- **Communication**: gRPC (to Cart Service)
- **Containerization**: Docker + Docker Compose
---
## 📁 Project Structure

```md
├── cart_proto
├── cmd
│   └── server
├── config
├── internal
│   ├── delivery
│   │   ├── grpc_client
│   │   ├── grpc_handler
│   │   └── http
│   ├── domain
│   ├── infrastructure
│   │   ├── database
│   │   ├── grpc_connection
│   │   ├── logger
│   │   ├── rabbitmq
│   │   └── redis
│   ├── migration
│   ├── model
│   ├── repository
│   │   ├── cachestore
│   │   ├── cart_service_repository
│   └── service
├── pkg
│   └── http_client
├── proto
└── tmp
```

## Getting Started
### Prerequisites
- Docker
- Go 1.21+
- <a href="https://github.com/adzi007/ecommerce-products-service" target="_blank">Ecommerce Product Service</a>
- <a href="https://github.com/adzi007/eccomerce-cart-service" target="_blank">Ecommerce Cart Service</a>
- RabbitMQ Container in the same network

## Running Locally (Docker)

1. Clone the project

```bash
git clone https://github.com/adzi007/ecommerce-order-service.git
cd ecommerce-order-service
```

2. CD into the ecommerce-cart-service directory and create an .env file or edit from .env.example following with fields bellow

```
PORT_APP=5002
DB_HOST=ecommerce-order-db
DB_PORT=5432
DB_USERNAME=YOUR_DB_USERNAME
DB_PASSWORD=YOUR_DB_PASSWORD
DB_NAME=ecommerce_app
API_GATEWAY=
PRODUCT_SERVICE_PATH=
RABBITMQ_USER=guest
RABBITMQ_PASS=guest
RABBITMQ_HOST=
RABBITMQ_PORT=5672
RABBITMQ_VHOST=
```

3. Build container

```
docker-compose up --build
```

The App will be running at `http://localhost:5002`
## Database Migration
1. Execute migration database
```
docker exec -it ecommerce-order-service /migrate
```

## API Documentation

<a href="https://www.postman.com/adzi/ecommerce-order-service" target="_blank">Postman Collections</a>