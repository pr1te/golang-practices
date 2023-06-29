package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pr1te/announcify-api/pkg/routes/v1"
	"go.uber.org/dig"
)

func InitRouter(app *fiber.App, container *dig.Container) {
	// register v1 api
	v1 := app.Group("/api/v1")

	routes.InitGreetingRouteV1(v1, container)
	routes.InitWorkspaceRouteV1(v1, container)
}
