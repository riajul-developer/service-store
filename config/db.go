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
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))
	DB = bun.NewDB(sqldb, pgdriver.New())
	log.Println("Database connected")
}
