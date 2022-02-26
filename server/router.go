package server

import (
	"github.com/gin-gonic/gin"

	"web-service-gin/config"
	"web-service-gin/controllers"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	config := config.GetConfig()
	router.SetTrustedProxies([]string{config.GetString("server.trustedProxies")})

	health := new(controllers.HealthController)

	router.GET("/health", health.Status)

	v1 := router.Group("v1")
	{
		albumGroup := v1.Group("album")
		{
			album := new(controllers.AlbumController)
			albumGroup.GET("/", album.GetAlbums)
			albumGroup.GET("/:id", album.GetAlbumByID)
			albumGroup.POST("/", album.PostAlbums)
		}
	}

	return router
}
