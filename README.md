This repo demonstrates a microservices architecture using gRPC for inter-service communication and GraphQL as the API gateway. 
It includes services for account management, product catalog, and order processing.


The project consists of the following main components:

    Account Service
    Catalog Service
    Order Service
    GraphQL API Gateway

Each service has its own database:

    Account and Order services use PostgreSQL
    Catalog service uses Elasticsearch
