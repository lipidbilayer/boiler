package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lipidbilayer/boiler/app/core"
)

type Auth struct {
	BaseController
}

func NewAuth(service *core.AppServices) *Auth {
	auth := &Auth{}
	auth.Services = service
	return auth
}

func (a Auth) Login(c *fiber.Ctx) error {
	return nil
}
