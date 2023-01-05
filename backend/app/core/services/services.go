package service

import (
	"net/http"

	"github.com/lipidbilayer/boiler/app/models"
)

type DatabaseService interface {
	Stop()
	GetUserWithPassword(*models.User) error
	ShowUser(*models.User) error
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
