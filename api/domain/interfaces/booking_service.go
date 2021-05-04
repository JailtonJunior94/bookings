package interfaces

import (
	"github.com/jailtonjunior94/bookings/api/application/dtos/requests"
	"github.com/jailtonjunior94/bookings/api/application/dtos/responses"
)

type IBookingService interface {
	Bookings(userId string) *responses.HttpResponse
	CreateBooking(userId string, request *requests.Booking) *responses.HttpResponse
}
