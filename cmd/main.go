package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/gofiber/fiber/v2"
	"github.com/pr1te/announcify-api/pkg/config"
	"github.com/pr1te/announcify-api/pkg/controllers"
	"github.com/pr1te/announcify-api/pkg/database"
	"github.com/pr1te/announcify-api/pkg/exceptions"
	"github.com/pr1te/announcify-api/pkg/logger"
	"github.com/pr1te/announcify-api/pkg/middlewares"
	"github.com/pr1te/announcify-api/pkg/repositories"
	"github.com/pr1te/announcify-api/pkg/routes"
	"github.com/pr1te/announcify-api/pkg/services"
	"go.uber.org/dig"
	"go.uber.org/zap"
)

var Version = "unset"

func main() {
	// create ioc
	container := dig.New()

	// load config
	conf, loadConfigError := config.Load()

	if loadConfigError != nil {
		log.Panicln(loadConfigError)
	}

	// create go fiber app
	app := fiber.New(fiber.Config{
		ErrorHandler: errorHandler,
	})

	// create logger
	logger, closeLogger, createLoggerError := logger.New(conf.Logger.Path, conf.Logger.Level)

	if createLoggerError != nil {
		log.Panicln(createLoggerError)
	}

	logger.Debugf("----- DEBUG MODE ENABLED -----")

	// graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	go func() {
		<-c

		logger.Debugln("gracefully shutting down")

		app.Shutdown()
	}()

	// establish database connection
	db := database.New(&database.ConnectionOptions{
		Name:     conf.Database.Name,
		Host:     conf.Database.Host,
		Port:     conf.Database.Port,
		Username: conf.Database.Username,
		Password: conf.Database.Password,
		Ssl:      conf.Database.Ssl,
	}, logger)

	if err := db.Connect(); err != nil {
		logger.Panicln(err)
	}

	logger.Infof("establish connection database connection to '%s:%s'", conf.Database.Host, conf.Database.Port)

	// register the dependencies to ioc
	container.Provide(func() *config.Configuration { return conf })
	container.Provide(func() *zap.SugaredLogger { return logger })
	container.Provide(func() *database.Database { return db })

	providers := []interface{}{
		// repositories
		repositories.NewHelper,
		repositories.NewWorkspace,
		repositories.NewLocalUser,

		// services
		services.NewLocalAuth,
		services.NewWorkspace,

		// controllers
		controllers.NewLocalAuth,
		controllers.NewWorkspace,
	}

	for _, provider := range providers {
		container.Provide(provider)
	}

	// apply middlewares
	app.Use(middlewares.NewHttpLogger(middlewares.HttpLogConfig{Logger: logger}))

	// register routes
	routes.InitRouter(app, container)

	app.Use(func(c *fiber.Ctx) error {
		return exceptions.NewNotFoundException("not found resource")
	})

	port := fmt.Sprintf(":%s", conf.Http.Port)

	logger.Infof("app version - %s", Version)
	logger.Infof("listening on port - %s", port)
	logger.Debugf("environment variables - %+v", conf)

	if err := app.Listen(port); err != nil {
		logger.Panicln(err)
	}

	// cleanup tasks goes here
	logger.Debugln("running cleanup tasks")

	defer closeLogger()
	defer db.Disconnect()
}