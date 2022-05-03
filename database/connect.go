package database

import (
	"fmt"
	"log"
	"os"

	"github.com/Chotiwitorratai/cloudmemo_backend/model"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)
var DB *gorm.DB

// ConnectDB connect to db
func ConnectDB() *gorm.DB{
   
 	err := godotenv.Load()
  	if err != nil {
    log.Fatalf("Error while reading config file %s", err)
  	}
    if err != nil {
		panic("failed to connect database")
	}
	DB_HOST := os.Getenv("DB_HOST")
  	DB_NAME := os.Getenv("DB_NAME")
 	DB_USER := os.Getenv("DB_USER")
 	DB_PORT := os.Getenv("DB_PORT")
  	DB_PASSWORD := os.Getenv("DB_PASSWORD")
 	dsn := fmt.Sprintf("host=%s user=%s dbname=%s port=%s password=%s", DB_HOST, DB_USER, DB_NAME, DB_PORT, DB_PASSWORD)

    // Connection URL to connect to Postgres Database
	DB, err = gorm.Open(postgres.Open(dsn))
    if err != nil {
        panic(err)
    }
	DB.AutoMigrate(
		&model.User{},
		&model.Memo{},
		&model.Comment{},
	)
	fmt.Println("Database Migrated")
    fmt.Println("Connection Opened to Database")

	return DB
}

