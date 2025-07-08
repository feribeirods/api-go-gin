package handler

import (
	"database/sql"

	"net/http"
	"strconv"

	"github.com/feribeirods/api-go-gin/repository"
	"github.com/gin-gonic/gin"
)

func DeleteGame(DB *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Get the ID param
		idParam := ctx.Param("id")
		id, err := strconv.ParseInt(idParam, 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error":   "Invalid ID format",
				"message": err.Error(),
			})
			return
		}

		// Call the repository to delete the game from the DB
		err = repository.DeleteGameByID(DB, id)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error":   "Failed to delete game",
				"message": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"message": "Game deleted successfully",
			"id":      id,
		})
	}
}
