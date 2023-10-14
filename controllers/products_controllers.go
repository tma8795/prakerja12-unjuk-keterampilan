package controllers

import (
	"ecommerce/configs"
	"ecommerce/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func CreateProductsController(c echo.Context) error {
	var productRequest models.Products
	c.Bind(&productRequest)

	result := configs.DB.Create(&productRequest)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Message: result.Error.Error(),
			Status:  false,
			Data:    nil,
		})
	}
	return c.JSON(http.StatusOK, models.BaseResponse{
		Message: "Berhasil menambah data",
		Status:  true,
		Data:    productRequest,
	})
}

// func GetProductsController(c echo.Context) error    {}
// func GetProductById(c echo.Context) error           {}
// func UpdateProductsController(c echo.Context) error {}
// func DeleteProductsController(c echo.Context) error {}
