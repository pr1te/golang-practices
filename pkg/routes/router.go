package routes

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/dig"
)

func InitRouter(app *fiber.App, container *dig.Container) {
	// register auth api
	auth := app.Group("/auth")

	InitLocalAuthRoute(auth, container)

	// register v1 api
	v1 := app.Group("/api/v1")

	InitGreetingRouteV1(v1, container)
	InitWorkspaceRouteV1(v1, container)
}
