package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"

	"github.com/s-union/canalia/internal/api"
	"github.com/s-union/canalia/internal/db"
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

func main() {
	Env()

	// Initialize database connection
	database, err := db.NewDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer database.Close()

	server := api.NewServer(database)

	e := echo.New()
	e.Use(echoMiddleware.Recover())
	e.Use(middleware.JWTAuth)
	
	// Manual route registration
	api.RegisterRoutes(e, server)

	log.Fatal(e.Start(":8080"))
}
