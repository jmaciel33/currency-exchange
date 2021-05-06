package web

import (
	"github.com/jmaciel33/currency-exchange/currency-exchange-api/internal/app/factory"
	"github.com/labstack/echo/v4"
	"net/http"
)

var currencyService = factory.GetCacheService()

func GetByIdHandler(ctx echo.Context) error {

	resource, err := currencyService.GetCurrencies()

	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	return ctx.JSON(http.StatusOK, resource)
}