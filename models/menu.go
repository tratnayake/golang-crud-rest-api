package models

import (
	"gorm.io/gorm"
)

type Menu struct {
	gorm.Model
	// ActiveFrom time.Time
	// ActiveTo   time.Time
	Products []Product `gorm:"many2many:menu_products"`
}
