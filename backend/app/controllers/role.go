package controllers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/lipidbilayer/boiler/app/core"
)

type Role struct {
	BaseController
}

func NewRole(service *core.AppServices) *Role {
	role := &Role{}
	role.Services = service
	return role
}

func (u *Role) Index(c *fiber.Ctx) error {
	items, err := u.Services.Database.IndexRole(c.Context())
	if err != nil {
		return err
	}
	return c.Status(http.StatusConflict).JSON(items)
}
