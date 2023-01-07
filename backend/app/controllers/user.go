package controllers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/lipidbilayer/boiler/app/core"
	"github.com/lipidbilayer/boiler/app/models"
)

type User struct {
	BaseController
}

func NewUser(service *core.AppServices) *User {
	user := &User{}
	user.Services = service
	return user
}

func (u *User) Index(c *fiber.Ctx) error {
	items, err := u.Services.Database.IndexUser(c.Context())
	if err != nil {
		return err
	}
	return c.Status(http.StatusOK).JSON(items)
}

func (u *User) Show(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return err
	}

	item := &models.User{ID: int64(id)}
	err = u.Services.Database.ShowUser(c.Context(), item)
	if err != nil {
		return err
	}
	return c.Status(http.StatusOK).JSON(item)
}

func (u *User) Create(c *fiber.Ctx) error {
	item, err := u.parseUserFromRequest(c)
	if err != nil {
		return err
	}

	item.ID = 0
	err = u.Services.Database.CreateUser(c.Context(), item)
	if err != nil {
		return err
	}
	return c.Status(http.StatusOK).JSON(item)
}

func (u *User) Update(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return err
	}

	item, err := u.parseUserFromRequest(c)
	if err != nil {
		return err
	}
	item.ID = int64(id)
	err = u.Services.Database.UpdateUser(c.Context(), item)
	if err != nil {
		return err
	}
	return c.Status(http.StatusOK).JSON(item)
}

func (u *User) Delete(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return err
	}

	item := &models.User{ID: int64(id)}
	err = u.Services.Database.DeleteUser(c.Context(), item)
	if err != nil {
		return err
	}
	return c.Status(http.StatusOK).JSON(item)
}

func (u *User) parseUserFromRequest(c *fiber.Ctx) (*models.User, error) {
	user := &models.User{}
	err := c.BodyParser(user)
	return user, err
}
