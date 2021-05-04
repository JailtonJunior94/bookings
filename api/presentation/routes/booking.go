package routes

import (
	"github.com/jailtonjunior94/bookings/api/infrastructure/ioc"

	"github.com/gofiber/fiber/v2"
)

func AddBookingRouter(router fiber.Router) {
	router.Get("/bookings", ioc.BookingController.Bookings)
	router.Post("/bookings", ioc.BookingController.CreateBooking)
}
