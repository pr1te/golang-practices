package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pr1te/announcify-api/pkg/authenticator"
	"github.com/pr1te/announcify-api/pkg/controllers"
	"go.uber.org/dig"
)

func InitMeRouteV1(router fiber.Router, container *dig.Container) {
	router.Get(
		"/me",
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
			return container.Invoke(func(controller *controllers.MeController) {
				controller.GetMyProfile(c)
			})
		},
	)
}
