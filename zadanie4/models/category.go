package models

import (
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Name     string
	Products []Product `gorm:"constraint:OnDelete:CASCADE;"`
}
