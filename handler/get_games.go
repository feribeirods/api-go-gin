package handler

import (
	"database/sql"
	"net/http"

	"github.com/feribeirods/api-go-gin/repository"
	"github.com/gin-gonic/gin"
)

func HandleGetGames(DB *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		games, err := repository.FindAllGames(DB)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error":   "Failed to fetch games",
				"message": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"games": games,
		})
	}
}
