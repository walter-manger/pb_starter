package handlers

import (
	"github.com/walter-manger/pb_starter/views/root"

	"github.com/labstack/echo/v5"
)

var (
	fromProtected bool = false
	isError       bool = false
)

func SetupRoutes(e *echo.Echo) {
	e.GET("/", homeHandler)
}

func homeHandler(c echo.Context) error {
	return renderView(c, root.HomeIndex(
		"| Home",
	))
}

func renderView(c echo.Context, cmp templ.Component) error {
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)

	return cmp.Render(c.Request().Context(), c.Response().Writer)
}
