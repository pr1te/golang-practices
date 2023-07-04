package routes

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/dig"
)

func InitGreetingRouteV1(router fiber.Router, container *dig.Container) {
	router.Get("/greeting", func(c *fiber.Ctx) error {
		error := c.SendString("Hello Greeting from NOWHERE")

		return error
	})
}
