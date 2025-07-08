package main

import (
	"github.com/feribeirods/api-go-gin/config"
	"github.com/feribeirods/api-go-gin/router"
)

func main() {
	//Initialize DB
	DB := config.InitDB()

	// Close DB Connection
	defer DB.Close()

	// Initialize Router
	router.InitializeRouter(DB)	
}