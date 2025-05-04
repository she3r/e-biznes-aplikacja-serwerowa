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

func TestCreateCategory(t *testing.T) {
	setupTestDB()
	e := setupEcho()

	category := models.Category{Name: "Test Category"}
	body, _ := json.Marshal(category)
	req := httptest.NewRequest(http.MethodPost, "/category", bytes.NewBuffer(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err := CreateCategory(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, rec.Code)

	var created models.Category
	_ = json.Unmarshal(rec.Body.Bytes(), &created)
	assert.Equal(t, "Test Category", created.Name)
}

func TestGetCategories(t *testing.T) {
	setupTestDB()
	e := setupEcho()

	db.DB.Create(&models.Category{Name: "Test Category"})

	req := httptest.NewRequest(http.MethodGet, "/categories", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err := GetCategories(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)

	var categories []models.Category
	_ = json.Unmarshal(rec.Body.Bytes(), &categories)
	assert.Len(t, categories, 1)
	assert.Equal(t, "Test Category", categories[0].Name)
}

func TestGetCategory(t *testing.T) {
	setupTestDB()
	e := setupEcho()

	category := models.Category{Name: "Test Category"}
	db.DB.Create(&category)

	req := httptest.NewRequest(http.MethodGet, "/category/"+strconv.Itoa(int(category.ID)), nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues(strconv.Itoa(int(category.ID)))

	err := GetCategory(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)

	var fetched models.Category
	_ = json.Unmarshal(rec.Body.Bytes(), &fetched)
	assert.Equal(t, "Test Category", fetched.Name)
}

func TestUpdateCategory(t *testing.T) {
	setupTestDB()
	e := setupEcho()

	category := models.Category{Name: "Old Name"}
	db.DB.Create(&category)

	updated := models.Category{Name: "New Name"}
	body, _ := json.Marshal(updated)
	req := httptest.NewRequest(http.MethodPut, "/category/"+strconv.Itoa(int(category.ID)), bytes.NewBuffer(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues(strconv.Itoa(int(category.ID)))

	err := UpdateCategory(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)

	var fetched models.Category
	_ = json.Unmarshal(rec.Body.Bytes(), &fetched)
	assert.Equal(t, "New Name", fetched.Name)
}

func TestDeleteCategory(t *testing.T) {
	setupTestDB()
	e := setupEcho()

	category := models.Category{Name: "To Delete"}
	db.DB.Create(&category)

	req := httptest.NewRequest(http.MethodDelete, "/category/"+strconv.Itoa(int(category.ID)), nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues(strconv.Itoa(int(category.ID)))

	err := DeleteCategory(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)

	var response map[string]string
	_ = json.Unmarshal(rec.Body.Bytes(), &response)
	assert.Equal(t, "Kategoria produktu została usunięta pomyślnie", response["message"])
}
