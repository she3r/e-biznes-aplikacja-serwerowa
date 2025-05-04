package controllers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"zadanie4_project/db"
	"zadanie4_project/models"
)

const ProductNotFoundErrorMsg = "Nie znaleziono takiego produktu"

func CreateProduct(c echo.Context) error {
	product := new(models.Product)
	if err := c.Bind(product); err != nil {
		return err
	}

	var category models.Category
	if err := db.DB.First(&category, product.CategoryID).Error; err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": CategoryNotFoundErrorMsg})
	}

	db.DB.Create(&product)
	return c.JSON(http.StatusCreated, product)
}

func GetProducts(c echo.Context) error {
	var products []models.Product
	db.DB.Preload("Category").Find(&products)
	return c.JSON(http.StatusOK, products)
}

func GetProduct(c echo.Context) error {
	id := c.Param("id")
	var product models.Product
	if err := db.DB.First(&product, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"message": "Product not found"})
	}
	return c.JSON(http.StatusOK, product)
}

func UpdateProduct(c echo.Context) error {
	id := c.Param("id")
	var product models.Product
	if err := db.DB.First(&product, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"message": ProductNotFoundErrorMsg})
	}

	updates := map[string]interface{}{}

	if err := c.Bind(&updates); err != nil {
		return err
	}

	if name, ok := updates["Name"].(string); ok {
		product.Name = name
	}
	if price, ok := updates["Price"].(float64); ok {
		product.Price = price
	}
	if categoryID, ok := updates["CategoryID"].(uint); ok {
		var category models.Category
		if err := db.DB.First(&category, categoryID).Error; err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"message": CategoryNotFoundErrorMsg})
		}
		product.CategoryID = categoryID
	}

	db.DB.Save(&product)
	return c.JSON(http.StatusOK, product)
}

func DeleteProduct(c echo.Context) error {
	id := c.Param("id")
	var product models.Product
	if err := db.DB.First(&product, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"message": ProductNotFoundErrorMsg})
	}

	db.DB.Delete(&product)
	return c.JSON(http.StatusOK, echo.Map{"message": "Produkt został usunięty pomyślnie"})
}
