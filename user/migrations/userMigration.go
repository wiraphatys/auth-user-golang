package main

import (
	"banky/config"
	"banky/database"
	"banky/user/entities"
)

func main() {
	cfg := config.GetConfig()
	db := database.NewPostgresDatabase(cfg)

	// migrate user schema
	db.GetDb().AutoMigrate(&entities.User{})
	database.CreateUserIDTrigger(db.GetDb())
}
