package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pr1te/announcify-api/pkg/controllers"
	"github.com/pr1te/announcify-api/pkg/libs/validator"
	"go.uber.org/dig"
)

func InitLocalAuthRoute(router fiber.Router, container *dig.Container) {
	router.Post("/local", func(c *fiber.Ctx) error {
		err := container.Invoke(func(controller *controllers.LocalAuthController, validator *validator.Validator) error {
			if err := controller.CreateAccount(c, validator); err != nil {
				return err
			}

			return nil
		})

		return err
	})

	router.Post("/local/login", func(c *fiber.Ctx) error {
		err := container.Invoke(func(controller *controllers.LocalAuthController, validator *validator.Validator) error {
			if err := controller.Login(c, validator); err != nil {
				return err
			}

			return nil
		})

		return err
	})
}
