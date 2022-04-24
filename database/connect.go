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
    // var err error
    // p := config.Config("DB_PORT")
    // port, err := strconv.ParseUint(p, 10, 32)

    // if err != nil {
	// 	panic("failed to connect database")
	// }	
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	DB_HOST := os.Getenv("DB_HOST")
	DB_NAME := os.Getenv("DB_NAME")
	DB_USER := os.Getenv("DB_USER")
	DB_PORT := os.Getenv("DB_PORT")
	DB_PASSWORD := os.Getenv("DB_PASSWORD")

	psqlInfo := fmt.Sprintf("host=%s user=%s dbname=%s port=%s password=%s", DB_HOST, DB_USER, DB_NAME, DB_PORT, DB_PASSWORD)
	db, err := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{
		SkipDefaultTransaction: true,
	})

	if err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}
//     // Connection URL to connect to Postgres Database
//     dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", config.Config("DB_HOST"), port, config.Config("DB_USER"), config.Config("DB_PASSWORD"), config.Config("DB_NAME"))
// 	DB, err = gorm.Open(postgres.Open(dsn))
// // DB, err = gorm.Open("sqlite3", "/tmp/gorm.db")
//     if err != nil {
//         panic(err)
//     }

    fmt.Println("Connection Opened to Database")

	return db
}

func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(
		&model.User{},
		&model.Memo{},
		&model.Comment{},
	)
	fmt.Println("Database Migrated")
}

