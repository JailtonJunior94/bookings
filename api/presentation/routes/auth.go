package routes

import (
	"github.com/jailtonjunior94/bookings/api/infrastructure/ioc"

	"github.com/gofiber/fiber/v2"
)

func AddAuthRouter(router fiber.Router) {
	router.Post("/token", ioc.AuthController.Authenticate)
}
