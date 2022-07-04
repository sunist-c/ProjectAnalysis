package visualization

import (
	"ProjectAnalysis/infrastructure/kafka"
	"ProjectAnalysis/infrastructure/mysql"
	"ProjectAnalysis/infrastructure/redis"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

func SplitUuid(uuid string) (requestType, days int, origin string) {
	arr := strings.Split(uuid, "-")
	days, _ = strconv.Atoi(arr[2])
	requestType, _ = strconv.Atoi(arr[1])
	origin = arr[0]
	return
}

// InsertRecord insert a record into mysql database from kafka message string
func InsertRecord(message string, client mysql.Client, redis redis.Client, callBack func(uuid string)) (ok bool, err error) {
	baseResp := BaseHiveProcessResponse{}
	err = json.Unmarshal([]byte(message), &baseResp)
	if err != nil {
		return false, err
	}

	requestType, days, origin := SplitUuid(baseResp.Uuid)
	result, err := redis.Read(origin)
	if err != nil {
		return false, err
	} else {
		cache := RequestCache{}
		cache.Format(result)

		if requestType == 1 {
			cache.Imported = true

			err = redis.Write(origin, cache.toString(), 1200)
			if err != nil {
				return false, err
			} else {
				callBack(baseResp.Uuid)
				return false, nil
			}
		} else if requestType == 2 {
			resp := CalculateCountryHiveProcessResponse{}
			json.Unmarshal([]byte(message), &resp)
			switch days {
			case 1:
				cache.Days1 = true
				cache.Record.DailyConfirmCase = resp.Data.Confirm
				cache.Record.DailyDeathCase = resp.Data.Death
				cache.Record.DailyRecoveredCase = resp.Data.Recovered
			case 7:
				cache.Days7 = true
				cache.Record.WeeklyConfirmCase = resp.Data.Confirm
				cache.Record.WeeklyDeathCase = resp.Data.Death
				cache.Record.WeeklyRecoveredCase = resp.Data.Recovered
			case 28:
				cache.Days28 = true
				cache.Record.MonthlyConfirmCase = resp.Data.Confirm
				cache.Record.MonthlyDeathCase = resp.Data.Death
				cache.Record.MonthlyRecoveredCase = resp.Data.Recovered
			case -1:
				cache.DaysTotal = true
				cache.Record.TotalConfirmCase = resp.Data.Confirm
				cache.Record.TotalDeathCase = resp.Data.Death
				cache.Record.TotalRecoveredCase = resp.Data.Recovered
			}
			if cache.Ready() {
				record := cache.Record
				_, err = client.Create(record)
				if err != nil {
					return false, err
				}
				err = redis.Remove(origin)
				log.Println(fmt.Sprintf("remove redis record failed with: %v", err))

				return true, nil
			} else {
				err = redis.Write(origin, cache.toString(), 1200)
				if err != nil {
					return false, err
				}

				return false, nil
			}
		} else if requestType == 3 {
			resp := CalculateProvinceHiveProcessResponse{}
			json.Unmarshal([]byte(message), &resp)
			switch days {
			case 1:
				cache.Days1 = true
				cache.Record.DailyConfirmCase = resp.Data.Confirm
				cache.Record.DailyDeathCase = resp.Data.Death
				cache.Record.DailyRecoveredCase = resp.Data.Recovered
			case 7:
				cache.Days7 = true
				cache.Record.WeeklyConfirmCase = resp.Data.Confirm
				cache.Record.WeeklyDeathCase = resp.Data.Death
				cache.Record.WeeklyRecoveredCase = resp.Data.Recovered
			case 28:
				cache.Days28 = true
				cache.Record.MonthlyConfirmCase = resp.Data.Confirm
				cache.Record.MonthlyDeathCase = resp.Data.Death
				cache.Record.MonthlyRecoveredCase = resp.Data.Recovered
			case -1:
				cache.DaysTotal = true
				cache.Record.TotalConfirmCase = resp.Data.Confirm
				cache.Record.TotalDeathCase = resp.Data.Death
				cache.Record.TotalRecoveredCase = resp.Data.Recovered
			}
			if cache.Ready() {
				record := cache.Record
				_, err = client.Create(record)
				if err != nil {
					return false, err
				}
				err = redis.Remove(origin)
				log.Println(fmt.Sprintf("remove redis record failed with: %v", err))

				return true, nil
			} else {
				err = redis.Write(origin, cache.toString(), 1200)
				if err != nil {
					return false, err
				}

				return false, nil
			}
		} else {
			return false, errors.New("unknown request type")
		}
	}
}

// SendRequest send a message to hive via kafka
func SendRequest(location string, request HiveProcessRequest, client kafka.ProducerClient, redis redis.Client) (uuid string, err error) {
	requestType, _, origin := SplitUuid(request.Uuid)
	date, _ := time.Parse("2006-01-02", request.Date)
	cache := RequestCache{
		Uuid:      origin,
		Date:      request.Date,
		Location:  location,
		Days1:     false,
		Days7:     false,
		Days28:    false,
		DaysTotal: false,
		Imported:  false,
		Record: Record{
			Uuid:                 origin,
			RefreshDate:          date,
			LocationName:         location,
			LocationType:         requestType,
			DailyConfirmCase:     0,
			DailyDeathCase:       0,
			DailyRecoveredCase:   0,
			WeeklyConfirmCase:    0,
			WeeklyDeathCase:      0,
			WeeklyRecoveredCase:  0,
			MonthlyConfirmCase:   0,
			MonthlyDeathCase:     0,
			MonthlyRecoveredCase: 0,
			TotalConfirmCase:     0,
			TotalDeathCase:       0,
			TotalRecoveredCase:   0,
		},
	}

	err = redis.Write(origin, cache.toString(), 1200)
	if err != nil {
		return "", err
	}

	_, err = client.Send(TopicRequest, request.ToString())
	if err != nil {
		return "", err
	}

	return origin, nil
}
