package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
	"github.com/walter-manger/pb_starter/handlers"
)

func main() {
	app := pocketbase.New()

	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		e.Router.Static("/public", "public")
		e.Router.Static("/favicon.ico", "public/images/favicon/favicon.ico")

		e.Router.GET("/report", func(c echo.Context) error {
			return c.String(http.StatusOK, "Hello world! & More & Update")
		}, apis.ActivityLogger(app))

		e.Router.GET("/", handlers.HomeHandler, apis.ActivityLogger(app))

		return nil
	})

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
