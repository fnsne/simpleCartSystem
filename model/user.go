package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Cart  Cart  `gorm:"foreignKey:UserID"`
	Order Order `gorm:"foreignKey:UserID"`
}
