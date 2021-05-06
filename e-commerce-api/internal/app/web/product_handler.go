package web

import (
	"github.com/jmaciel33/currency-exchange/e-commerce-api/internal/app/factory"
	"github.com/jmaciel33/currency-exchange/e-commerce-api/internal/models"
	"github.com/labstack/echo/v4"
	"net/http"
)

var productService = factory.GetProductService()

func CreateHandler(ctx echo.Context) error {

	var data models.AddProduct

	if err := ctx.Bind(&data); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := productService.CreateNewProduct(data); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, map[string]string{
		"message": "Product was added.",
	})
}

func GetByIdHandler(ctx echo.Context) error {
	resource, err := productService.FindProductById(ctx.Param("id"))

	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	return ctx.JSON(http.StatusOK, resource)
}
func ListProductsHandler(ctx echo.Context) error {
	products := productService.ListAll()
	return ctx.JSON(http.StatusOK, products)
}

func UpdateHandler(ctx echo.Context) error {

	var data models.Product

	if err := ctx.Bind(&data); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := productService.UpdateProduct(data); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, map[string]string{
		"message": "Product has been updated.",
	})
}
