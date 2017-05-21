package routers

import (
	"github.com/labstack/echo"
	"github.com/alexyslozada/dc-jam/handlers"
)

func StartDesplazamiento(e *echo.Echo) {
	desplazamiento := e.Group("/api/desplazamiento")
	desplazamiento.GET("", handlers.Desplazamiento)
}

func StartEducativas(e *echo.Echo) {
	educativas := e.Group("/api/educativas")
	educativas.GET("", handlers.Educativas)
}
