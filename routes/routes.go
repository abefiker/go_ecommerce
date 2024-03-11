package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/abefiker/go_ecommerce/internals/handlers"
)

func RegisterRoutes(e *echo.Echo) {
	productGroup := e.Group("/product")
	productGroup.GET("/product",handlers.Product)
	productGroup.GET("/list/:id",handlers.GetProduct)
	productGroup.GET("/lists",handlers.GetAllUsers)
	productGroup.POST("/create", handlers.CreateProduct)
	productGroup.PUT("/update/:id", handlers.UpdateProduct)
	productGroup.DELETE("/delete/:id",handlers.DeleteProduct)
}