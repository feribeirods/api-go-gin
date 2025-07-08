package router

import (
	"database/sql"

	"github.com/feribeirods/api-go-gin/handler"
	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine, DB *sql.DB) {
	v1 := router.Group("/api/v1")

	v1.GET("/games", handler.HandleGetGames(DB))
	v1.POST("/games", handler.PostNewGame(DB))
	v1.PUT("/games/:id", handler.UpdateGame(DB))
	v1.DELETE("/games/:id", handler.DeleteGame(DB))
}
