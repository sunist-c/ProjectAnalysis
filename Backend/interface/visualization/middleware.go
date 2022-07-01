package visualization

import "github.com/gin-gonic/gin"

// QueryParamHandler the handler which used in QueryParamMiddleware, to query the params in url
func QueryParamHandler(ctx *gin.Context) {
	location, date := ctx.Query(":location"), ctx.Query(":date")
	if location == "" || date == "" {
		ctx.Abort()
	} else {

	}
	ctx.Next()
}

// QueryParamMiddleware the middleware which query the params in url and set to context
func QueryParamMiddleware() gin.HandlerFunc {
	return QueryParamHandler
}
