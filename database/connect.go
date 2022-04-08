package database

import (
	"fmt"
	"strconv"

	"github.com/Chotiwitorratai/cloudmemo_backend/config"
	"github.com/Chotiwitorratai/cloudmemo_backend/internal/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// ConnectDB connect to db
func ConnectDB() {
    var err error
    p := config.Config("DB_PORT")
    port, err := strconv.ParseUint(p, 10, 32)

    if err != nil {
		panic("failed to connect database")
	}

    // Connection URL to connect to Postgres Database
    dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", config.Config("DB_HOST"), port, config.Config("DB_USER"), config.Config("DB_PASSWORD"), config.Config("DB_NAME"))
	DB, err = gorm.Open(postgres.Open(dsn))

    if err != nil {
        panic(err)
    }

    fmt.Println("Connection Opened to Database")

    // Migrate the database
    DB.AutoMigrate(&model.Memo{})
    DB.AutoMigrate(&model.User{})
    fmt.Println("Database Migrated")

}

