package main

import (
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/storage/postgres/v2"
	"github.com/pr1te/left-it-api/pkg/config"
)

func NewSession(conf *config.Configuration) *session.Store {
	port, _ := strconv.Atoi(conf.Database.Port)

	postgresStorage := postgres.New(postgres.Config{
		Host:     conf.Database.Host,
		Port:     port,
		Username: conf.Database.Username,
		Password: conf.Database.Password,
		SSLMode:  conf.Database.Ssl,
		Database: conf.Database.Name,
		Table:    "_sessions",
	})

	storage := session.New(session.Config{
		Storage:        postgresStorage,
		CookieHTTPOnly: true,
		CookieSecure:   true,
		Expiration:     24 * time.Hour,
	})

	return storage
}
