package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func InitDB() *sql.DB {
	// Load .env files
	err := godotenv.Load()
	
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Set the DB variables
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")

	// Create the DB connection string
	connString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&charset=utf8mb4&loc=Local", dbUser, dbPass, dbHost, dbPort, dbName)

	// Open the DB connection
	DB, err := sql.Open("mysql", connString)

	if err != nil {
		log.Fatal("Error connecting database")
	}

	// Connection Pooling and Timeout
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)
	DB.SetConnMaxLifetime(time.Minute * 5)
	
	// Test Connection
	DB.Ping()

	fmt.Println("Successfully connected to database!")

	return DB
}