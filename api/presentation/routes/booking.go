package routes

import (
	"github.com/jailtonjunior94/bookings/api/infrastructure/ioc"
	"github.com/jailtonjunior94/bookings/api/presentation/middlewares"

	"github.com/gofiber/fiber/v2"
)

func AddBookingRouter(router fiber.Router) {
	router.Get("/bookings", middlewares.Protected(), ioc.BookingController.Bookings)
	router.Post("/bookings", middlewares.Protected(), ioc.BookingController.CreateBooking)
}
