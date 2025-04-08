package models

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name       string
	Price      float64
	CategoryID uint     `gorm:"not null;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Category   Category `gorm:"foreignKey:CategoryID" json:"-"`
	Baskets    []Basket `gorm:"many2many:basket_products;"`
}
