package visualization

import (
	"github.com/gin-gonic/gin"
	"strings"
)

// MapData the structure of map data
type MapData struct {
	Longitude            float64 `json:"longitude"`
	Latitude             float64 `json:"latitude"`
	LocationName         string  `json:"location_name"`
	LocationType         string  `json:"location_type"`
	DailyConfirmCase     int     `json:"daily_confirm_case"`
	DailyDeathCase       int     `json:"daily_death_case"`
	DailyRecoveredCase   int     `json:"daily_recovered_case"`
	WeeklyConfirmCase    int     `json:"weekly_confirm_case"`
	WeeklyDeathCase      int     `json:"weekly_death_case"`
	WeeklyRecoveredCase  int     `json:"weekly_recovered_case"`
	MonthlyConfirmCase   int     `json:"monthly_confirm_case"`
	MonthlyDeathCase     int     `json:"monthly_death_case"`
	MonthlyRecoveredCase int     `json:"monthly_recovered_case"`
	TotalConfirmCase     int     `json:"total_confirm_case"`
	TotalDeathCase       int     `json:"total_death_case"`
	TotalRecoveredCase   int     `json:"total_recovered_case"`
}

// MapDataRequest the structure of map data interface request
type MapDataRequest struct {
}

// MapDataResponse the structure of map data interface response
type MapDataResponse struct {
	BaseResponse
	Data []MapData `json:"data"`
}

// MapDataHandler the handler of map data interface
func MapDataHandler(ctx *gin.Context) {
	_location, _ := ctx.Get("location")
	_date, _ := ctx.Get("date")
	location, _ := _location.(string)
	date, _ := _date.(string)

	locationArr := strings.Split(location, ":")

	// query country
	if len(locationArr) == 1 {
		result, err := application.QueryCountryData(location, date)
		if err != nil {
			ctx.JSON(500, BaseResponse{
				ErrorCode: 5000,
				Message:   err.Error(),
			})
		}

		data := make([]MapData, len(result))
		for i, record := range result {
			data[i] = MapData{
				Longitude:            0,
				Latitude:             0,
				LocationName:         record.LocationName,
				LocationType:         TypeProvince.toString(),
				DailyConfirmCase:     record.DailyConfirmCase,
				DailyDeathCase:       record.DailyDeathCase,
				DailyRecoveredCase:   record.DailyRecoveredCase,
				WeeklyConfirmCase:    record.WeeklyConfirmCase,
				WeeklyDeathCase:      record.WeeklyDeathCase,
				WeeklyRecoveredCase:  record.WeeklyRecoveredCase,
				MonthlyConfirmCase:   record.MonthlyConfirmCase,
				MonthlyDeathCase:     record.MonthlyDeathCase,
				MonthlyRecoveredCase: record.MonthlyRecoveredCase,
				TotalConfirmCase:     record.TotalConfirmCase,
				TotalDeathCase:       record.TotalDeathCase,
				TotalRecoveredCase:   record.TotalRecoveredCase,
			}
		}
		data[0].LocationType = TypeCountry.toString()

		ctx.JSON(200, MapDataResponse{
			BaseResponse: BaseResponse{
				ErrorCode: 0,
				Message:   "",
			},
			Data: data,
		})
	}
}
