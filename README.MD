# Odery

A digital product ordering backend service written in <b>java</b> and <b>goLang</b> with microservice architecture and a
well-thought-out system design implementation. Java services are <b>spring boot</b> while golang are <b>gin</b>

### Services

1. Auth services(GoLang)
2. Gateway service(Java)
3. User service(Java)
4. Product service(Java)
5. Order service(Java)
6. Payment service(GoLang)
7. Notification service(GoLang)
8. Discovery service(Java)

### Databases

1. Postgres
2. Mongo

User service uses postgres while notification,product,order and payment services all uses mongodb

### Cache

1. Redis

Used in User and Product services

### Messaging

1. RabbitMq

Implemented in product,order,payment and notification services

### Microservice diagram

<img src="assets/ordery1.jpg">

### System design explanation

1. Gateway service: Api Gateway redirects all request to appropriate endpoint both open and authenticated endpoints.
2. User service: User create account on user service. postgresql was used.
3. Auth service: Logs in on auth service using username and password gets jwt token. with direct db call to users
   database for latency.
4. Product service: server the client on products details mongodb was used.
5. Order service: post order saves and get posted out through RabbitMQ used mongodb receives payment and update order
   status from same RabbitMQ.
6. Payment service: Receives from RabbitMQ order and process payment sends queue for orders saves with  mongodb
7. Notification service:  Receives from RabbitMQ on Payment and Order uses mongodb
8. Discovery service: Uses eureka for server discovery
