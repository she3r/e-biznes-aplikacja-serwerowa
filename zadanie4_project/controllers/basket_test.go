package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
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

	err := CreateBasket(c)

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

	err := GetBasket(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
}

func insertTestData(t *testing.T) (models.Basket, models.Product) {
	basket := models.Basket{}
	product := models.Product{Name: "Test Product", Price: 10.0}

	if err := db.DB.Create(&basket).Error; err != nil {
		t.Fatalf("failed to create basket: %v", err)
	}
	if err := db.DB.Create(&product).Error; err != nil {
		t.Fatalf("failed to create product: %v", err)
	}
	return basket, product
}

func TestAddProductToBasket_Success(t *testing.T) {
	setupTestDB()
	e := setupEcho()

	basket, product := insertTestData(t)

	req := httptest.NewRequest(http.MethodPost, "/baskets/:basket_id/products/:product_id", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("basket_id", "product_id")
	c.SetParamValues(toStr(basket.ID), toStr(product.ID))

	err := AddProductToBasket(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)

	// Verify association
	var updatedBasket models.Basket
	err = db.DB.Preload("Products").First(&updatedBasket, basket.ID).Error
	assert.NoError(t, err)
	assert.Len(t, updatedBasket.Products, 1)
	assert.Equal(t, product.ID, updatedBasket.Products[0].ID)
}

func TestRemoveProductFromBasket_Success(t *testing.T) {
	setupTestDB()
	e := setupEcho()

	basket, product := insertTestData(t)
	// Add product to basket first
	db.DB.Model(&basket).Association("Products").Append(&product)

	req := httptest.NewRequest(http.MethodDelete, "/baskets/:basket_id/products/:product_id", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("basket_id", "product_id")
	c.SetParamValues(toStr(basket.ID), toStr(product.ID))

	err := RemoveProductFromBasket(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)

	// Verify removal
	var updatedBasket models.Basket
	err = db.DB.Preload("Products").First(&updatedBasket, basket.ID).Error
	assert.NoError(t, err)
	assert.Len(t, updatedBasket.Products, 0)
}

func TestAddProductToBasket_BasketNotFound(t *testing.T) {
	setupTestDB()
	e := setupEcho()

	product := models.Product{Name: "Test Product", Price: 10.0}
	db.DB.Create(&product)

	req := httptest.NewRequest(http.MethodPost, "/baskets/:basket_id/products/:product_id", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("basket_id", "product_id")
	c.SetParamValues("999", toStr(product.ID))

	err := AddProductToBasket(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusNotFound, rec.Code)
}

func TestRemoveProductFromBasket_ProductNotFound(t *testing.T) {
	setupTestDB()
	e := setupEcho()

	basket := models.Basket{}
	db.DB.Create(&basket)

	req := httptest.NewRequest(http.MethodDelete, "/baskets/:basket_id/products/:product_id", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("basket_id", "product_id")
	c.SetParamValues(toStr(basket.ID), "999")

	err := RemoveProductFromBasket(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusNotFound, rec.Code)
}

// Helper to convert uint to string
func toStr(id uint) string {
	return fmt.Sprintf("%d", id)
}
