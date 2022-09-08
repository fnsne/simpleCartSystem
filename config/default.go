//go:build default

package config

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	_ "gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewUserDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("default.db"), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	return db
}

var GinMode = gin.DebugMode

func DropTestTable(table interface{}) {
	fmt.Println("cannot using this method on product db")
}
