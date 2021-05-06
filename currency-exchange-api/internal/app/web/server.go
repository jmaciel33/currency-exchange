package web

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

func Server() *echo.Echo {
	app := echo.New()

	app.Use(middleware.Recover())
	app.Use(middleware.Gzip())

	app.GET("/", func(ctx echo.Context) error {
		return ctx.JSON(http.StatusOK, map[string]string{
			"status": "OK",
		})
	})

	crud := app.Group("/v1/currency")
	{
		crud.GET("", GetByIdHandler)
	}

	return app
}