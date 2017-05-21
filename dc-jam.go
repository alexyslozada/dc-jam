package main

import (
	"github.com/labstack/echo/middleware"
	"github.com/labstack/echo"
	"github.com/alexyslozada/dc-jam/routers"
	"os"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Static("/", "public")

	routers.StartAll(e)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	e.Logger.Fatal(e.Start(":" + port))
}
