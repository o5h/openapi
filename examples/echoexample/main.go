package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	server := echo.New()

	if err := server.Start(":8080"); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
