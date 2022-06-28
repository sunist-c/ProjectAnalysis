# 1. 架构

## 1.1 Overview

![Structure - ProjcetAnalysis](https://s3.bmp.ovh/imgs/2022/06/27/8296f0d2497c9a00.png)

## 1.2 Frontend - 展示数据

+ 使用Vue.js构建Web
+ 使用Echarts执行图形可视化
+ 使用Material作为UI框架

## 1.3 Backend - 响应前端/调起大数据服务

+ 使用go-gin构建http服务
+ 使用go-zero拆分微服务
+ 使用OAuth2.0进行客户端认证
+ 使用Kafka进行日志收集
+ 使用Redis和ClickHouse作为临时数据存储
+ 使用Mysql作为持久化数据存储
+ 使用cron定时处理scrapy爬虫任务

## 1.4 Big Data - 大数据离线处理

+ 使用SpringRPC进行rpc调用
+ 使用Hive作为大数据离线处理工具
+ 使用Mysql作为持久化数据存储
+ 使用HBase作为数据仓库

# 2. 功能点 

## 2.1 前端功能点

### Overview

> 竞品参考： [COVID-19 Map - Johns Hopkins Coronavirus Resource Center (jhu.edu)](https://coronavirus.jhu.edu/map.html)

![Frontend Overview - Johns Hopkins University](https://s3.bmp.ovh/imgs/2022/06/27/b1f86f00cf6eba2c.png)

+ 功能1: 地图可视化
  + 分为国际地图、国内地图、城市地图三个层级，每个层级分别可视化展示热力图或者散点图
  + 国际地图点击中国国界可以跳转到国内地图
  + 国内地图点击省份可以跳转到城市地图
  + 最高精度精确到城市/区
+ 功能2: 图形可视化
  + 随地图可视化功能同步变化，用户可选展示信息
  + 可展示信息如下：
    + 当前区域每周死亡/每周感染/每周治愈病例
    + 当前区域每天死亡/每天感染/每天治愈病例
+ 功能3: TopN图表
  + 随地图可视化功能同步变化，用户可选展示信息
  + 可展示信息如下：
    + 当前区域感染TopN
    + 当前区域治愈TopN
    + 当前区域死亡TopN
+ 功能4: 当前区域信息展示
  + 随地图可视化功能同步变化，向用户展示以下信息：
    + 当前区域累计病例
    + 当前区域累计死亡
    + 当前区域累计治愈
    + 当前区域本月病例
    + 当前区域本月死亡
    + 当前区域本月治愈
    + 当前区域每日病例
    + 当前区域每日死亡
    + 当前区域每日治愈
+ 功能5: 当前区域新闻展示
  + 随地图可视化功能同步变化，向用户展示当前区域的疫情相关新闻与摘要

## 2.2 后端功能点

+ 功能1: 基于OAuth2.0协议和JWT进行客户端验证与用户验证
+ 功能2: 基于Kafka进行日志收集和消息分发
+ 功能3: 基于gRPC协议和go-zero/go-gin框架，提供http服务，并使用Docker构建微服务集群
+ 功能4: 基于scrapy和cron进行定期数据爬取数据，使用rpc调用大数据端进行离线数据处理
+ 功能5: 使用redis和ClickHouse进行数据缓存

## 2.3 大数据端功能点

+ 功能1: 基于爬虫数据计算并更新地图数据

# 3. 接口

## 3.1 大数据端接口

1. UpdateData: 更新特定区域的数据接口(grpc)

```protobuf
syntax = "proto3";

option java_package = "cn.sunist.project-analysis.grpc";
option java_multiple_file = false;
option java_outer_classname = "UpdateDataService";

// 更新区域数据rpc请求
message UpdateDataRequest {
	message MapData {
		string LocationID = 1;
		string LocationType = 2;
	}
	int32 RequestTime = 1;
	repeated MapData MapRequest = 2;
}

// 更新区域数据rpc响应
message UpdateDataResponse {
	message MapData {
		string LocationID = 1;
		string LocationType = 2;
		bool LocationUpdateStatus = 3;
	}
	repeated MapData MapResponse = 1;
}

// 更新区域数据rpc服务: 根据时间和地点列表，更新全部请求区域的数据，并将处理结果返回
service UpdateData {
	rpc updateData(UpdateDataRequest) returns (UpdateDataResponse);
}
```

## 3.2 后端接口

1. OAuth模块
   1. NewClient: 新建OAuth客户端(http)
   2. NewUser: 新建OAuth用户(http)
   3. Token: 申请OAuth的Token(http)
   4. RefreshToken: 申请OAuth的RefreshToken(http)
   5. Introspect: 验证OAuth的Token(http)
2. Overview模块
   1. MapOverview: 获取某区域的疫情概况(http)
3. Statistics模块
   1. MapStatistic: 获取某区域的疫情统计(http)
4. Cron模块
   1. Process: 执行Cron任务(http)
   2. Report: 查看Cron任务的执行状况(http)
   3. Options: 更改Cron任务配置(http)
5. Trace模块
   缺省，有空再做

> Author of this page: [SunistC](https://www.sunist.cn/post/ProjectAnalysisDesignDocument) - 28, Jun, 2022