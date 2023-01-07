package controllers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/lipidbilayer/boiler/app/core"
	"github.com/lipidbilayer/boiler/app/core/services/jwt"
	"github.com/lipidbilayer/boiler/app/models"
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
	user, err := a.parseUserInfo(c)
	if err != nil {
		return err
	}

	err = a.Services.Database.GetUserWithPassword(c.Context(), user)
	if err != nil {
		return err
	}

	token, err := jwt.GenerateToken(user.ID, user.Username)
	if err != nil {
		return err
	}

	return c.Status(http.StatusOK).JSON(map[string]string{
		"token": token,
	})
}

func (a Auth) Profile(c *fiber.Ctx) error {
	tokenString, err := jwt.TokenExtractor(c)
	if err != nil {
		return err
	}
	userID, err := a.Services.AuthService.GetAuthUserID(tokenString)
	if err != nil {
		return err
	}
	user := &models.User{ID: userID}
	err = a.Services.Database.ShowUser(c.Context(), user)
	if err != nil {
		return err
	}

	return c.Status(http.StatusOK).JSON(user)
}

func (a Auth) parseUserInfo(c *fiber.Ctx) (*models.User, error) {
	user := &models.User{}
	err := c.BodyParser(user)
	return user, err
}
