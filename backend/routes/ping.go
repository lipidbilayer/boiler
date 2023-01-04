package routes

import "github.com/gofiber/fiber/v2"

func Ping(routes fiber.Router) {
	ping := routes.Group("/ping")

	ping.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("pong")
	})
}
