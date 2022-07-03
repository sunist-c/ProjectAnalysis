package visualization

import (
	"ProjectAnalysis/application/cron"
	_map "ProjectAnalysis/domain/map"
	"ProjectAnalysis/domain/visualization"
	"ProjectAnalysis/infrastructure/common"
	"ProjectAnalysis/infrastructure/kafka"
	"ProjectAnalysis/infrastructure/mysql"
	"ProjectAnalysis/infrastructure/redis"
	"errors"
	"fmt"
	"log"
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

func (a Application) ListenAndServe() {
	a.kafkaConsumer.Listen(visualization.TopicResponse)
	select {
	case message := <-a.kafkaConsumer.MessageChan:
		status, err := visualization.InsertRecord(message, a.mysqlDatabase, a.redisClient, a.callBack)
		if status {
			log.Println(fmt.Sprintf("Insert a record successfully: %v", message))
		}

		if err != nil {
			log.Println(fmt.Sprintf("Serve a message errored with: %v", err))
		}
	}
}

func (a Application) TickAndServe() {
	select {
	case message := <-cron.ReadChan:
		request := visualization.HiveProcessRequest{}
		request.Date = message.Date
		request.FormatUuid(fmt.Sprintf("%v:%v", message.CountryName, message.ProvinceName))
		_map.Info.AddRecord(message.CountryName, message.ProvinceName)

		request.FormatType(visualization.RequestImport.ToString())

		data := visualization.RequestImportData{
			Confirm:          int(message.Confirm),
			Death:            int(message.Death),
			Recovered:        int(message.Recovered),
			RefreshTime:      message.Date,
			LocationCountry:  message.CountryName,
			LocationProvince: message.ProvinceName,
		}
		request.FormatData(data)

		visualization.SendRequest(fmt.Sprintf("%v:%v", message.CountryName, message.ProvinceName), request, a.kafkaProducer, a.redisClient)
	}
}

func (a Application) QueryProvinceData(country, province, date string) (record visualization.Record, err error) {
	_, ok := _map.Info.Counties[country]
	if !ok {
		return visualization.Record{}, errors.New("no such country")
	}

	record = visualization.Record{}
	record.Uuid = common.GenerateMd5Len16(fmt.Sprintf("%v:%v", country, province), date)
	if ok, err := a.mysqlDatabase.Retrieve(&record); err != nil {
		return visualization.Record{}, err
	} else {
		if ok {
			return record, nil
		} else {
			return visualization.Record{}, errors.New("query failed")
		}
	}
}

func (a Application) QueryCountryData(country, date string) (records []visualization.Record, err error) {
	provinces, ok := _map.Info.Counties[country]
	if !ok {
		return nil, errors.New("no such country")
	}

	records = make([]visualization.Record, len(provinces.Provinces)+1)
	records[0].Uuid = common.GenerateMd5Len16(fmt.Sprintf("%v:%v", country, ""), date)
	records[0].LocationType = 2

	for i, province := range provinces.Provinces {
		records[i+1], err = a.QueryProvinceData(country, province, date)
		records[0].DailyConfirmCase += records[i+1].DailyConfirmCase
		records[0].DailyDeathCase += records[i+1].DailyDeathCase
		records[0].DailyRecoveredCase += records[i+1].DailyRecoveredCase
		records[0].MonthlyConfirmCase += records[i+1].MonthlyConfirmCase
		records[0].MonthlyDeathCase += records[i+1].MonthlyDeathCase
		records[0].MonthlyRecoveredCase += records[i+1].MonthlyRecoveredCase
		records[0].TotalConfirmCase += records[i+1].TotalConfirmCase
		records[0].TotalDeathCase += records[i+1].TotalDeathCase
		records[0].TotalRecoveredCase += records[i+1].TotalRecoveredCase
		if err != nil {
			return nil, err
		}
	}

	return records, nil
}

func (a Application) callBack(uuid string) {
	_, _, origin := visualization.SplitUuid(uuid)
	result, _ := a.redisClient.Read(origin)
	cache := visualization.RequestCache{}
	cache.Format(result)

	// send today request
	request := visualization.HiveProcessRequest{}
	request.FormatType(visualization.RequestProvince.ToString())
	request.Date = cache.Date
	request.FormatLength(1)
	request.FormatData(visualization.RequestCalculateData(1))
	request.FormatUuid(cache.Location)
	visualization.SendRequest(cache.Location, request, a.kafkaProducer, a.redisClient)

	// send 7-day request
	request = visualization.HiveProcessRequest{}
	request.FormatType(visualization.RequestProvince.ToString())
	request.Date = cache.Date
	request.FormatLength(7)
	request.FormatData(visualization.RequestCalculateData(7))
	request.FormatUuid(cache.Location)
	visualization.SendRequest(cache.Location, request, a.kafkaProducer, a.redisClient)

	// send 28-day request
	request = visualization.HiveProcessRequest{}
	request.FormatType(visualization.RequestProvince.ToString())
	request.Date = cache.Date
	request.FormatLength(28)
	request.FormatData(visualization.RequestCalculateData(28))
	request.FormatUuid(cache.Location)
	visualization.SendRequest(cache.Location, request, a.kafkaProducer, a.redisClient)

	// send all-day request
	request = visualization.HiveProcessRequest{}
	request.FormatType(visualization.RequestProvince.ToString())
	request.Date = cache.Date
	request.FormatLength(-1)
	request.FormatData(visualization.RequestCalculateData(-1))
	request.FormatUuid(cache.Location)
	visualization.SendRequest(cache.Location, request, a.kafkaProducer, a.redisClient)
}
