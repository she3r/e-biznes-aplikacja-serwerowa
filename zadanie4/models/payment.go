package models

import (
	"gorm.io/gorm"
)

type Payment struct {
	gorm.Model
	Date     string `json:"datetime"`
	Value    float64
	ClientID uint   `gorm:"not null;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Client   Client `gorm:"foreignKey:ClientID" json:"-"`
}
