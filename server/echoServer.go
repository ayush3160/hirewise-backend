package server

import (
	"fmt"

	"github.com/ayush3160/hirewise-backend/config"
	"github.com/ayush3160/hirewise-backend/database"

	"github.com/labstack/echo/v4"
)

type echoServer struct {
	app  *echo.Echo
	conf *config.Config
	db   database.Database
}

func NewEchoServer(conf *config.Config, db database.Database) *echoServer {
	return &echoServer{
		app:  echo.New(),
		conf: conf,
		db:   db,
	}
}

func (s *echoServer) Start() error {
	serverUrl := fmt.Sprintf(":%d", s.conf.Server.Port)

	if err := s.app.Start(serverUrl); err != nil {
		return err
	}
	return nil
}
