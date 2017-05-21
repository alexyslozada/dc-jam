package main

import (
	"github.com/labstack/echo/middleware"
	"github.com/labstack/echo"
	"github.com/alexyslozada/dc-jam/routers"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Static("/", "public")

	routers.StartAll(e)

	e.Logger.Fatal(e.Start(":8080"))
}
