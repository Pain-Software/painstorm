package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"painstorm/controller"
)

func NewRouter(weatherController *controller.WeatherController) *gin.Engine {
	router := gin.Default()
	router.Use(cors.Default())
	
	baseRouter := router.Group("/api")
	weatherRouter := baseRouter.Group("/weather")
	searchRouter := weatherRouter.Group("/search")

	baseRouter.GET("/version", weatherController.Version)

	searchRouter.GET("/current", weatherController.Current)
	searchRouter.GET("/generate", weatherController.Generate)
	searchRouter.GET("/intensity", weatherController.RainIntensity)

	return router
}
