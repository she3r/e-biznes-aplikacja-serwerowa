package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"zadanie4_project/models"
)

var DB *gorm.DB

func Init() {
	var err error
	DB, err = gorm.Open(sqlite.Open("zadanie4.db"), &gorm.Config{})
	if err != nil {
		panic("Nie udało się połączyć z bazą danych")
	}

	DB.AutoMigrate(&models.Product{}, &models.Category{}, &models.Basket{}, &models.Employee{}, &models.Client{}, &models.Payment{})
}
