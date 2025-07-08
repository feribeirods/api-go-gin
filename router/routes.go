package router

import (
	"github.com/feribeirods/api-go-gin/handler"
	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")

	v1.GET("/games", handler.HandleGetGames)
	v1.POST("/games", handler.PostNewGame)
	v1.PUT("/games", handler.UpdateGame)
	v1.DELETE("/games", handler.DeleteGame)
}