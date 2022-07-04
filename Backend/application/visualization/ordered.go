package visualization

import (
	_map "ProjectAnalysis/domain/map"
	"ProjectAnalysis/domain/visualization"
	"errors"
	"sort"
)

func (a Application) QueryWorldOrderList(date string, count int, methods []string) (records map[string][]visualization.Record, err error) {
	records = make(map[string][]visualization.Record, count)

	totalRecords := make([]visualization.Record, 0, len(_map.Info.Counties))
	if len(_map.Info.Counties)-count < 0 {
		return nil, errors.New("no enough countries")
	}

	for country, _ := range _map.Info.Counties {
		record, err := a.QueryCountryData(country, date)
		if err != nil {
			return nil, err
		}
		totalRecords = append(totalRecords, record[0])
	}

	startIndex := len(_map.Info.Counties) - count
	for _, method := range methods {
		switch method {
		case "daily_confirm_list":
			sort.Sort(byDailyConfirm(totalRecords))
			records[method] = totalRecords[startIndex:]
		case "weekly_confirm_list":
			sort.Sort(byWeeklyConfirm(totalRecords))
			records[method] = totalRecords[startIndex:]
		case "monthly_confirm_list":
			sort.Sort(byMonthlyConfirm(totalRecords))
			records[method] = totalRecords[startIndex:]
		case "total_confirm_list":
			sort.Sort(byTotalConfirm(totalRecords))
			records[method] = totalRecords[startIndex:]
		case "daily_death_list":
			sort.Sort(byDailyDeath(totalRecords))
			records[method] = totalRecords[startIndex:]
		case "weekly_death_list":
			sort.Sort(byWeeklyDeath(totalRecords))
			records[method] = totalRecords[startIndex:]
		case "monthly_death_list":
			sort.Sort(byMonthlyDeath(totalRecords))
			records[method] = totalRecords[startIndex:]
		case "total_death_list":
			sort.Sort(byTotalDeath(totalRecords))
			records[method] = totalRecords[startIndex:]
		case "daily_recovered_list":
			sort.Sort(byDailyRecovered(totalRecords))
			records[method] = totalRecords[startIndex:]
		case "weekly_recovered_list":
			sort.Sort(byWeeklyRecovered(totalRecords))
			records[method] = totalRecords[startIndex:]
		case "monthly_recovered_list":
			sort.Sort(byMonthlyRecovered(totalRecords))
			records[method] = totalRecords[startIndex:]
		case "total_recovered_list":
			sort.Sort(byTotalRecovered(totalRecords))
			records[method] = totalRecords[startIndex:]
		default:
			return nil, errors.New("unknown sort method")
		}
	}

	return records, nil
}

func (a Application) QueryCountryOrderedList(country, date string, count int, methods []string) (records map[string][]visualization.Record, err error) {
	records = make(map[string][]visualization.Record, count)

	totalRecords := make([]visualization.Record, 0, len(_map.Info.Counties[country].Provinces))
	if len(_map.Info.Counties[country].Provinces)-count < 0 {
		return nil, errors.New("no enough countries")
	}

	record, err := a.QueryCountryData(country, date)
	if err != nil {
		return nil, err
	}

	totalRecords = record[1:]

	startIndex := len(_map.Info.Counties) - count + 1
	for _, method := range methods {
		switch method {
		case "daily_confirm_list":
			sort.Sort(byDailyConfirm(totalRecords))
			records[method] = totalRecords[startIndex:]
		case "weekly_confirm_list":
			sort.Sort(byWeeklyConfirm(totalRecords))
			records[method] = totalRecords[startIndex:]
		case "monthly_confirm_list":
			sort.Sort(byMonthlyConfirm(totalRecords))
			records[method] = totalRecords[startIndex:]
		case "total_confirm_list":
			sort.Sort(byTotalConfirm(totalRecords))
			records[method] = totalRecords[startIndex:]
		case "daily_death_list":
			sort.Sort(byDailyDeath(totalRecords))
			records[method] = totalRecords[startIndex:]
		case "weekly_death_list":
			sort.Sort(byWeeklyDeath(totalRecords))
			records[method] = totalRecords[startIndex:]
		case "monthly_death_list":
			sort.Sort(byMonthlyDeath(totalRecords))
			records[method] = totalRecords[startIndex:]
		case "total_death_list":
			sort.Sort(byTotalDeath(totalRecords))
			records[method] = totalRecords[startIndex:]
		case "daily_recovered_list":
			sort.Sort(byDailyRecovered(totalRecords))
			records[method] = totalRecords[startIndex:]
		case "weekly_recovered_list":
			sort.Sort(byWeeklyRecovered(totalRecords))
			records[method] = totalRecords[startIndex:]
		case "monthly_recovered_list":
			sort.Sort(byMonthlyRecovered(totalRecords))
			records[method] = totalRecords[startIndex:]
		case "total_recovered_list":
			sort.Sort(byTotalRecovered(totalRecords))
			records[method] = totalRecords[startIndex:]
		default:
			return nil, errors.New("unknown sort method")
		}
	}

	return records, nil
}
