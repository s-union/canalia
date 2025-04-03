package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"

	"github.com/s-union/canalia/generated"
	"github.com/s-union/canalia/handler"
	"github.com/s-union/canalia/middleware"
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
	server := handler.NewServer()

	e := echo.New()
	e.Use(echoMiddleware.Recover())
	e.Use(middleware.JWTAuth)
	generated.RegisterHandlers(e, &server)

	log.Fatal(e.Start(":8080"))
}
