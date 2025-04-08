package models

import (
	"gorm.io/gorm"
)

type Basket struct {
	gorm.Model
	Products []Product `gorm:"many2many:basket_products;"`
}
