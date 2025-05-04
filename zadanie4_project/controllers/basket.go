package controllers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"zadanie4_project/db"
	"zadanie4_project/models"
)

const BasketNotFoundErrorMsg = "Nie znaleziono koszyka"

func CreateBasket(c echo.Context) error {
	basket := new(models.Basket)
	if err := db.DB.Create(&basket).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "Nie udało się utworzyć koszyka"})
	}
	return c.JSON(http.StatusCreated, basket)
}

func AddProductToBasket(c echo.Context) error {
	basketID, productID := GetBasketIdAndProductId(c)

	basket, product, err, isErr := GetBasketAndProduct(c, basketID, productID)
	if isErr {
		return err
	}

	db.DB.Model(&basket).Association("Products").Append(&product)
	return c.JSON(http.StatusOK, basket)
}

func RemoveProductFromBasket(c echo.Context) error {
	basketID, productID := GetBasketIdAndProductId(c)

	basket, product, err, isErr := GetBasketAndProduct(c, basketID, productID)
	if isErr {
		return err
	}

	db.DB.Model(&basket).Association("Products").Delete(&product)
	return c.JSON(http.StatusOK, basket)
}

func GetBasket(c echo.Context) error {
	id := c.Param("id")
	var basket models.Basket
	if err := db.DB.Preload("Products").First(&basket, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"message": BasketNotFoundErrorMsg})
	}
	return c.JSON(http.StatusOK, basket)
}

func GetBasketAndProduct(c echo.Context, basketID string, productID string) (models.Basket, models.Product, error, bool) {
	var basket models.Basket
	if err := db.DB.First(&basket, basketID).Error; err != nil {
		return models.Basket{}, models.Product{}, c.JSON(http.StatusNotFound, echo.Map{"message": BasketNotFoundErrorMsg}), true
	}

	var product models.Product
	if err := db.DB.First(&product, productID).Error; err != nil {
		return models.Basket{}, models.Product{}, c.JSON(http.StatusNotFound, echo.Map{"message": ProductNotFoundErrorMsg}), true
	}
	return basket, product, nil, false
}

func GetBasketIdAndProductId(c echo.Context) (string, string) {
	basketID := c.Param("basket_id")
	productID := c.Param("product_id")
	return basketID, productID
}
