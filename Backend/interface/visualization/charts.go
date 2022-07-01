package visualization

import "github.com/gin-gonic/gin"

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

}
