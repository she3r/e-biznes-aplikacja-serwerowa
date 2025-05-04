package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"zadanie4_project/db"
	"zadanie4_project/models"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestCreateProduct(t *testing.T) {
	setupTestDB()
	e := setupEcho()

	// Create a category first
	category := models.Category{Name: "Electronics"}
	db.DB.Create(&category)

	product := models.Product{
		Name:       "Smartphone",
		Price:      999.99,
		CategoryID: category.ID,
	}
	body, _ := json.Marshal(product)
	req := httptest.NewRequest(http.MethodPost, "/products", bytes.NewBuffer(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err := CreateProduct(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, rec.Code)

	var created models.Product
	_ = json.Unmarshal(rec.Body.Bytes(), &created)
	assert.Equal(t, "Smartphone", created.Name)
	assert.Equal(t, category.ID, created.CategoryID)
}

func TestGetProducts(t *testing.T) {
	setupTestDB()
	e := setupEcho()

	category := models.Category{Name: "Books"}
	db.DB.Create(&category)

	db.DB.Create(&models.Product{Name: "Go Programming", Price: 49.99, CategoryID: category.ID})

	req := httptest.NewRequest(http.MethodGet, "/products", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err := GetProducts(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)

	var products []models.Product
	_ = json.Unmarshal(rec.Body.Bytes(), &products)
	assert.Len(t, products, 1)
	assert.Equal(t, "Go Programming", products[0].Name)
}

func TestGetProduct(t *testing.T) {
	setupTestDB()
	e := setupEcho()

	category := models.Category{Name: "Toys"}
	db.DB.Create(&category)

	product := models.Product{Name: "Lego Set", Price: 59.99, CategoryID: category.ID}
	db.DB.Create(&product)

	req := httptest.NewRequest(http.MethodGet, "/products/"+strconv.Itoa(int(product.ID)), nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues(strconv.Itoa(int(product.ID)))

	err := GetProduct(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)

	var fetched models.Product
	_ = json.Unmarshal(rec.Body.Bytes(), &fetched)
	assert.Equal(t, "Lego Set", fetched.Name)
}

func TestUpdateProduct(t *testing.T) {
	setupTestDB()
	e := setupEcho()

	category := models.Category{Name: "Home"}
	newCategory := models.Category{Name: "Garden"}
	db.DB.Create(&category)
	db.DB.Create(&newCategory)

	product := models.Product{Name: "Lamp", Price: 29.99, CategoryID: category.ID}
	db.DB.Create(&product)

	updates := map[string]interface{}{
		"Name":       "LED Lamp",
		"Price":      34.99,
		"CategoryID": newCategory.ID,
	}
	body, _ := json.Marshal(updates)
	req := httptest.NewRequest(http.MethodPut, "/products/"+strconv.Itoa(int(product.ID)), bytes.NewBuffer(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues(strconv.Itoa(int(product.ID)))

	err := UpdateProduct(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)

	var updated models.Product
	_ = json.Unmarshal(rec.Body.Bytes(), &updated)
	assert.Equal(t, "LED Lamp", updated.Name)
	assert.Equal(t, 34.99, updated.Price)
}

func TestDeleteProduct(t *testing.T) {
	setupTestDB()
	e := setupEcho()

	category := models.Category{Name: "Office"}
	db.DB.Create(&category)

	product := models.Product{Name: "Notebook", Price: 5.99, CategoryID: category.ID}
	db.DB.Create(&product)

	req := httptest.NewRequest(http.MethodDelete, "/products/"+strconv.Itoa(int(product.ID)), nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues(strconv.Itoa(int(product.ID)))

	err := DeleteProduct(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)

	var response map[string]string
	_ = json.Unmarshal(rec.Body.Bytes(), &response)
	assert.Equal(t, "Produkt został usunięty pomyślnie", response["message"])
}
