package visualization

import (
	"ProjectAnalysis/application/cron"
	_map "ProjectAnalysis/domain/map"
	"ProjectAnalysis/domain/visualization"
	"fmt"
	"log"
	"time"
)

func (a Application) ListenAndServe() {
	log.Println("start listen kafka message on topic: ", visualization.TopicResponse)
	go a.kafkaConsumer.Listen(visualization.TopicResponse, a.HandleResponse)
	select {}
}

func (a Application) HandleResponse(message string) {
	status, err := visualization.InsertRecord(message, a.mysqlDatabase, a.redisClient, a.callBack)
	if status {
		log.Println(fmt.Sprintf("Insert a record successfully: %v", message))
	} else if err != nil {
		log.Println(fmt.Sprintf("Serve a message errored with: %v", err))
	} else {
		log.Println(fmt.Sprintf("Receive a message"))
	}
}

func (a Application) TickAndServe() {
	count := 0
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
	default:
		time.Sleep(time.Second * 10)
		count += 1
		if count >= 4320 {
			for country, _ := range _map.Info.Counties {
				request := visualization.HiveProcessRequest{}
				request.FormatDate(time.Now())
				request.FormatUuid(fmt.Sprintf("%v:%v", country, ""))
				request.FormatType(visualization.RequestCountry.ToString())
				request.FormatData(visualization.RequestCalculateData(1))
				visualization.SendRequest(fmt.Sprintf("%v:%v", country, ""), request, a.kafkaProducer, a.redisClient)
				request.FormatData(visualization.RequestCalculateData(7))
				visualization.SendRequest(fmt.Sprintf("%v:%v", country, ""), request, a.kafkaProducer, a.redisClient)
				request.FormatData(visualization.RequestCalculateData(28))
				visualization.SendRequest(fmt.Sprintf("%v:%v", country, ""), request, a.kafkaProducer, a.redisClient)
				request.FormatData(visualization.RequestCalculateData(-1))
				visualization.SendRequest(fmt.Sprintf("%v:%v", country, ""), request, a.kafkaProducer, a.redisClient)
			}
		}
	}
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
