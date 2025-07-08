package handler

import (
	"database/sql"

	"net/http"

	"github.com/feribeirods/api-go-gin/repository"
	"github.com/feribeirods/api-go-gin/request"
	"github.com/gin-gonic/gin"
)

func PostNewGame(DB *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req request.GameRequest

		// Binding JSON
		if err := ctx.BindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error":   "Invalid JSON",
				"message": err.Error(),
			})
			return
		}

		// Validate the values recieved
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

		// Translate to the model
		game := req.ToModel()

		// Call the repository to insert the data in the DB
		if err := repository.InsertGame(DB, game); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error":   "Database error",
				"message": err.Error(),
			})
			return
		}

		// Send a success message to the client
		ctx.JSON(http.StatusCreated, gin.H{
			"message": "Game created successfully",
			"data":    game,
		})
	}
}
