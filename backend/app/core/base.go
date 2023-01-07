package core

import (
	service "github.com/lipidbilayer/boiler/app/core/services"
	dbgobun "github.com/lipidbilayer/boiler/app/core/services/db-gobun"
	"github.com/lipidbilayer/boiler/app/core/services/envconfig"
	"github.com/lipidbilayer/boiler/app/core/services/jwt"
	"github.com/lipidbilayer/boiler/app/core/services/statik"
	_ "github.com/lipidbilayer/boiler/lib/files"
)

type AppServices struct {
	File        service.EmbeddedFile
	Database    service.DatabaseService
	AuthService service.AuthService
	Config      service.ConfigService
}

var Services *AppServices

func Init() {
	statikFS := statik.New()
	config := envconfig.New()
	database := dbgobun.New(config)
	auth := jwt.New(config, statikFS)

	Services = &AppServices{File: statikFS, Config: config, Database: database, AuthService: auth}
}

func Stop() {
	Services.Database.Stop()
}
