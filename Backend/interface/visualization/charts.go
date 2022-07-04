package visualization

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
)

// ChartsPoint the structure of a poing in ChartsData
type ChartsPoint struct {
	StartDate            string `json:"start_date"`
	EndDate              string `json:"end_date"`
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

// ChartsData the structure of charts data
type ChartsData struct {
	LocationName string        `json:"location_name"`
	LocationType string        `json:"location_type"`
	HistoryData  []ChartsPoint `json:"history_data"`
}

// ChartsDataRequest the structure of charts data interface request
type ChartsDataRequest struct {
}

// ChartsDataResponse the structure of charts data interface response
type ChartsDataResponse struct {
	BaseResponse
	Data ChartsData `json:"data"`
}

// ChartsDataHandler the handler of charts data interface
func ChartsDataHandler(ctx *gin.Context) {
	_location, _ := ctx.Get("location")
	_date, _ := ctx.Get("date")
	location, _ := _location.(string)
	date, _ := _date.(string)
	_sep, _count := ctx.Query("separation"), ctx.Query("count")
	sep, err := strconv.Atoi(_sep)
	if err != nil {
		ctx.JSON(400, BaseResponse{
			ErrorCode: 4002,
			Message:   "bad_request: unknown separation format",
		})
		return
	}

	count, err := strconv.Atoi(_count)
	if err != nil {
		ctx.JSON(400, BaseResponse{
			ErrorCode: 4002,
			Message:   "bad_request: unknown count format",
		})
		return
	}

	locationArr := strings.Split(location, ":")

	// Query Country
	if len(locationArr) == 1 {
		records, err := application.QueryCurrentCountryData(location, date, sep, count)
		if err != nil {
			ctx.JSON(500, BaseResponse{
				ErrorCode: 5000,
				Message:   err.Error(),
			})
			return
		}

		data := make([]ChartsPoint, count)
		for i, record := range records {
			data[i] = ChartsPoint{
				StartDate:            record.RefreshDate.AddDate(0, 0, sep*-1).Format("2006-01-02MST"),
				EndDate:              record.RefreshDate.Format("2006-01-02MST"),
				TotalConfirmCase:     record.TotalConfirmCase,
				MonthlyConfirmCase:   record.MonthlyConfirmCase,
				WeeklyConfirmCase:    record.WeeklyConfirmCase,
				DailyConfirmCase:     record.DailyConfirmCase,
				TotalDeathCase:       record.TotalDeathCase,
				MonthlyDeathCase:     record.MonthlyDeathCase,
				WeeklyDeathCase:      record.WeeklyDeathCase,
				DailyDeathCase:       record.DailyDeathCase,
				TotalRecoveredCase:   record.TotalRecoveredCase,
				MonthlyRecoveredCase: record.MonthlyRecoveredCase,
				WeeklyRecoveredCase:  record.WeeklyRecoveredCase,
				DailyRecoveredCase:   record.DailyRecoveredCase,
			}
		}

		ctx.JSON(200, ChartsDataResponse{
			BaseResponse: BaseResponse{
				ErrorCode: 0,
				Message:   "",
			},
			Data: ChartsData{
				LocationName: location,
				LocationType: TypeCountry.toString(),
				HistoryData:  data,
			},
		})
		return
	} else if len(locationArr) == 2 {
		records, err := application.QueryCurrentProvinceData(locationArr[0], locationArr[1], date, sep, count)
		if err != nil {
			ctx.JSON(500, BaseResponse{
				ErrorCode: 5000,
				Message:   err.Error(),
			})
			return
		}

		data := make([]ChartsPoint, count)
		for i, record := range records {
			data[i] = ChartsPoint{
				StartDate:            record.RefreshDate.AddDate(0, 0, sep*-1).Format("2006-01-02MST"),
				EndDate:              record.RefreshDate.Format("2006-01-02MST"),
				TotalConfirmCase:     record.TotalConfirmCase,
				MonthlyConfirmCase:   record.MonthlyConfirmCase,
				WeeklyConfirmCase:    record.WeeklyConfirmCase,
				DailyConfirmCase:     record.DailyConfirmCase,
				TotalDeathCase:       record.TotalDeathCase,
				MonthlyDeathCase:     record.MonthlyDeathCase,
				WeeklyDeathCase:      record.WeeklyDeathCase,
				DailyDeathCase:       record.DailyDeathCase,
				TotalRecoveredCase:   record.TotalRecoveredCase,
				MonthlyRecoveredCase: record.MonthlyRecoveredCase,
				WeeklyRecoveredCase:  record.WeeklyRecoveredCase,
				DailyRecoveredCase:   record.DailyRecoveredCase,
			}
		}

		ctx.JSON(200, ChartsDataResponse{
			BaseResponse: BaseResponse{
				ErrorCode: 0,
				Message:   "",
			},
			Data: ChartsData{
				LocationName: location,
				LocationType: TypeProvince.toString(),
				HistoryData:  data,
			},
		})
		return
	} else {
		ctx.JSON(400, BaseResponse{
			ErrorCode: 4002,
			Message:   "bad_request: unknown location format",
		})
		return
	}
}
