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

func TestCreateClient(t *testing.T) {
	setupTestDB()
	e := setupEcho()

	client := models.Client{
		Name:    "Jan Kowalski",
		Address: "jan@example.com",
	}
	body, _ := json.Marshal(client)
	req := httptest.NewRequest(http.MethodPost, "/clients", bytes.NewBuffer(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err := CreateClient(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, rec.Code)

	var created models.Client
	_ = json.Unmarshal(rec.Body.Bytes(), &created)
	assert.Equal(t, "Jan Kowalski", created.Name)
	assert.Equal(t, "jan@example.com", created.Address)
}

func TestGetClients(t *testing.T) {
	setupTestDB()
	e := setupEcho()

	// Insert a sample client
	db.DB.Create(&models.Client{Name: "Anna Nowak", Address: "anna@example.com"})

	req := httptest.NewRequest(http.MethodGet, "/clients", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err := GetClients(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)

	var clients []models.Client
	_ = json.Unmarshal(rec.Body.Bytes(), &clients)
	assert.Len(t, clients, 1)
	assert.Equal(t, "Anna Nowak", clients[0].Name)
}
