package controllers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"zadanie4_project/db"
	"zadanie4_project/models"
)

func CreateEmployee(c echo.Context) error {
	employee := new(models.Employee)
	if err := c.Bind(employee); err != nil {
		return err
	}
	if err := db.DB.Create(&employee).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "Nie udało się utworzyć pracownika"})
	}
	return c.JSON(http.StatusCreated, employee)
}

func GetEmployees(c echo.Context) error {
	var employees []models.Employee
	if err := db.DB.Find(&employees).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "Nie udało się wydobyć pracowników"})
	}
	return c.JSON(http.StatusOK, employees)
}
