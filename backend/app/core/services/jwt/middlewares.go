package jwt

import (
	"github.com/gofiber/fiber/v2"
	service "github.com/lipidbilayer/boiler/app/core/services"
	"github.com/lipidbilayer/boiler/app/models"
)

var extractorList []Extractor = []Extractor{TokenHeaderExtractor, TokenQueryExtractor, TokenParamExtractor, TokenCookieExtractor}

type Extractor func(*fiber.Ctx) (string, error)

func NewJWTMiddleware(auth service.AuthService, database service.DatabaseService) fiber.Handler {

	return func(c *fiber.Ctx) error {
		tokenString, err := TokenExtractor(c)
		if err != nil {
			return err
		}
		// need filter to bypass auth controller
		claim, err := auth.CheckToken(tokenString)
		if err != nil {
			return err
		}

		userID, err := auth.GetAuthUserID(tokenString)
		if err != nil {
			return err
		}

		user := &models.User{ID: userID}
		err = database.ShowUser(c.Context(), user)
		if err != nil {
			return err
			// throw user not found or unexpected error
		}
		c.Locals(TOKEN_CLAIMS_KEY, claim)
		c.Locals(AUTHENTICATED_USER_KEY, user)
		return c.Next()
	}
}

func TokenExtractor(c *fiber.Ctx) (string, error) {
	var err error
	token := ""
	for _, value := range extractorList {
		token, err = value(c)
		if token != "" {
			return token, nil
		}
	}
	return "", err
}

func TokenHeaderExtractor(c *fiber.Ctx) (string, error) {
	bearer := c.Get("Authorization")
	return GetAuthToken(bearer), nil
}

func TokenQueryExtractor(c *fiber.Ctx) (string, error) {
	token := c.Query("auth_token", "")
	return token, nil
}

func TokenParamExtractor(c *fiber.Ctx) (string, error) {
	token := c.Params("auth_token", "")
	return token, nil
}

func TokenCookieExtractor(c *fiber.Ctx) (string, error) {
	token := c.Cookies("auth_token", "")
	return token, nil
}
