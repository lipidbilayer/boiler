package jwt

import (
	"log"

	"github.com/golang-jwt/jwt/v4"
	"github.com/lipidbilayer/boiler/lib/apperror"
)

func GetAuthToken(authorization string) string {
	if len(authorization) > 7 { // char count "Bearer " ==> 7
		return authorization[7:]
	}

	return ""
}

func CheckToken(tokenString string) (interface{}, error) {
	token, _, err := ParseToken(tokenString)
	if err == nil && token.Valid && !IsInBlocklist(tokenString) {
		return token.Claims, nil
	} else {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				log.Print("That's not even a token")
				err = apperror.NewError(err, "Unknown Token", apperror.AuthError)
			} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
				log.Print("Timing is everything, Token is either expired or not active yet")
				err = apperror.NewError(err, "Token expired", apperror.AuthError)
			} else {
				log.Printf("Couldn't handle this token: %v", err)
				err = apperror.NewError(err, "Unknown error token", apperror.AuthError)
			}
		} else {
			log.Printf("Couldn't handle this token: %v", err)
		}
	}
	return nil, err
}
