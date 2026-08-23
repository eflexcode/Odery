# Odery
A digital product ordering service written in java and goLang with microservice architecture
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

Used in User both Product services

### Messaging
1. RabbitMq

Implemented in product,order,payment and notification services


### Microservice diagram