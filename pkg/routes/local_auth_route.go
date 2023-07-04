package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pr1te/announcify-api/pkg/controllers"
	"go.uber.org/dig"
)

func InitLocalAuthRoute(router fiber.Router, container *dig.Container) {
	router.Post("/local", func(c *fiber.Ctx) error {
		err := container.Invoke(func(controller *controllers.LocalAuthController) {
			controller.CreateAccount(c)
		})

		return err
	})
}
