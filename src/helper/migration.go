package helper

import (
	"golang-web/src/config"
	"golang-web/src/models"
)

func Migration() {
	config.DB.AutoMigrate(&models.Product{})
	config.DB.AutoMigrate(&models.User{})
}
