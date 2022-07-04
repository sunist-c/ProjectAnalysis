package visualization

import (
	"github.com/gin-gonic/gin"
	"time"
)

// QueryParamHandler the handler which used in QueryParamMiddleware, to query the params in url
func QueryParamHandler(ctx *gin.Context) {
	location, date := ctx.Param("location"), ctx.Param("date")
	if location == "" || date == "" {
		ctx.JSON(400, BaseResponse{
			ErrorCode: 4001,
			Message:   "bad_request: location or date field is empty",
		})
		ctx.Abort()
	} else {
		_, err := time.Parse("2006-01-02", date)
		if err != nil {
			ctx.JSON(400, BaseResponse{
				ErrorCode: 4002,
				Message:   err.Error(),
			})
			ctx.Abort()
		}
		ctx.Set("location", location)
		ctx.Set("date", date)
	}
	ctx.Next()
}

// QueryParamMiddleware the middleware which query the params in url and set to context
func QueryParamMiddleware() gin.HandlerFunc {
	return QueryParamHandler
}
