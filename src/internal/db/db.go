package db

import (
	"context"
	"fmt"
	"os"

	"github.com/Jerinji2016/grpc-template/src/pkg/logger"
	"github.com/jackc/pgx/v5/pgxpool"
)

var DB *pgxpool.Pool

func InitDB() {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	conn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", dbUsername, dbPassword, dbHost, dbPort, dbName)

	var err error
	DB, err = pgxpool.New(context.Background(), conn)
	if err != nil {
		logger.FatalLog("Failed to connect to DB: %s", conn)
	}
	logger.InfoLog("Database connection initialized")
}

func CloseDB() {
	DB.Close()
	logger.InfoLog("Database connection closed")
}
