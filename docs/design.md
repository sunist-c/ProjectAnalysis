# Structure Design

## 1. Domain-Driven Design

We devided our system to four parts via ddd, the following graphic is the structure overview:

![Domain-Driven Design Structure Overview](/_media/structure_3.png)

### 1.1 User Interface

This layer is designed for public service:

- Web Frontend: the visualization frontend pressed by website frontend
- HTTP Interface: the http service provide for frontend and third-party system to request our data
- Manage Interface: the http service provide for system managers to manage the whole system

### 1.2 Application

This layer contains the abstract service in our system:

- Visualization Data Processor: the service which provide data for visualization frontend
- Data Spider: the service which get data for third-party website and database via spider
- BigData Processor: the service which process the huge-count data to structured data for backend via hadoop

### 1.3 Domain

This layer is the logical implement of our system:

- Hive Implement: the implement of hive which calculate the huge-count data
- Backend Handler: the implement of backend which provide backend service

### 1.4 Infrastructure

This layer contains the basic components of our system:

- Mysql: the database which stored system data
- Redis: the cacher which cached system data
- Kafka: the message-queue which delivered message between components
- ClickHouse: the database which stored system data
- CeylonFramework: the backend framework which provide http service, database access, cron task, log agent, authentication, gateway, .etc
- Hadoop: the big data framework which provide big data calculating

## 2. Component Calling

### 2.1 Frontend -> Backend Calling

The Frontend -> Backend request and response process is like the following graphic:

![Frontend -> Backend Calling Overview](/_media/structure_1.png)

1. Frontend -[request]-> Authenticator: authenticator will receive the request
2. Authenticator -[query]-> Mysql/Redis: authenticator will ask for mysql/redis to query the user-info and client-info
3. Mysql/Redis -[response]-> Authenticator: 

### 2.2 Cron Task -> Hive Implement Calling

The Cron Task -> Hive Implement calling process is like the following graphic:

![Cron Task -> Hive Implement Calling Overview](/_media/structure_2.png)

> Author of this page: [SunistC](https://www.sunist.cn/post/ProjectAnalysisDesignDocument) - 29, Jun, 2022