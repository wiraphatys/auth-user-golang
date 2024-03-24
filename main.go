package main

import (
	"banky/config"
	"banky/database"
	"banky/server"
)

func main() {
	cfg := config.GetConfig()
	db := database.NewPostgresDatabase(cfg)

	server.NewFiberServer(cfg, db.GetDb()).Start()
}
