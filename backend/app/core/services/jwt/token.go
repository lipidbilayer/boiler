package jwt

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type AuthClaim struct {
	LocationID int64 `json:"location_id"`
	jwt.RegisteredClaims
}

func generateToken(id int64, username string, expiration time.Duration) (string, error) {
	authClaim := AuthClaim{}
	authClaim.RegisteredClaims = jwt.RegisteredClaims{
		Subject:   username,
		ID:        fmt.Sprint(id),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(expiration * time.Minute)),
		NotBefore: jwt.NewNumericDate(time.Now()),
	}

	if isIssuerExists {
		authClaim.Issuer = issuer
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, authClaim)

	tokenString, err := token.SignedString(privateKey)
	return tokenString, err
}

func parseToken(tokenString string) (*jwt.Token, *AuthClaim, error) {
	token, err := jwt.ParseWithClaims(tokenString, &AuthClaim{}, func(token *jwt.Token) (interface{}, error) {
		return publicKey, nil
	})

	if err != nil {
		return nil, nil, err
	}

	if authClaim, ok := token.Claims.(*AuthClaim); ok && token.Valid {
		return token, authClaim, nil
	}
	return nil, nil, jwt.NewValidationError("Claim not valid", 100)
}
