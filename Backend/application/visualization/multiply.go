package visualization

import (
	_map "ProjectAnalysis/domain/map"
	"ProjectAnalysis/domain/visualization"
	"ProjectAnalysis/infrastructure/common"
	"errors"
	"fmt"
	"time"
)

func (a Application) QueryCurrentProvinceData(country, province, date string, sep, count int) (records []visualization.Record, err error) {
	_, ok := _map.Info.Counties[country]
	if !ok {
		return nil, errors.New("no such country")
	}

	records = make([]visualization.Record, count)
	currentDate, _ := time.Parse("2006-01-02MST", date)
	for i := 0; i < count; i++ {
		unitRecord := visualization.Record{}
		for j := 0; j < sep; j++ {
			record := visualization.Record{}
			d := currentDate.Format("2006-01-02MST")
			record.Uuid = common.GenerateMd5Len16(fmt.Sprintf("%v:%v", country, province), d)
			ok, err := a.mysqlDatabase.Retrieve(&record)
			if err != nil {
				return nil, err
			} else if !ok {
				return nil, errors.New("no such records")
			}

			if j == 0 {
				unitRecord.LocationType = 3
				unitRecord.Uuid = record.Uuid
				unitRecord.LocationName = record.LocationName
				unitRecord.RefreshDate = currentDate
			}

			unitRecord.DailyConfirmCase += record.DailyConfirmCase
			unitRecord.DailyDeathCase += record.DailyDeathCase
			unitRecord.DailyRecoveredCase += record.DailyRecoveredCase
			unitRecord.MonthlyConfirmCase += record.MonthlyConfirmCase
			unitRecord.MonthlyDeathCase += record.MonthlyDeathCase
			unitRecord.MonthlyRecoveredCase += record.MonthlyRecoveredCase
			unitRecord.TotalConfirmCase += record.TotalConfirmCase
			unitRecord.TotalDeathCase += record.TotalDeathCase
			unitRecord.TotalRecoveredCase += record.TotalRecoveredCase

			currentDate = currentDate.AddDate(0, 0, 1)
		}
		records[i] = unitRecord
	}

	return records, nil
}

func (a Application) QueryCurrentCountryData(country, date string, sep, count int) (records []visualization.Record, err error) {
	provinces, ok := _map.Info.Counties[country]
	if !ok {
		return nil, errors.New("no such country")
	}

	records = make([]visualization.Record, count)
	currentDate, _ := time.Parse("2006-01-02MST", date)
	for i := 0; i < count; i++ {
		unitRecord := visualization.Record{}
		for j := 0; j < sep; j++ {
			record := visualization.Record{}
			d := currentDate.Format("2006-01-02MST")
			record.Uuid = common.GenerateMd5Len16(fmt.Sprintf("%v:%v", country, ""), d)
			ok, err := a.mysqlDatabase.Retrieve(&record)
			if ok && err != nil {
				if j == 0 {
					unitRecord.LocationType = 2
					unitRecord.Uuid = record.Uuid
					unitRecord.LocationName = record.LocationName
					unitRecord.RefreshDate = currentDate
				}

				unitRecord.DailyConfirmCase += record.DailyConfirmCase
				unitRecord.DailyDeathCase += record.DailyDeathCase
				unitRecord.DailyRecoveredCase += record.DailyRecoveredCase
				unitRecord.MonthlyConfirmCase += record.MonthlyConfirmCase
				unitRecord.MonthlyDeathCase += record.MonthlyDeathCase
				unitRecord.MonthlyRecoveredCase += record.MonthlyRecoveredCase
				unitRecord.TotalConfirmCase += record.TotalConfirmCase
				unitRecord.TotalDeathCase += record.TotalDeathCase
				unitRecord.TotalRecoveredCase += record.TotalRecoveredCase
			} else {
				if j == 0 {
					unitRecord.Uuid = common.GenerateMd5Len16(fmt.Sprintf("%v:%v", country, ""), date)
					unitRecord.LocationType = 2
					unitRecord.LocationName = country
					unitRecord.RefreshDate = currentDate
				}

				tempRecord := visualization.Record{}
				for _, province := range provinces.Provinces {
					pRecord := visualization.Record{}
					pRecord.Uuid = common.GenerateMd5Len16(fmt.Sprintf("%v:%v", country, province), d)
					ok, err := a.mysqlDatabase.Retrieve(&pRecord)
					if err != nil {
						return nil, err
					} else if !ok {
						return nil, errors.New("no such records")
					}

					tempRecord.DailyConfirmCase += pRecord.DailyConfirmCase
					tempRecord.DailyDeathCase += pRecord.DailyDeathCase
					tempRecord.DailyRecoveredCase += pRecord.DailyRecoveredCase
					tempRecord.MonthlyConfirmCase += pRecord.MonthlyConfirmCase
					tempRecord.MonthlyDeathCase += pRecord.MonthlyDeathCase
					tempRecord.MonthlyRecoveredCase += pRecord.MonthlyRecoveredCase
					tempRecord.TotalConfirmCase += pRecord.TotalConfirmCase
					tempRecord.TotalDeathCase += pRecord.TotalDeathCase
					tempRecord.TotalRecoveredCase += pRecord.TotalRecoveredCase
				}

				unitRecord.DailyConfirmCase += tempRecord.DailyConfirmCase
				unitRecord.DailyDeathCase += tempRecord.DailyDeathCase
				unitRecord.DailyRecoveredCase += tempRecord.DailyRecoveredCase
				unitRecord.MonthlyConfirmCase += tempRecord.MonthlyConfirmCase
				unitRecord.MonthlyDeathCase += tempRecord.MonthlyDeathCase
				unitRecord.MonthlyRecoveredCase += tempRecord.MonthlyRecoveredCase
				unitRecord.TotalConfirmCase += tempRecord.TotalConfirmCase
				unitRecord.TotalDeathCase += tempRecord.TotalDeathCase
				unitRecord.TotalRecoveredCase += tempRecord.TotalRecoveredCase
			}

			currentDate = currentDate.AddDate(0, 0, 1)
		}
		records[i] = unitRecord
	}

	return records, nil
}
