package apperror

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

// New creates a new middleware handler
func AppErrorHandler(c *fiber.Ctx, err error) error {
	code := http.StatusInternalServerError

	var errorResult interface{}
	errorResult = map[string]interface{}{
		"err":     err.Error(),
		"code":    1000,
		"message": "Something went wrong",
	}

	if commonErr, ok := err.(CommonError); ok {
		// if commonErr.Code==apperror.NotFoundError
		errorResult = commonErr
		switch commonErr.Code {
		case NotFoundError:
			code = http.StatusNotFound
		case AuthError:
			code = http.StatusUnauthorized
		default:
			code = http.StatusInternalServerError
		}

	} else {

	}
	if err != nil {
		return c.Status(code).JSON(errorResult)
	}
	// Return err if exist, else move to next handler
	return nil
}
