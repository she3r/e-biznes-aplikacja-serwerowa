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

func TestCreatePayment(t *testing.T) {
	setupTestDB()
	e := setupEcho()

	// Create a client first
	client := models.Client{Name: "Jan Kowalski", Address: "ul. Kwiatowa 12"}
	db.DB.Create(&client)

	payment := models.Payment{
		Date:     "2025-05-04T12:00:00Z",
		Value:    199.99,
		ClientID: client.ID,
	}
	body, _ := json.Marshal(payment)
	req := httptest.NewRequest(http.MethodPost, "/payments", bytes.NewBuffer(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err := CreatePayment(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, rec.Code)

	var created models.Payment
	_ = json.Unmarshal(rec.Body.Bytes(), &created)
	assert.Equal(t, 199.99, created.Value)
	assert.Equal(t, client.ID, created.ClientID)
}
