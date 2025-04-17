package main

import (
	"hirewise-backend/config"
	"hirewise-backend/database"
	"hirewise-backend/server"
	"os"
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
