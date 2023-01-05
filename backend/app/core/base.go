package core

import service "github.com/lipidbilayer/boiler/app/core/services"

type AppServices struct {
	Database    service.DatabaseService
	AuthService service.AuthService
	Config      service.ConfigService
}

var Services *AppServices

func Init() {

}

func Stop() {

}
