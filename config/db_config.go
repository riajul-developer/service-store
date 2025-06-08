package config

import (
	"context"
	"database/sql"
	"log"
	"os"
	"service-store/internal/models"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

var DB *bun.DB

func ConnectDB() {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("DATABASE_URL not set in .env")
	}

	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))

	if err := sqldb.Ping(); err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}

	DB = bun.NewDB(sqldb, pgdialect.New())

	modelsToCreate := []interface{}{
		(*models.User)(nil),
		(*models.Role)(nil),
		(*models.Permission)(nil),
		(*models.RolePermission)(nil),
		(*models.Hub)(nil),
		(*models.Vendor)(nil),
		(*models.Category)(nil),
	}

	ctx := context.Background()

	for _, model := range modelsToCreate {
		_, err := DB.NewCreateTable().
			Model(model).
			IfNotExists().
			Exec(ctx)
		if err != nil {
			log.Fatalf("Failed to create table for model %T: %v", model, err)
		}
	}

	log.Println("Database connected successfully")
}
