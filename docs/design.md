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
3. Mysql/Redis -[response]-> Authenticator: database return the result of the query requseted
4. Authenticator -[next]-> QPS Gateway: authenticator complete the authorize process, press the context to QPS Gateway
5. QPS Gateway -[next]-> Handler: QPS Gateway limited the qps of client/user, press the context to service handler
6. Handler -[query]-> ClickHouse: service handler do actions to ClickHouse database with logical operation
7. ClickHouse -[response]-> Handler: ClickHouse return the result of the query/action requested
8. Handler -[response]-> Frontend: service handler return the http response to frontend

### 2.2 Cron Task -> Hive Implement Calling

The Cron Task -> Hive Implement calling process is like the following graphic:

![Cron Task -> Hive Implement Calling Overview](/_media/structure_2.png)

1. CRON -[time tick]-> Processor: the cron task wakes up the processor which process the cron task
2. Processor -[exec]-> Data Spider: the processor will execute spiders to get covid-19 data
3. Data Spider -[return]-> Processor: the spider return the result obtained
4. Processor -[call]-> Clean & Pre-Processor: the processor will press the data to data cleaner and data pre-processor to clean the dirty data and complete the pre-process progress
5. Clean & Pre-Processor -[write]-> Kafka: the cleaner and pre-processor will press the data to Kafka with writing message
6. Kafka -[push]-> Hive: Kafka will send the message which pushed from data cleaner and pre-processor to Hive
7. Hive <-[calculate]-> Hive: Hive will calculate the huge-count data with HBase
8. Hive -[write]-> Kafka: Hive will press the data calculated to Kafka with writing message
9. Kafka -[push]-> Processor: Kafka will send the message which pushed from Hive to ClickHouse Processor
10. Processor -[update]-> ClickHouse: Processor do actions to ClickHouse with logical operation

> Author of this page: [SunistC](https://www.sunist.cn/post/ProjectAnalysisDesignDocument) - 28, Jun, 2022 (Updated at 30, Jun, 2022)