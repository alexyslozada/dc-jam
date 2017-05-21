package routers

import "github.com/labstack/echo"

func StartAll(e *echo.Echo) {
	StartDesplazamiento(e)
}
