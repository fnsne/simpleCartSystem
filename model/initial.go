package model

import (
	"fmt"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	err := db.AutoMigrate(
		&Cart{},
		&User{},
		&Product{},
		&OrderProduct{},
		&Order{},
	)
	if err != nil {
		fmt.Println("migrate error:", err.Error())
	} else {
		fmt.Println("No Migrate Error.")
	}
}
