package controllers

import (
	"net/http"
	"time"

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

type AuthRefresh struct {
	RefreshToken string `json:"refresh_token"`
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

	token, err := a.Services.AuthService.GenerateAccessToken(user.ID, user.Username)
	if err != nil {
		return err
	}

	refreshToken, err := a.Services.AuthService.GenerateRefreshToken(user.ID)
	if err != nil {
		return err
	}
	now := time.Now()
	user.LastLoginAt = &now

	err = a.Services.Database.UpdateUser(c.Context(), user)
	if err != nil {
		return err
	}

	return c.Status(http.StatusOK).JSON(map[string]string{
		"token":         token,
		"refresh_token": refreshToken,
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

func (a Auth) RefreshToken(c *fiber.Ctx) error {
	params := &AuthRefresh{}
	err := c.BodyParser(params)
	if err != nil {
		return err
	}

	userID, err := a.Services.AuthService.GetAuthUserID(params.RefreshToken)
	if err != nil {
		return err
	}

	user := &models.User{ID: userID}
	err = a.Services.Database.ShowUser(c.Context(), user)
	if err != nil {
		return err
	}

	token, err := a.Services.AuthService.GenerateAccessToken(user.ID, user.Username)
	if err != nil {
		return err
	}

	refreshToken, err := a.Services.AuthService.GenerateRefreshToken(user.ID)
	if err != nil {
		return err
	}

	return c.Status(http.StatusOK).JSON(map[string]string{
		"token":         token,
		"refresh_token": refreshToken,
	})
}

func (a Auth) parseUserInfo(c *fiber.Ctx) (*models.User, error) {
	user := &models.User{}
	err := c.BodyParser(user)
	return user, err
}
