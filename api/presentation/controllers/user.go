package controllers

import (
	"github.com/jailtonjunior94/bookings/api/application/dtos/requests"
	"github.com/jailtonjunior94/bookings/api/domain/interfaces"

	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	Service interfaces.IUserService
}

func NewUserController(u interfaces.IUserService) *UserController {
	return &UserController{Service: u}
}

func (u *UserController) Create(c *fiber.Ctx) error {
	request := new(requests.Register)
	if err := c.BodyParser(request); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"error": "Unprocessable Entity"})
	}

	if err := request.IsValid(); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	response := u.Service.CreateUser(request)
	return c.Status(response.StatusCode).JSON(response.Data)
}
