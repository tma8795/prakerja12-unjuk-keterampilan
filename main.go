package main

import (
	"ecommerce/configs"
	"ecommerce/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	configs.InitDatabase()
	e := echo.New()
	routes.InitRoutes(e)
	e.Start(":8000")
}
