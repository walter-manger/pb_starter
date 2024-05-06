package handlers

import (
	"github.com/a-h/templ"
	"github.com/labstack/echo/v5"
	"github.com/walter-manger/pb_starter/views/home"
)

var (
	fromProtected bool = false
	isError       bool = false
)

func SetupRoutes(e *echo.Echo) {
	e.GET("/", HomeHandler)
}

func HomeHandler(c echo.Context) error {
	homeView := home.Home()
	return renderView(c, home.HomeIndex(
		"| Home",
		homeView,
	))
}

func renderView(c echo.Context, cmp templ.Component) error {
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)

	return cmp.Render(c.Request().Context(), c.Response().Writer)
}
