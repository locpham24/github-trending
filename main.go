package main

import (
	"github.com/labstack/echo"
	"github.com/locpham24/github-trending/handler"
	"log"
	"os"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	e := echo.New()

	handler.InitRouter(e)
	e.Logger.Fatal(e.Start(":" + port))
}
