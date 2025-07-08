package router

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)


func InitializeRouter(DB *sql.DB) {
	// Initialize Router
	router := gin.Default()
	
	// Initialize Routes
	initializeRoutes(router, DB)

	// Run the server
	router.Run() // listen and serve on 0.0.0.0:8080
}