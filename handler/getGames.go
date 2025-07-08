package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)



func HandleGetGames (ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "GET All Games",
	})
}