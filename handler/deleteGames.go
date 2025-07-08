package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)


func DeleteGame (ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "Delete an Existant Game",
	})
}