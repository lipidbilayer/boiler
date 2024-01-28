package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lipidbilayer/boiler/app/controllers"
	"github.com/lipidbilayer/boiler/app/core"
	"github.com/lipidbilayer/boiler/app/core/services/jwt"
)

func Api(routes fiber.Router, service *core.AppServices) {
	user := controllers.NewUser(service)
	userRoute := routes.Group("/api/user", jwt.NewJWTMiddleware(core.Services.AuthService, core.Services.Database))

	userRoute.Get("/", user.Index)
	userRoute.Get("/:id", user.Show)
	userRoute.Post("/", user.Create)
	userRoute.Patch("/:id", user.Update)
	userRoute.Delete("/:id", user.Delete)

	role := controllers.NewRole(service)
	roleRoute := routes.Group("/api/role", jwt.NewJWTMiddleware(core.Services.AuthService, core.Services.Database))

	roleRoute.Get("/", role.Index)
}
