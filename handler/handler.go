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

func PostNewGame (ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "POST New Game",
	})
}

func UpdateGame (ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "Update an Existant Game",
	})
}

func DeleteGame (ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "Delete an Existant Game",
	})
}