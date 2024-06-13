// Package config provides configuration settings and initialization for the application
package config

// Import the required packages
import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/labstack/gommon/log" // used for logging messages to the console
	"gorm.io/driver/mysql"           // provides the MySQL dialect for GORM
	"gorm.io/gorm"                   // ORM library
	"os"
)

// DB Declare a global variable DB of type *gorm.DB (=> a pointer) to hold the database connection
var DB *gorm.DB

// InitDB Initialize the database connection
func InitDB() {
	err := godotenv.Load()
	if err != nil {
		log.Warn("No .env file found, using default environment variables")
	}
	// Get the database connection details from environment variables
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")
	if port == "" {
		port = "3306" // Default MySQL port
	}

	// Define the Data Source Name (DSN) which includes the database connection details
	// Format: "username:password@protocol(address)/dbname?param=value"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		username, password, host, port, dbname)

	// Open a connection to the database using GORM with the provided DSN
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	// Check if there was an error while connecting to the database
	if err != nil {
		// If there's an error, log a fatal message and exit the program
		log.Fatal("Failed to connect to database: ", err)
	}

	// Log a success message indicating that the database connection was established successfully
	log.Info("Database connected successfully.")
}
