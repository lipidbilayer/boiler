package service

import "net/http"

type DatabaseService interface {
	Stop()
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
