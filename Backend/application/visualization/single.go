package visualization

import (
	_map "ProjectAnalysis/domain/map"
	"ProjectAnalysis/domain/visualization"
	"ProjectAnalysis/infrastructure/common"
	"errors"
	"fmt"
)

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

	record := visualization.Record{}
	record.Uuid = common.GenerateMd5Len16(fmt.Sprintf("%v:%v", country, ""), date)
	success, err := a.mysqlDatabase.Retrieve(&record)
	if success && err == nil {
		records[0] = record
	} else {
		records[0].Uuid = common.GenerateMd5Len16(fmt.Sprintf("%v:%v", country, ""), date)
		records[0].LocationName = country
		records[0].LocationType = 2
	}

	for i, province := range provinces.Provinces {
		records[i+1], err = a.QueryProvinceData(country, province, date)
		if err != nil {
			return nil, err
		}

		if success {
			records[0].DailyConfirmCase += records[i+1].DailyConfirmCase
			records[0].DailyDeathCase += records[i+1].DailyDeathCase
			records[0].DailyRecoveredCase += records[i+1].DailyRecoveredCase
			records[0].MonthlyConfirmCase += records[i+1].MonthlyConfirmCase
			records[0].MonthlyDeathCase += records[i+1].MonthlyDeathCase
			records[0].MonthlyRecoveredCase += records[i+1].MonthlyRecoveredCase
			records[0].TotalConfirmCase += records[i+1].TotalConfirmCase
			records[0].TotalDeathCase += records[i+1].TotalDeathCase
			records[0].TotalRecoveredCase += records[i+1].TotalRecoveredCase
			a.mysqlDatabase.Create(records[0])
		}
	}

	return records, nil
}
