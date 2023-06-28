package router

import "github.com/gofiber/fiber/v2"

func InitGreetingRoute(router fiber.Router) {
	v1 := router.Group("/v1")

	v1.Get("/greeting", func(c *fiber.Ctx) error {
		error := c.SendString("Hello Greeting from NOWHERE")

		return error
	})
}
