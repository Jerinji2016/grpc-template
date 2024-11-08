package db

import (
	"fmt"
	"os"

	"github.com/Jerinji2016/grpc-template/src/internal/models"
	"github.com/Jerinji2016/grpc-template/src/pkg/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

var schemas = []interface{} {
	&models.User{}, 
	&models.Post{},
}

func InitDB() {
	dns := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	var err error
	DB, err = gorm.Open(postgres.Open(dns), &gorm.Config{})
	if err != nil {
		logger.FatalLog("Failed to connect to DB: %s", dns)
	}
	logger.InfoLog("Database connection initialized")

	if err := DB.AutoMigrate(schemas...); err != nil {
		logger.FatalLog("Failed to migrate Database: %v", err)
	}
}

func CloseDB() {
	db, err := DB.DB()
	if err != nil {
		logger.FatalLog("Failed to get DB")
	}
	db.Close()
	logger.InfoLog("Database connection closed")
}
