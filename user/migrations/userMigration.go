package main

import (
	"banky/config"
	"banky/database"
	"banky/user/entities"
)

func main() {
	cfg := config.GetConfig()
	db := database.NewPostgresDatabase(cfg)
	db.GetDb().AutoMigrate(&entities.User{})
}
