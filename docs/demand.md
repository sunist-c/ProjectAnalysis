# System Demand

## Covid-19 Data Visualization

### 1. Map data visualization

+ Divided into three levels: international map, domestic map, and city map, each level visualizes heat map or scatter plot respectively
+ International map Click on the border of China to jump to the domestic map
+ Domestic map Click on the province to jump to the city map
+ Highest accuracy down to city/district

### 2. Graph Visualization

+ Changes synchronously with the map visualization function, users can choose to display information
+ The information that can be displayed is as follows:
    + Weekly deaths/weekly infections/weekly cured cases in current region
    + Daily deaths/daily infections/daily cured cases in the current region

### 3. TopN Chart

+ Changes synchronously with the map visualization function, users can choose to display information
+ The information that can be displayed is as follows:
    + The current area is infected with TopN
    + Current area heals TopN
    + Current area death TopN

### 4. Display of current area information

+ Changes synchronously with the map visualization function, showing the following information to the user:
    + Cumulative cases in the current area
    + Cumulative deaths in current area
    + Cumulative healing in the current area
    + Cases of the month in the current region
    + Current region died this month
    + Current area healed this month
    + Current area daily cases
    + daily deaths in current area
    + Daily heal in current area

### 5. Current regional news display

+ Simultaneously changes with the map visualization function, showing users the news and summaries of the current area of ​​the epidemic

## Contribution System

+ 功能1: 基于OAuth2.0协议和JWT进行客户端验证与用户验证
+ 功能2: 基于Kafka进行日志收集和消息分发
+ 功能3: 基于gRPC协议和go-zero/go-gin框架，提供http服务，并使用Docker构建微服务集群
+ 功能4: 基于scrapy和cron进行定期数据爬取数据，使用rpc调用大数据端进行离线数据处理
+ 功能5: 使用redis和ClickHouse进行数据缓存

## Big Data Algorithm

+ 功能1: 基于爬虫数据计算并更新地图数据