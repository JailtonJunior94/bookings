package controllers

import (
	"github.com/jailtonjunior94/bookings/api/application/dtos/requests"
	"github.com/jailtonjunior94/bookings/api/domain/interfaces"
	"github.com/jailtonjunior94/bookings/api/infrastructure/adapters"

	"github.com/gofiber/fiber/v2"
)

type BookingController struct {
	Jwt     adapters.IJwtAdapter
	Service interfaces.IBookingService
}

func NewBookingController(u interfaces.IBookingService, j adapters.IJwtAdapter) *BookingController {
	return &BookingController{Service: u, Jwt: j}
}

func (u *BookingController) CreateBooking(c *fiber.Ctx) error {
	userId, err := u.Jwt.ExtractClaims(c.Get("Authorization"))
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Token inválido ou expirado"})
	}

	request := new(requests.Booking)
	if err := c.BodyParser(request); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"error": "Unprocessable Entity"})
	}

	if err := request.IsValid(); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	response := u.Service.CreateBooking(*userId, request)
	return c.Status(response.StatusCode).JSON(response.Data)
}

func (u *BookingController) Bookings(c *fiber.Ctx) error {
	userId, err := u.Jwt.ExtractClaims(c.Get("Authorization"))
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Token inválido ou expirado"})
	}

	response := u.Service.Bookings(*userId)
	return c.Status(response.StatusCode).JSON(response.Data)
}
