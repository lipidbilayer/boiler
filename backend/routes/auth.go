package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lipidbilayer/boiler/app/controllers"
	"github.com/lipidbilayer/boiler/app/core"
)

func Auth(routes fiber.Router, service *core.AppServices) {
	auth := controllers.NewAuth(service)
	ping := routes.Group("/auth")

	ping.Get("/login", auth.Login)
}
