package repositories

import (
	"github.com/jailtonjunior94/bookings/api/domain/entities"
	"github.com/jailtonjunior94/bookings/api/domain/interfaces"
	"github.com/jailtonjunior94/bookings/api/infrastructure/database"
)

type BookingRepository struct {
	Client database.IMongoConnection
}

func NewBookingRepository(client database.IMongoConnection) interfaces.IBookingRepository {
	return &BookingRepository{Client: client}
}

func (r *BookingRepository) Add(b *entities.Booking) (booking *entities.Booking, err error) {
	return nil, nil
}

func (r *BookingRepository) Get() (bookings []entities.Booking, err error) {
	return nil, nil
}
