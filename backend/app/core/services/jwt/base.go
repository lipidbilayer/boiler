package jwt

import (
	"strconv"

	"github.com/golang-jwt/jwt/v4"
	service "github.com/lipidbilayer/boiler/app/core/services"
	"github.com/lipidbilayer/boiler/lib/apperror"
)

type JWTAuth struct {
}

func New(config service.ConfigService, file service.EmbeddedFile) *JWTAuth {
	Init(config, file)
	return &JWTAuth{}
}

func (a *JWTAuth) GenerateRefreshToken(id int64) (string, error) {
	return generateToken(id, "", refreshExpiration)
}

func (a *JWTAuth) GenerateAccessToken(id int64, username string) (string, error) {
	return generateToken(id, username, expiration)
}

func (a *JWTAuth) ParseToken(tokenString string) (interface{}, error) {
	token, _, err := parseToken(tokenString)
	return token, err
}

func (a *JWTAuth) CheckToken(token string) (interface{}, error) {
	return CheckToken(token)
}

func (a *JWTAuth) GetAuthToken(authHeader string) string {
	return GetAuthToken(authHeader)
}

func (a *JWTAuth) GetAuthUserID(tokenString string) (int64, error) {
	_, claim, err := parseToken(tokenString)
	if err != nil {
		return 0, a.errorDatabase(err)
	}

	return strconv.ParseInt(claim.ID, 10, 64)
}

func (a *JWTAuth) errorDatabase(err error) error {
	switch err.(type) {
	case *jwt.ValidationError:
		return apperror.NewError(err, "token tidak di temukan", apperror.AuthError)
	}
	return err
}
