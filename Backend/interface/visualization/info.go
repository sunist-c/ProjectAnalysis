package visualization

import (
	"github.com/gin-gonic/gin"
)

type MapInfoRequest struct {
}

type MapInfoResponse struct {
	BaseResponse
	Data map[string][]string `json:"data"`
}

func MapInfoHandler(ctx *gin.Context) {
	location := ctx.Query("location")
	if location == "world" {
		data := application.QueryWholeMap()
		ctx.JSON(200, MapInfoResponse{
			BaseResponse: BaseResponse{
				ErrorCode: 0,
				Message:   "",
			},
			Data: data,
		})
		return
	} else {
		result := make(map[string][]string, 1)
		result[location] = application.QueryCountryMap(location)
		ctx.JSON(200, MapInfoResponse{
			BaseResponse: BaseResponse{
				ErrorCode: 0,
				Message:   "",
			},
			Data: result,
		})
		return
	}
}
