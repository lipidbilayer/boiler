package main

import (
	"log"

	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/lipidbilayer/boiler/app/core"
	"github.com/lipidbilayer/boiler/lib/apperror"
	"github.com/lipidbilayer/boiler/routes"
)

func main() {

	// start service
	core.Init()
	defer core.Stop()

	// start web
	app := fiber.New(fiber.Config{
		ErrorHandler: apperror.AppErrorHandler,
	})

	//middleware
	app.Use(recover.New())
	app.Use(logger.New())

	//routes
	routes.Ping(app)
	routes.Auth(app, core.Services)
	routes.Api(app, core.Services)

	//start fiber
	log.Fatal(app.Listen(":8000"))
}
