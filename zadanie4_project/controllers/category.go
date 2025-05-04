package controllers

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
	"zadanie4_project/db"
	"zadanie4_project/models"
)

const CategoryNotFoundErrorMsg = "Nie znaleziono kategorii produktu"

func CreateCategory(c echo.Context) error {
	category := new(models.Category)
	if err := c.Bind(category); err != nil {
		return err
	}
	db.DB.Create(&category)
	return c.JSON(http.StatusCreated, category)
}

func GetCategories(c echo.Context) error {
	var categories []models.Category
	// po to zeby wyswietlaly sie produkty w relacji - trzeba przeladowac
	getPreloadedDb().Find(&categories)
	return c.JSON(http.StatusOK, categories)
}

func getPreloadedDb() *gorm.DB {
	return db.DB.Preload("Products")
}

func GetCategory(c echo.Context) error {
	id := c.Param("id")
	var category models.Category
	if err := getPreloadedDb().First(&category, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"message": CategoryNotFoundErrorMsg})
	}
	return c.JSON(http.StatusOK, category)
}

func UpdateCategory(c echo.Context) error {
	id := c.Param("id")
	var category models.Category
	if err := getPreloadedDb().First(&category, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"message": CategoryNotFoundErrorMsg})
	}

	if err := c.Bind(&category); err != nil {
		return err
	}

	db.DB.Save(&category)
	return c.JSON(http.StatusOK, category)
}

func DeleteCategory(c echo.Context) error {
	id := c.Param("id")
	var category models.Category
	if err := getPreloadedDb().First(&category, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"message": CategoryNotFoundErrorMsg})
	}

	db.DB.Delete(&category)
	return c.JSON(http.StatusOK, echo.Map{"message": "Kategoria produktu została usunięta pomyślnie"})
}
