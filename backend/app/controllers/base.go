package controllers

import (
	"github.com/lipidbilayer/boiler/app/core"
	"github.com/lipidbilayer/boiler/app/models"
)

type BaseController struct {
	Services          *core.AppServices
	AuthenticatedUser *models.User
}
