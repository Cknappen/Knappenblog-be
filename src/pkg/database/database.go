package database

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	DB *gorm.DB
)

func SetLocal[T any](c *fiber.Ctx, key string, value T) {
	c.Locals(key, value)
}
func GetLocal[T any](c *fiber.Ctx, key string) T {
	return c.Locals(key).(T)
}

func InitDB() *gorm.DB {
	// Load environment variables from the .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file, check example.env for requiered fields.")
	}

	// Read requiered fields for establishing a db connection with GORM.
	pgHostname := os.Getenv("PG_HOSTNAME")
	pgUserName := os.Getenv("PG_USERNAME")
	pgPassword := os.Getenv("PG_PASSWORD")
	pgDatabase := os.Getenv("PG_DATABASE")
	pgPort := os.Getenv("PG_PORT")
	pgSslMode := os.Getenv("PG_SSLMODE")
	pgTimeZone := os.Getenv("PG_TIMEZONE")

	// Dsn as listed from GORM docs for postgres: https://gorm.io/docs/connecting_to_the_database.html
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s", pgHostname, pgUserName, pgPassword, pgDatabase, pgPort, pgSslMode, pgTimeZone)

	// Try to establish connection to Postgresql Database.
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent), // You can set the log mode as needed
	})
	// Database not found.
	if err != nil {
		panic("Failed to connect to database")
	}

	DB = db

	return DB
}
