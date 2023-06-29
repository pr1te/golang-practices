package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pr1te/announcify-api/pkg/controllers"
	"go.uber.org/dig"
)

func InitWorkspaceRouteV1(router fiber.Router, container *dig.Container) {
	router.Get("/workspaces", func(c *fiber.Ctx) error {
		err := container.Invoke(func(controller *controllers.WorkspaceController) {
			controller.GetWorkspace(c)
		})

		return err
	})

	router.Post("/workspaces", func(c *fiber.Ctx) error {
		err := container.Invoke(func(controller *controllers.WorkspaceController) {
			controller.CreateWorkspace(c)
		})

		return err
	})
}
