package main

import (
	"ProjectAnalysis/application/cron"
	"ProjectAnalysis/application/visualization"
	dto "ProjectAnalysis/domain/visualization"
	"ProjectAnalysis/infrastructure/config"
	"ProjectAnalysis/infrastructure/kafka"
	"ProjectAnalysis/infrastructure/mysql"
	"ProjectAnalysis/infrastructure/redis"
	api "ProjectAnalysis/interface/visualization"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

const (
	KafkaConfigPath string = "/Users/sunist/Projects/GitHub/ProjectAnalysis/Backend/config/kafka.json"
	MysqlConfigPath string = "/Users/sunist/Projects/GitHub/ProjectAnalysis/Backend/config/mysql.json"
	RedisConfigPath string = "/Users/sunist/Projects/GitHub/ProjectAnalysis/Backend/config/redis.json"
	GinConfigPath   string = "/Users/sunist/Projects/GitHub/ProjectAnalysis/Backend/config/gin.json"
	SpiderPath      string = "/Users/sunist/Projects/GitHub/ProjectAnalysis/DataSpider/main.py"
	DataTempPath    string = "/Users/sunist/Projects/GitHub/ProjectAnalysis/Backend/release/data.csv"
	DataResultPath  string = "/Users/sunist/Projects/GitHub/ProjectAnalysis/Backend/release/data_result.csv"

	GenerateConfig bool = false
)

func main() {
	// define config structs
	kafkaConfig := kafka.ClusterConfig{}
	mysqlConfig := mysql.Config{}
	redisConfig := redis.Config{}
	ginConfig := api.Config{}

	// generate config file
	if GenerateConfig {
		bytes_kafka, _ := json.Marshal(kafkaConfig)
		bytes_mysql, _ := json.Marshal(mysqlConfig)
		bytes_redis, _ := json.Marshal(redisConfig)
		bytes_gin, _ := json.Marshal(ginConfig)
		fmt.Println(string(bytes_kafka))
		fmt.Println(string(bytes_mysql))
		fmt.Println(string(bytes_redis))
		fmt.Println(string(bytes_gin))
	}

	// load config files
	cfgLoader := config.Configuration{}
	cfgLoader.Load(KafkaConfigPath, &kafkaConfig)
	cfgLoader.Load(MysqlConfigPath, &mysqlConfig)
	cfgLoader.Load(RedisConfigPath, &redisConfig)
	cfgLoader.Load(GinConfigPath, &ginConfig)
	cron.SpiderPath = SpiderPath
	cron.DataResultPath = DataResultPath
	cron.DataTempPath = DataTempPath
	log.Println("loaded config")

	// connect to requirements
	kafkaProducer := kafka.ProducerClient{}
	kafkaConsumer := kafka.ConsumerClient{}
	mysqlClient := mysql.Client{}
	redisClient := redis.Client{}
	kafkaProducer.Connect(kafkaConfig)
	kafkaConsumer.Connect(kafkaConfig)
	redisClient.Connect(redisConfig)
	mysqlClient.Connect(mysqlConfig)
	record := dto.Record{}
	mysqlClient.Sync(record)
	log.Println("connected to requirements")

	// init instances
	engine := gin.Default()
	var application *visualization.Application
	application = visualization.NewApplicationWithOpts(visualization.ApplicationOptions{
		KafkaProducer: kafkaProducer,
		KafkaConsumer: kafkaConsumer,
		MysqlDatabase: mysqlClient,
		RedisClient:   redisClient,
	})

	// start up services
	api.StartServer(engine, application, ginConfig)
}
