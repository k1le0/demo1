package route

import (
	"demo1/service"
	"github.com/labstack/echo/v4"
)

func Route(e *echo.Echo) {
	e.GET("/", service.Hello)
}
