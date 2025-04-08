package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"zadanie4_project/controllers"
	"zadanie4_project/db"
)

func main() {
	db.Init()

	e := echo.New()

	e.POST("/products", controllers.CreateProduct)
	e.GET("/products", controllers.GetProducts)
	e.GET("/products/:id", controllers.GetProduct)
	e.PUT("/products/:id", controllers.UpdateProduct)
	e.DELETE("/products/:id", controllers.DeleteProduct)

	e.POST("/categories", controllers.CreateCategory)
	e.GET("/categories", controllers.GetCategories)
	e.GET("/categories/:id", controllers.GetCategory)
	e.PUT("/categories/:id", controllers.UpdateCategory)
	e.DELETE("/categories/:id", controllers.DeleteCategory)

	e.POST("/baskets", controllers.CreateBasket)
	e.GET("/baskets/:id", controllers.GetBasket)
	e.POST("/baskets/:basket_id/products/:product_id", controllers.AddProductToBasket)
	e.DELETE("/baskets/:basket_id/products/:product_id", controllers.RemoveProductFromBasket)

	e.POST("/employees", controllers.CreateEmployee)
	e.GET("/employees", controllers.GetEmployees)

	e.POST("/clients", controllers.CreateClient)
	e.GET("/clients", controllers.GetClients)

	e.POST("/payments", controllers.CreatePayment)

	// CORS middleware
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))

	e.Logger.Fatal(e.Start(":8080"))
}
