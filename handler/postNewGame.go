package handler

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PostNewGame  (DB *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "Post New game",
		})
	}
}