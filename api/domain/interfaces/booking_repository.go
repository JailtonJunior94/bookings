package interfaces

import "github.com/jailtonjunior94/bookings/api/domain/entities"

type IBookingRepository interface {
	Add(b *entities.Booking) (booking *entities.Booking, err error)
	Get(userId string) (bookings []entities.Booking, err error)
}
