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

### 1.2 Covid-19 Visualization Data

Visualization Data module is used for frontend to visualize the covid-19 situation. Here is the overview of visualization progress:

![Authentication Overview](/_media/interface_2.png)

1. Get Map Data

**Description**

This interface provides map data to frontend for visualization base on location and date

**URL**

```
GET /map-data/:location/:date
```

**Params**

| Field | Description | Necessary | Example |
| :--: | :--: | :--: | :--: |
| location | the main location of the map | [x] | "world", "united_states" |
| date | the refresh date of the map data | [x] | "2022.1.1", "latest" |

**Header**

| Field | Description | Necessary | Example |
| :--: | :--: | :--: | :--: |
| AccessToken | the access-token of the frontend application | [x] | "38f3493f345c736a1e8684cd144f4d05" |

To learn more about access-token, please turn to [oauth]().

**Request Body**

None

**Example Request Body**

```json
{
  
}
```

**Response Body**

```js

```

**Example Response Body**

```json
{
  
}
```

**Error Code**



2. Get Charts Data

**Description**

This interface provides charts data to frontend for visualization base on location and date

**URL**

```
GET /charts-data/:location/:date
```

**Params**

| Field | Description | Necessary | Example |
| :--: | :--: | :--: | :--: |
| location | the main location of the map | [x] | "world", "united_states" |
| date | the refresh date of the map data | [x] | "2022.1.1", "latest" |

**Header**

| Field | Description | Necessary | Example |
| :--: | :--: | :--: | :--: |
| AccessToken | the access-token of the frontend application | [x] | "38f3493f345c736a1e8684cd144f4d05" |

To learn more about access-token, please turn to [oauth]().

**Request Body**

None

**Example Request Body**

```json
{
  
}
```

**Response Body**

```js

```

**Example Response Body**

```json
{
  
}
```

**Error Code**


3. Get Ordered Data

**Description**

This interface provides ordered data to frontend for visualization base on location and date

**URL**

```
GET /ordered-data/:location/:date
```

**Params**

| Field | Description | Necessary | Example |
| :--: | :--: | :--: | :--: |
| location | the main location of the map | [x] | "world", "united_states" |
| date | the refresh date of the map data | [x] | "2022.1.1", "latest" |

**Header**

| Field | Description | Necessary | Example |
| :--: | :--: | :--: | :--: |
| AccessToken | the access-token of the frontend application | [x] | "38f3493f345c736a1e8684cd144f4d05" |

To learn more about access-token, please turn to [oauth]().

**Request Body**

None

**Example Request Body**

```json
{
  
}
```

**Response Body**

```js

```

**Example Response Body**

```json
{
  
}
```

**Error Code**


4. Get Overview Data

**Description**

This interface provides overview data to frontend for visualization base on location and date

**URL**

```
GET /overview-data/:location/:date
```

**Params**

| Field | Description | Necessary | Example |
| :--: | :--: | :--: | :--: |
| location | the main location of the map | [x] | "world", "united_states" |
| date | the refresh date of the map data | [x] | "2022.1.1", "latest" |

**Header**

| Field | Description | Necessary | Example |
| :--: | :--: | :--: | :--: |
| AccessToken | the access-token of the frontend application | [x] | "38f3493f345c736a1e8684cd144f4d05" |

To learn more about access-token, please turn to [oauth]().

**Request Body**

None

**Example Request Body**

```json
{
  
}
```

**Response Body**

```js

```

**Example Response Body**

```json
{
  
}
```

**Error Code**



    
## 2. Data Spider

### 2.1 Raw Data

Reference: [Covid-19 Data](https://github.com/CSSEGISandData/COVID-19)

Example of the data set:

```csv
FIPS,Admin2,Province_State,Country_Region,Last_Update,Lat,Long_,Confirmed,Deaths,Recovered,Active,Combined_Key,Incident_Rate,Case_Fatality_Ratio
,,,Afghanistan,2022-01-06 04:22:09,33.93911,67.709953,158245,7367,,,Afghanistan,406.5033236325175,4.65543935037442
,,,Albania,2022-01-06 04:22:09,41.1533,20.1683,212021,3224,,,Albania,7367.468204878727,1.5206040911041832
,,,Algeria,2022-01-06 04:22:09,28.0339,1.6596,220415,6310,,,Algeria,502.6448287672428,2.862781571127192
,,,Andorra,2022-01-06 04:22:09,42.5063,1.5218,25289,141,,,Andorra,32730.214197890386,0.5575546680374867
,,,Angola,2022-01-06 04:22:09,-11.2027,17.8739,86636,1789,,,Angola,263.60157472092664,2.0649614478969482
,,,Antigua and Barbuda,2022-01-06 04:22:09,17.0608,-61.7964,4486,119,,,Antigua and Barbuda,4580.916591781717,2.652697280427998
,,,Argentina,2022-01-06 04:22:09,-38.4161,-63.6167,5915695,117346,,,Argentina,13089.043695387734,1.9836384397775748
,,,Armenia,2022-01-06 04:22:09,40.0691,45.0382,345255,7989,,,Armenia,11651.29044820625,2.3139418690533082
,,Australian Capital Territory,Australia,2022-01-06 04:22:09,-35.4735,149.0124,8021,15,,,"Australian Capital Territory, Australia",1873.6276570894647,0.18700910110958732
,,New South Wales,Australia,2022-01-06 04:22:09,-33.8688,151.2093,342133,689,,,"New South Wales, Australia",4214.49864498645,0.20138367243148134
,,Northern Territory,Australia,2022-01-06 04:22:09,-12.4634,130.8456,985,1,,,"Northern Territory, Australia",504.071661237785,0.10152284263959391
,,Queensland,Australia,2022-01-06 04:22:09,-27.4698,153.0251,46731,8,,,"Queensland, Australia",913.5177402013487,0.017119257024245146
```

### 2.2 Structured Data

Structure of the data:

```python
data.rename(
        columns={'Province/State': 'location_province', 'Province_State': 'location_province',
                 'Country/Region': 'location_country', 'Country_Region': 'location_country',
                 'Last_Update': 'refresh_time', 'Last Update': 'refresh_time', 'Confirmed': 'daily_confirm',
                 'Deaths': 'daily_death', 'Recovered': 'daily_recovered'},
        inplace=True)
```

Example of the data:

```json
{
  "data": [
    {
      "confirm":1,
      "countryName":"jdh",
      "death":2,
      "recovered":3,
      "refreshTime":1656465382198
    },
    {
      "confirm":3,
      "countryName":"China",
      "death":0,
      "recovered":0,
      "refreshTime":1656465382198
    }
  ],
  "time":7,
  "uuid":"hewfwiufwuf"
}
```

## 3. Kafka Message

## 4. Hive Implement

### 4.1 RPC Calling (Deprecated)

We use kafka to replace the rpc calling. But we still provide the protobuf file:

```protobuf
syntax = "proto3";

option java_package = "cn.sunist.project-analysis.grpc";
option java_multiple_files = false;
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

### 4.2 Listen Kafka Message

```java
public void listen(ConsumerRecord<?, ?> record) throws Exception {
    // ...implement
}
```

### 4.3 Send Kafka Message

```java
private void send(Response response) {
    // ...implement
}
```

> Author of this page: [SunistC](https://www.sunist.cn) - 29, Jun, 2022 (Updated at 30, Jun, 2022)