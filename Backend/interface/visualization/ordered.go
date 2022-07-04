package visualization

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
)

// OrderedPoint the structure of data in ordered-list
type OrderedPoint struct {
	LocationName         string `json:"location_name"`
	TotalConfirmCase     int    `json:"total_confirm_case"`
	MonthlyConfirmCase   int    `json:"monthly_confirm_case"`
	WeeklyConfirmCase    int    `json:"weekly_confirm_case"`
	DailyConfirmCase     int    `json:"daily_confirm_case"`
	TotalDeathCase       int    `json:"total_death_case"`
	MonthlyDeathCase     int    `json:"monthly_death_case"`
	WeeklyDeathCase      int    `json:"weekly_death_case"`
	DailyDeathCase       int    `json:"daily_death_case"`
	TotalRecoveredCase   int    `json:"total_recovered_case"`
	MonthlyRecoveredCase int    `json:"monthly_recovered_case"`
	WeeklyRecoveredCase  int    `json:"weekly_recovered_case"`
	DailyRecoveredCase   int    `json:"daily_recovered_case"`
}

// OrderedData the structure of ordered data
type OrderedData struct {
	LocationName string                    `json:"location_name"`
	LocationType string                    `json:"location_type"`
	OrderedList  map[string][]OrderedPoint `json:"ordered_list"`
}

// OrderedDataRequest the structure of map data interface request
type OrderedDataRequest struct {
}

// OrderedDataResponse the structure of map data interface response
type OrderedDataResponse struct {
	BaseResponse
	Data OrderedData `json:"data"`
}

// OrderedDataHandler the handler of map data interface
func OrderedDataHandler(ctx *gin.Context) {
	_location, _ := ctx.Get("location")
	_date, _ := ctx.Get("date")
	location, _ := _location.(string)
	date, _ := _date.(string)
	location = strings.Trim(location, ":")
	_methods := ctx.Query("methods")
	_count := ctx.Query("count")
	count, err := strconv.Atoi(_count)
	if err != nil {
		ctx.JSON(400, BaseResponse{
			ErrorCode: 4002,
			Message:   "unknown count format",
		})
		return
	}

	methods := strings.Split(strings.Trim(_methods, ","), ",")
	if len(methods) == 1 && methods[0] == "" {
		ctx.JSON(400, BaseResponse{
			ErrorCode: 4002,
			Message:   "empty method field",
		})
		return
	}

	switch location {
	case "world":
		result, err := application.QueryWorldOrderList(date, count, methods)
		if err != nil {
			ctx.JSON(500, BaseResponse{
				ErrorCode: 5000,
				Message:   err.Error(),
			})
			return
		} else {
			data := make(map[string][]OrderedPoint)
			for s, records := range result {
				data[s] = make([]OrderedPoint, len(records))
				for i := 0; i < len(records); i++ {
					data[s][i] = OrderedPoint{
						LocationName:         records[i].LocationName,
						TotalConfirmCase:     records[i].TotalConfirmCase,
						MonthlyConfirmCase:   records[i].MonthlyConfirmCase,
						WeeklyConfirmCase:    records[i].WeeklyConfirmCase,
						DailyConfirmCase:     records[i].DailyConfirmCase,
						TotalDeathCase:       records[i].TotalDeathCase,
						MonthlyDeathCase:     records[i].MonthlyDeathCase,
						WeeklyDeathCase:      records[i].WeeklyDeathCase,
						DailyDeathCase:       records[i].DailyDeathCase,
						TotalRecoveredCase:   records[i].TotalRecoveredCase,
						MonthlyRecoveredCase: records[i].MonthlyRecoveredCase,
						WeeklyRecoveredCase:  records[i].WeeklyRecoveredCase,
						DailyRecoveredCase:   records[i].DailyRecoveredCase,
					}
				}
			}

			ctx.JSON(200, OrderedDataResponse{
				BaseResponse: BaseResponse{
					ErrorCode: 0,
					Message:   "",
				},
				Data: OrderedData{
					LocationName: location,
					LocationType: TypeCountry.toString(),
					OrderedList:  data,
				},
			})
			return
		}
	default:
		result, err := application.QueryCountryOrderedList(location, date, count, methods)
		if err != nil {
			ctx.JSON(500, BaseResponse{
				ErrorCode: 5000,
				Message:   err.Error(),
			})
			return
		} else {
			data := make(map[string][]OrderedPoint)
			for s, records := range result {
				data[s] = make([]OrderedPoint, len(records))
				for i := 0; i < len(records); i++ {
					data[s][i] = OrderedPoint{
						LocationName:         records[i].LocationName,
						TotalConfirmCase:     records[i].TotalConfirmCase,
						MonthlyConfirmCase:   records[i].MonthlyConfirmCase,
						WeeklyConfirmCase:    records[i].WeeklyConfirmCase,
						DailyConfirmCase:     records[i].DailyConfirmCase,
						TotalDeathCase:       records[i].TotalDeathCase,
						MonthlyDeathCase:     records[i].MonthlyDeathCase,
						WeeklyDeathCase:      records[i].WeeklyDeathCase,
						DailyDeathCase:       records[i].DailyDeathCase,
						TotalRecoveredCase:   records[i].TotalRecoveredCase,
						MonthlyRecoveredCase: records[i].MonthlyRecoveredCase,
						WeeklyRecoveredCase:  records[i].WeeklyRecoveredCase,
						DailyRecoveredCase:   records[i].DailyRecoveredCase,
					}
				}
			}

			ctx.JSON(200, OrderedDataResponse{
				BaseResponse: BaseResponse{
					ErrorCode: 0,
					Message:   "",
				},
				Data: OrderedData{
					LocationName: location,
					LocationType: TypeProvince.toString(),
					OrderedList:  data,
				},
			})
			return
		}
	}
}
