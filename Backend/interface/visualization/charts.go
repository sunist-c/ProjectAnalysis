package visualization

import (
	"github.com/gin-gonic/gin"
	"strings"
)

// ChartsPoint the structure of a poing in ChartsData
type ChartsPoint struct {
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
	Value     int    `json:"value"`
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

	locationArr := strings.Split(location, ":")

	// TODO: implement
	// Query Country
	if len(locationArr) == 1 {
		_, err := application.QueryCountryData(location, date)
		if err != nil {
			ctx.JSON(500, BaseResponse{
				ErrorCode: 5000,
				Message:   err.Error(),
			})
		}

		//history := make([]ChartsPoint, len(result))
		//for i, value := range result {
		//
		//}

		ctx.JSON(200, ChartsDataResponse{
			BaseResponse: BaseResponse{
				ErrorCode: 0,
				Message:   "",
			},
			Data: ChartsData{
				LocationName: location,
				LocationType: TypeCountry.toString(),
				HistoryData:  []ChartsPoint{},
			},
		})
	}
}
