package middleware

import (
	"os"
	"time"

	"github.com/evasbr/osam-backend/app/errors"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	IdUser string `json:"id_user"`
	jwt.RegisteredClaims
}

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

func AuthMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		tokenStr := c.Cookies("token")

		if tokenStr == "" {
			return errors.HttpError{
				Messages:   []string{"Missing auth token"},
				StatusCode: fiber.StatusUnauthorized,
			}
		}

		token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})
		if err != nil || !token.Valid {
			return errors.HttpError{
				Messages:   []string{"Invalid or expired token"},
				StatusCode: fiber.StatusUnauthorized,
			}
		}

		claims, ok := token.Claims.(*Claims)
		if !ok {
			return errors.HttpError{
				Messages:   []string{"Invalid token claims"},
				StatusCode: fiber.StatusUnauthorized,
			}
		}

		if claims.ExpiresAt.Time.Before(time.Now()) {
			return errors.HttpError{
				Messages:   []string{"Invalid token"},
				StatusCode: fiber.StatusUnauthorized,
			}
		}

		c.Locals("idUser", claims.IdUser)

		return c.Next()
	}
}
