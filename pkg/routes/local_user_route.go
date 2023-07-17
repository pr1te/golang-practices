package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pr1te/left-it-api/pkg/authenticator"
	"github.com/pr1te/left-it-api/pkg/controllers"
	"github.com/pr1te/left-it-api/pkg/libs/validator"
	"go.uber.org/dig"
)

func InitLocalUserRouteAuth(router fiber.Router, container *dig.Container) {
	router.Post("/local", func(c *fiber.Ctx) error {
		err := container.Invoke(func(controller *controllers.LocalUserController, validator *validator.Validator) error {
			if err := controller.CreateAccount(c, validator); err != nil {
				return err
			}

			return nil
		})

		return err
	})

	router.Post(
		"/local/login",
		func(c *fiber.Ctx) error {
			return container.Invoke(func(authenticator *authenticator.Authenticator) error {
				if _, err := authenticator.Strategies.Local.Authenticate(c); err != nil {
					return err
				}

				return nil
			})
		},
	)
}
