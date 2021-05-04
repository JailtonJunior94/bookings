package middlewares

import (
	"github.com/gofiber/fiber/v2"

	jwtware "github.com/gofiber/jwt/v2"
	"github.com/jailtonjunior94/bookings/api/infrastructure/environments"
)

func Protected() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey:   []byte(environments.JwtSecret),
		ErrorHandler: jwtError,
	})
}

func jwtError(c *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "JWT ausente ou malformado"})
	}
	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Token inválido ou expirado"})
}
