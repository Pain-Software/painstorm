package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"painstorm/controller"
)

func NewRouter(weatherController *controller.WeatherController) *gin.Engine {
	router := gin.Default()
	router.Use(cors.Default())
	// add swagger
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	baseRouter := router.Group("/api")
	weatherRouter := baseRouter.Group("/weather")
	searchRouter := weatherRouter.Group("/search")

	baseRouter.GET("/version", weatherController.Version)

	searchRouter.GET("/retrieve", weatherController.Retrieve)
	searchRouter.GET("/current", weatherController.Current)
	searchRouter.GET("/generate", weatherController.Generate)
	searchRouter.GET("/intensity", weatherController.RainIntensity)

	return router
}
