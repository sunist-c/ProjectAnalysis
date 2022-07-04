package visualization

import (
	"ProjectAnalysis/application/cron"
	"ProjectAnalysis/application/visualization"
	"github.com/gin-gonic/gin"
)

func StartServer(engine *gin.Engine, visualizer *visualization.Application, cfg Config) {
	gin.SetMode(cfg.Mode)
	application = visualizer
	Bind(engine)
	engine.Run(cfg.Address)
}

func Bind(engine *gin.Engine) {
	group := engine.Group("/")
	group.Use(QueryParamMiddleware())
	group.GET("/map-data/:location/:date", MapDataHandler)
	group.GET("/charts-data/:location/:date", ChartsDataHandler)
	group.GET("/ordered-data/:location/:date", OrderedDataHandler)
	group.GET("/overview-data/:location/:date", OverviewDataHandler)
	engine.GET("map-info", MapInfoHandler)
	engine.GET("update", func(context *gin.Context) {
		cron.LoadCsv()
		context.JSON(200, BaseResponse{
			ErrorCode: 0,
			Message:   "",
		})
	})
}
