package visualization

import (
	"github.com/gin-gonic/gin"
	"strings"
)

// OverviewData the structure of overview data
type OverviewData struct {
	LocationName         string `json:"location_name"`
	LocationType         string `json:"location_type"`
	DailyConfirmCase     int    `json:"daily_confirm_case"`
	DailyDeathCase       int    `json:"daily_death_case"`
	DailyRecoveredCase   int    `json:"daily_recovered_case"`
	WeeklyConfirmCase    int    `json:"weekly_confirm_case"`
	WeeklyDeathCase      int    `json:"weekly_death_case"`
	WeeklyRecoveredCase  int    `json:"weekly_recovered_case"`
	MonthlyConfirmCase   int    `json:"monthly_confirm_case"`
	MonthlyDeathCase     int    `json:"monthly_death_case"`
	MonthlyRecoveredCase int    `json:"monthly_recovered_case"`
	TotalConfirmCase     int    `json:"total_confirm_case"`
	TotalDeathCase       int    `json:"total_death_case"`
	TotalRecoveredCase   int    `json:"total_recovered_case"`
}

// OverviewDataRequest the structure of map data interface request
type OverviewDataRequest struct {
}

// OverviewDataResponse the structure of map data interface response
type OverviewDataResponse struct {
	BaseResponse
	Data OverviewData `json:"data"`
}

// OverviewDataHandler the handler of map data interface
func OverviewDataHandler(ctx *gin.Context) {
	_location, _ := ctx.Get("location")
	_date, _ := ctx.Get("date")
	location, _ := _location.(string)
	date, _ := _date.(string)
	location = strings.Trim(location, ":")
	locationArr := strings.Split(location, ":")

	if len(locationArr) == 1 {
		result, err := application.QueryCountryData(location, date)
		if err != nil {
			ctx.JSON(500, BaseResponse{
				ErrorCode: 5000,
				Message:   err.Error(),
			})
			return
		}

		if len(result) == 0 {
			ctx.JSON(500, BaseResponse{
				ErrorCode: 0,
				Message:   "empty country data",
			})
			return
		}

		ctx.JSON(200, OverviewDataResponse{
			BaseResponse: BaseResponse{
				ErrorCode: 0,
				Message:   "",
			},
			Data: OverviewData{
				LocationName:         result[0].LocationName,
				LocationType:         TypeCountry.toString(),
				DailyConfirmCase:     result[0].DailyConfirmCase,
				DailyDeathCase:       result[0].DailyDeathCase,
				DailyRecoveredCase:   result[0].DailyRecoveredCase,
				WeeklyConfirmCase:    result[0].WeeklyConfirmCase,
				WeeklyDeathCase:      result[0].WeeklyDeathCase,
				WeeklyRecoveredCase:  result[0].WeeklyRecoveredCase,
				MonthlyConfirmCase:   result[0].MonthlyConfirmCase,
				MonthlyDeathCase:     result[0].MonthlyDeathCase,
				MonthlyRecoveredCase: result[0].WeeklyRecoveredCase,
				TotalConfirmCase:     result[0].TotalConfirmCase,
				TotalDeathCase:       result[0].TotalDeathCase,
				TotalRecoveredCase:   result[0].TotalRecoveredCase,
			},
		})
		return
	}
}
