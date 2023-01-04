package main

import (
	"log"

	"github.com/gofiber/fiber/v2/middleware/logger"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/lipidbilayer/boiler/routes"
)

func main() {
	app := fiber.New()

	//middleware
	app.Use(recover.New())
	app.Use(logger.New())

	//routes
	routes.Ping(app)

	//start fiber
	log.Fatal(app.Listen(":8000"))
}
