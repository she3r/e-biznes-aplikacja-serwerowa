package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"zadanie4_project/db"
	"zadanie4_project/models"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestCreateEmployee(t *testing.T) {
	setupTestDB()
	e := setupEcho()

	employee := models.Employee{
		Name:    "Alicja Nowak",
		Address: "a.nowak@gmail.com",
	}
	body, _ := json.Marshal(employee)
	req := httptest.NewRequest(http.MethodPost, "/employees", bytes.NewBuffer(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err := CreateEmployee(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, rec.Code)

	var created models.Employee
	_ = json.Unmarshal(rec.Body.Bytes(), &created)
	assert.Equal(t, "Alicja Nowak", created.Name)
	assert.Equal(t, "a.nowak@gmail.com", created.Address)
}

func TestGetEmployees(t *testing.T) {
	setupTestDB()
	e := setupEcho()

	db.DB.Create(&models.Employee{Name: "Tomasz Zieliński", Address: "a.nowak@gmail.com"})

	req := httptest.NewRequest(http.MethodGet, "/employees", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err := GetEmployees(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)

	var employees []models.Employee
	_ = json.Unmarshal(rec.Body.Bytes(), &employees)
	assert.Len(t, employees, 1)
	assert.Equal(t, "Tomasz Zieliński", employees[0].Name)
}
