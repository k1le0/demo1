package main

import (
	"demo1/route"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	route.Route(e)

	_ = e.Start(":8080")
}
