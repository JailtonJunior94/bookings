package routes

import (
	"github.com/jailtonjunior94/bookings/api/infrastructure/ioc"

	"github.com/gofiber/fiber/v2"
)

func AddUserRouter(router fiber.Router) {
	router.Post("/users", ioc.UserController.Create)
}
