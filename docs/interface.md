# Interface Design

## 1. HTTP Service

### 1.1 Authentication

Authentication module is used for authorizing if client/user/request is invalid. Here is the overview of authentication progress:

![Authentication Overview](/_media/interface_1.jpg)

1. Request AuthorizeCode

2. Request AccessToken with AuthorizeCode

3. Request AccessToken with RefreshToken

4. Request AccessToken directly

5. Introspect AccessToken

6. Create Client

7. Create Users

### Covid-19 Visualization Data

Visualization Data module is used for frontend to visualize the covid-19 situation. Here is the overview of visualization progress:

![Authentication Overview](/_media/interface_2.png)

1. Get Visualization Data

2. Get Raw Data

3. Calculate Visualization Data

4. Statistic Visualization Data

## 1.2 Data Spider

## 1.3 Kafka Message

## 1.4 Hive Implement

> Author of this page: [SunistC](https://www.sunist.cn) - 29, Jun, 2022 (Updated at 30, Jun, 2022)