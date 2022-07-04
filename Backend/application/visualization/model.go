package visualization

import (
	"ProjectAnalysis/infrastructure/kafka"
	"ProjectAnalysis/infrastructure/mysql"
	"ProjectAnalysis/infrastructure/redis"
)

// Application the visualization application
type Application struct {
	kafkaProducer kafka.ProducerClient
	kafkaConsumer kafka.ConsumerClient
	mysqlDatabase mysql.Client
	redisClient   redis.Client
}

// ApplicationOptions the options of visualization application
type ApplicationOptions struct {
	KafkaProducer kafka.ProducerClient
	KafkaConsumer kafka.ConsumerClient
	MysqlDatabase mysql.Client
	RedisClient   redis.Client
}

// NewApplicationWithOpts use application options to create an application instance
func NewApplicationWithOpts(opt ApplicationOptions) *Application {
	return &Application{
		kafkaProducer: opt.KafkaProducer,
		kafkaConsumer: opt.KafkaConsumer,
		mysqlDatabase: opt.MysqlDatabase,
		redisClient:   opt.RedisClient,
	}
}
