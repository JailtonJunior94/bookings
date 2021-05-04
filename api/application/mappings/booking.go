package mappings

import (
	"github.com/jailtonjunior94/bookings/api/application/dtos/requests"
	"github.com/jailtonjunior94/bookings/api/application/dtos/responses"
	"github.com/jailtonjunior94/bookings/api/domain/entities"
)

func ToBookingEntity(r *requests.Booking, u *entities.User) (e *entities.Booking) {
	entity := new(entities.Booking)
	entity.NewBooking(r.Date, *u)

	return entity
}

func ToBookingResponse(e *entities.Booking) (r *responses.BookingResponse) {
	return &responses.BookingResponse{
		ID:   e.ID.Hex(),
		Date: e.Date,
		User: ToUserResponse(&e.User),
	}
}
