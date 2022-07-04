package visualization

import "github.com/gin-gonic/gin"

// OrderedPoint the structure of data in ordered-list
type OrderedPoint struct {
	Index        int    `json:"index"`
	LocationName string `json:"location_name"`
	LocationType string `json:"location_type"`
	Value        int    `json:"value"`
}

// OrderedData the structure of ordered data
type OrderedData struct {
	LocationName string         `json:"location_name"`
	LocationType string         `json:"location_type"`
	OrderedList  []OrderedPoint `json:"ordered_list"`
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

}
