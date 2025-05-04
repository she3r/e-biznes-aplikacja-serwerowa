package controllers_tests

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"zadanie4_project/controllers"
	"zadanie4_project/db"
	"zadanie4_project/models"
)

func setupTestDB() {
	database, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	_ = database.AutoMigrate(&models.Basket{}, &models.Product{}, &models.Client{}, &models.Category{}, &models.Employee{}, &models.Payment{})
	db.DB = database
}

func setupEcho() *echo.Echo {
	e := echo.New()
	return e
}

func TestCreateBasket(t *testing.T) {
	setupTestDB()
	e := setupEcho()

	req := httptest.NewRequest(http.MethodPost, "/basket", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err := controllers.CreateBasket(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, rec.Code)

	var basket models.Basket
	err = json.Unmarshal(rec.Body.Bytes(), &basket)
	assert.NoError(t, err)
	assert.NotZero(t, basket.ID)
}

func TestGetBasket(t *testing.T) {
	setupTestDB()
	e := setupEcho()

	// Create a basket first
	basket := models.Basket{}
	db.DB.Create(&basket)
	basketId := strconv.Itoa(int(basket.ID))
	req := httptest.NewRequest(http.MethodGet, "/basket/"+basketId, nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues(basketId)

	err := controllers.GetBasket(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
}
