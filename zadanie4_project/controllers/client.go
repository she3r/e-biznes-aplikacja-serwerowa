package controllers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"zadanie4_project/db"
	"zadanie4_project/models"
)

func CreateClient(c echo.Context) error {
	client := new(models.Client)
	if err := c.Bind(client); err != nil {
		return err
	}
	if err := db.DB.Create(&client).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "Nie udało się utworzyć klienta"})
	}
	return c.JSON(http.StatusCreated, client)
}

func GetClients(c echo.Context) error {
	var clients []models.Client
	if err := db.DB.Find(&clients).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "Nie udało się wydobyć klientów"})
	}
	return c.JSON(http.StatusOK, clients)
}
