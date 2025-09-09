package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"

	"github.com/s-union/canalia/internal/api"
	db "github.com/s-union/canalia/internal/db/generated"
	"github.com/s-union/canalia/internal/middleware"
)

func Env() {
	environment, ok := os.LookupEnv("GO_ENV")
	if !ok {
		environment = "local"
	}
	err := godotenv.Load(fmt.Sprintf(".env.%s", environment))
	if err != nil {
		log.Fatalf("Error loading .env.%s file", environment)
	}
}

func ConnectDB() (*db.Queries, *pgxpool.Pool) {
	// Initialize database connection pool
	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		log.Fatal("DATABASE_URL environment variable is required")
	}

	ctx := context.Background()
	dbPool, err := pgxpool.New(ctx, databaseURL)
	if err != nil {
		log.Fatal("Failed to create database pool:", err)
	}

	// Test database connection
	if err := dbPool.Ping(ctx); err != nil {
		log.Fatal("Failed to ping database:", err)
	}

	// Initialize database queries
	queries := db.New(dbPool)
	return queries, dbPool
}

func main() {
	Env()

	queries, dbPool := ConnectDB()
	defer dbPool.Close()

	// Initialize server with database queries
	server := api.NewServer(queries)

	e := echo.New()
	e.Use(echoMiddleware.Recover())
	e.Use(middleware.JWTAuth)

	// Manual route registration
	api.RegisterRoutes(e, server)

	log.Fatal(e.Start(":8080"))
}
