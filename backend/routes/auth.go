package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lipidbilayer/boiler/app/controllers"
	"github.com/lipidbilayer/boiler/app/core"
)

func Auth(routes fiber.Router, service *core.AppServices) {
	auth := controllers.NewAuth(service)
	authRoute := routes.Group("/auth")

	authRoute.Post("/login", auth.Login)
	authRoute.Post("/refresh/token", auth.RefreshToken)
	authRoute.Get("/profile", auth.Profile)
}
