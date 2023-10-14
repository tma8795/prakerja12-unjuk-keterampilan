package controllers

import (
	"ecommerce/models"
	"ecommerce/repositories"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func CreateProductsController(c echo.Context) error {
	var productRequest models.Products
	c.Bind(&productRequest)

	err := repositories.AddProducts(&productRequest)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Message: err.Error(),
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

func GetProductsController(c echo.Context) error {
	var products []models.Products
	err := repositories.GetProducts(&products)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Message: err.Error(),
			Status:  false,
			Data:    nil,
		})
	}
	return c.JSON(http.StatusOK, models.BaseResponse{
		Message: "Berhasil mengambil data",
		Status:  true,
		Data:    products,
	})
}

func GetProductByIdController(c echo.Context) error {
	var err error
	var product models.Products
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Message: err.Error(),
			Status:  false,
			Data:    nil,
		})
	}

	product.Id = uint(id)

	err = repositories.GetProduct(&product)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, models.BaseResponse{
		Message: "Berhasil mengambil data",
		Status:  true,
		Data:    product,
	})
}

// func UpdateProductsController(c echo.Context) error {}

func DeleteProductsController(c echo.Context) error {
	var err error
	var product models.Products
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Message: err.Error(),
			Status:  false,
			Data:    nil,
		})
	}

	product.Id = uint(id)

	err = repositories.DeleteProduct(&product)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Message: err.Error(),
			Status:  false,
			Data:    nil,
		})
	}
	return c.JSON(http.StatusOK, models.BaseResponse{
		Message: "Berhasil menghapus data",
		Status:  true,
		Data:    product.Id,
	})
}
