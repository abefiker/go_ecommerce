package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/abefiker/go_ecommerce/internals/config"
	"github.com/abefiker/go_ecommerce/routes"
)

func main() {
	config.Databaseinit()
	
	e := echo.New()

	e.Use(middleware.Logger())
    e.Use(middleware.Recover())

	routes.RegisterRoutes(e)
	e.Logger.Fatal(e.Start(":1323"))
}
