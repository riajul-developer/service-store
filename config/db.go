package config

import (
	"database/sql"
	"log"
	"os"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/driver/pgdriver"
)

var DB *bun.DB

func ConnectDB() {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("DATABASE_URL not set in .env")
	}

	// Create the PostgreSQL connector
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))

	// Initialize Bun with the PostgreSQL connector
	DB = bun.NewDB(sqldb, nil) // Removed pgdriver.New() here

	// Check if the database is reachable
	if err := DB.Ping(); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	log.Println("Database connected successfully")
}
