package main

import (
	"github.com/ameykpatil/go-uploader/constants"
	"github.com/ameykpatil/go-uploader/controllers"
	"github.com/gin-gonic/gin"
	"net/http"
)

func getMainEngine() *gin.Engine {
	router := gin.New() // Create a gin router without any default middleware.

	// Add recovery (crash-free) middleware
	router.Use(gin.Recovery())

	// Add logger middleware only when running in debug mode (GIN_MODE env variable)
	if gin.Mode() == gin.DebugMode {
		router.Use(gin.Logger())
	}

	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	assetGroup := router.Group("/asset")
	{
		assetGroup.GET("/:assetId", controllers.GetAsset)
		assetGroup.POST("", controllers.AddAsset)
		assetGroup.PUT("/:assetId", controllers.UpdateAsset)
	}

	return router
}

func main() {
	router := getMainEngine()
	router.Run(":" + constants.Env.Port)
}
