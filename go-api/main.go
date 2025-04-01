package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/s-union/canalia/generated"
	"github.com/s-union/canalia/handler"
	"github.com/s-union/canalia/middleware"
)

func Env() {
	enviroment, ok := os.LookupEnv("GO_ENV")
	if !ok {
		enviroment = "local"
	}
	err := godotenv.Load(fmt.Sprintf(".env.%s", enviroment))
	if err != nil {
		log.Fatalf("Error loading .env.%s file", enviroment)
	}
}

func main() {
	Env()
	server := handler.NewServer()

	e := echo.New()
	e.Use(middleware.JWTAuth)
	generated.RegisterHandlers(e, &server)

	log.Fatal(e.Start(":8080"))
}
