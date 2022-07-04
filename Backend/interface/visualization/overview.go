package visualization

import "github.com/gin-gonic/gin"

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

}
