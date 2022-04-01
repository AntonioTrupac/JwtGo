package config

import "antonio.trupac/models"

func MigrateTable() {
	db := GetDB()

	db.AutoMigrate(&models.User{})
}
