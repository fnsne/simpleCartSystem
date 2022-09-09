package initial

import (
	"github.com/gin-gonic/gin"
	"shopline-question/config"
	"shopline-question/model"
	"shopline-question/model/repository"
)

func Initial() {
	db := config.NewUserDB()
	repository.Initial(db)
	config.DB = db
	gin.SetMode(config.GinMode)
}

func MigrateDB() {
	db := config.NewUserDB()
	model.Migrate(db)
}
