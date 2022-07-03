package visualization

import (
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
	engine.Use(QueryParamMiddleware())
	engine.GET("/map-data/:location/:date", MapDataHandler)
	engine.GET("/charts-data/:location/:date", ChartsDataHandler)
	engine.GET("/ordered-data/:location/:date", OrderedDataHandler)
	engine.GET("/overview-data/:location/:date", OverviewDataHandler)
}
