package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateGame (ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "Update an Existant Game",
	})
}