package handler

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/feribeirods/api-go-gin/repository"
	"github.com/feribeirods/api-go-gin/request"
	"github.com/gin-gonic/gin"
)

func UpdateGame(DB *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Get the ID param
		idParam := ctx.Param("id")
		id, err := strconv.ParseInt(idParam, 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid ID",
			})
			return
		}

		// Binding JSON
		var req request.GameRequest
		if err := ctx.BindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error":   "Invalid JSON",
				"message": err.Error(),
			})
			return
		}

		// Validate the values passed in the body
		if errs := req.Validate(); len(errs) > 0 {
			var messages []string
			for _, e := range errs {
				messages = append(messages, e.Error())
			}
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error":   "Validation failed",
				"message": messages,
			})
			return
		}

		game := req.ToModel()

		err = repository.UpdateGameByID(DB, id, game)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update", "message": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"message": "Game updated successfully",
			"data":    game,
		})
	}
}
