package controllers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"zadanie4_project/db"
	"zadanie4_project/models"
)

func CreatePayment(c echo.Context) error {
	payment := new(models.Payment)
	if err := c.Bind(payment); err != nil {
		return err
	}
	db.DB.Create(&payment)
	return c.JSON(http.StatusCreated, payment)
}
