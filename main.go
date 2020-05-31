package main

import (
	"github.com/labstack/echo"
	"github.com/locpham24/github-trending/handler"
)

func main() {
	e := echo.New()

	handler.InitRouter(e)
	e.Logger.Fatal(e.Start(":7000"))
}
