package controllers

import (
	"github.com/jailtonjunior94/bookings/api/application/dtos/requests"
	"github.com/jailtonjunior94/bookings/api/domain/interfaces"

	"github.com/gofiber/fiber/v2"
)

type AuthController struct {
	Service interfaces.IAuthService
}

func NewAuthController(u interfaces.IAuthService) *AuthController {
	return &AuthController{Service: u}
}

func (u *AuthController) Authenticate(c *fiber.Ctx) error {
	request := new(requests.Login)
	if err := c.BodyParser(request); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"error": "Unprocessable Entity"})
	}

	if err := request.IsValid(); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	response := u.Service.Authenticate(request)
	return c.Status(response.StatusCode).JSON(response.Data)
}
