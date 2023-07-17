package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pr1te/left-it-api/pkg/authenticator"
	"github.com/pr1te/left-it-api/pkg/controllers"
	"go.uber.org/dig"
)

func InitSessionRouteAuth(router fiber.Router, container *dig.Container) {
	router.Delete(
		"/session",
		func(c *fiber.Ctx) error {
			err := container.Invoke(func(authenticator *authenticator.Authenticator) error {
				if _, err := authenticator.Strategies.Session.Authenticate(c); err != nil {
					return err
				}

				return nil
			})

			if err != nil {
				return err
			}

			return c.Next()
		},
		func(c *fiber.Ctx) error {
			err := container.Invoke(func(controller *controllers.SessionController) error {
				if err := controller.Destroy(c); err != nil {
					return err
				}

				c.Status(204)

				return nil
			})

			return err
		},
	)
}
