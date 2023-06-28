package router

import (
	"github.com/gofiber/fiber/v2"
	router "github.com/pr1te/announcify-api/pkg/router/routes"
)

func Init(app *fiber.App) {
	api := app.Group("/api")

	router.InitGreetingRoute(api)
}
