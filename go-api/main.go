package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/s-union/canalia/generated"
	"github.com/s-union/canalia/handler"
)

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func main() {
	server := handler.NewServer()

	e := echo.New()
	generated.RegisterHandlers(e, &server)

	log.Fatal(e.Start(":8080"))
}
