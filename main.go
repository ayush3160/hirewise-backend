package main

import (
	"os"

	"github.com/ayush3160/hirewise-backend/config"
	"github.com/ayush3160/hirewise-backend/database"
	"github.com/ayush3160/hirewise-backend/server"
)

func main() {
	_, err := os.Stat("config.yaml")
	if os.IsNotExist(err) {
		err := config.CreateDefaultConfig()
		if err != nil {
			panic("failed to create default config file")
		}
	}

	conf := config.GetConfig()
	db := database.NewPostgresDatabase(conf)
	server.NewEchoServer(conf, db).Start()
}
