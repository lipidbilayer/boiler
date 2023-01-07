package service

import (
	"context"
	"net/http"

	"github.com/lipidbilayer/boiler/app/models"
)

type DatabaseService interface {
	Stop()
	GetUserWithPassword(context.Context, *models.User) error
	IndexUser(context.Context) ([]*models.User, error)
	ShowUser(context.Context, *models.User) error
	CreateUser(context.Context, *models.User) error
	UpdateUser(context.Context, *models.User) error
	DeleteUser(context.Context, *models.User) error
}

type ConfigService interface {
	GetDebugMode() bool
	GetDatabaseURL() string
	GetJWTRealmName() string
	GetJWTExpiration() int64
	GetJWTIssuerName() string
	GetJWTPublicKeyPath() string
	GetJWTPrivateKeyPath() string
	GetSentryDSN() string
}

type AuthService interface {
	ParseToken(tokenString string) (interface{}, error)
	CheckToken(token string) (interface{}, error)
	GetAuthToken(authHeader string) string
	GetAuthUserID(token string) (int64, error)
}

type EmbeddedFile interface {
	GetEmbeddedFile() http.FileSystem
}
