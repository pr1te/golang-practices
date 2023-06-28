package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/gofiber/fiber/v2"
	"github.com/pr1te/announcify-api/pkg/common/config"
	"github.com/pr1te/announcify-api/pkg/common/database"
	"github.com/pr1te/announcify-api/pkg/common/logger"
	"github.com/pr1te/announcify-api/pkg/router"
	"go.uber.org/dig"
)

var Version = "unset"

func main() {
	// create ioc
	container := dig.New()

	// load config
	config, loadConfigError := config.Load()

	if loadConfigError != nil {
		log.Panicln(loadConfigError)
	}

	// create go fiber app
	app := fiber.New()

	// create logger
	logger, closeLogger, createLoggerError := logger.New(config.Logger.Path, config.Logger.Level)

	if createLoggerError != nil {
		log.Panicln(createLoggerError)
	}

	logger.Debugf("----- DEBUG MODE ENABLED -----")

	// graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	go func() {
		_ = <-c

		logger.Debugln("gracefully shutting down")

		_ = app.Shutdown()
	}()

	// establish database connection
	db := database.New(&database.ConnectionOptions{
		Name:     config.Database.Name,
		Host:     config.Database.Host,
		Port:     config.Database.Port,
		Username: config.Database.Username,
		Password: config.Database.Password,
		Ssl:      config.Database.Ssl,
	}, logger)

	if err := db.Connect(); err != nil {
		logger.Panicln(err)
	}

	logger.Infof("establish connection database connection to '%s:%s'", config.Database.Host, config.Database.Port)

	// register the dependencies to ioc
	container.Provide(db)
	container.Provide(config)
	container.Provide(logger)

	// register routes
	router.Init(app)

	port := fmt.Sprintf(":%s", config.Http.Port)

	logger.Infof("app version - %s", Version)
	logger.Infof("listening on port - %s", port)
	logger.Debugf("environment variables - %+v", config)

	if err := app.Listen(port); err != nil {
		logger.Panicln(err)
	}

	// cleanup tasks goes here
	logger.Debugln("running cleanup tasks")

	defer closeLogger()
	defer db.Disconnect()
}
