package main

import (
	"github.com/labstack/echo"
	"github.com/locpham24/github-trending/db"
	"github.com/locpham24/github-trending/handler"
	"os"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		//log.Fatal("$PORT must be set")
		port = "7000"
	}

	DB := &db.Sql{}
	DB.Connect()
	defer DB.Close()

	e := echo.New()

	handler.InitRouter(e)
	e.Logger.Fatal(e.Start(":" + port))
}
