package routes

import (
	"ecommerce/controllers"

	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo) {
	e.POST("/products", controllers.CreateProductsController)
	// e.GET("/products", controllers.GetProductsController)
	// e.GET("/products/:id", controllers.GetProductById)
	// e.PUT("/products/:id", controllers.UpdateProductsController)
	// e.DELETE("/products/:id", controllers.DeleteProductsController)
}
