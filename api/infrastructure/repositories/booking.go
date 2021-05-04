package repositories

import (
	"context"
	"errors"
	"log"

	"github.com/jailtonjunior94/bookings/api/domain/entities"
	"github.com/jailtonjunior94/bookings/api/domain/interfaces"
	"github.com/jailtonjunior94/bookings/api/infrastructure/database"
	"github.com/jailtonjunior94/bookings/api/infrastructure/environments"
)

type BookingRepository struct {
	Client database.IMongoConnection
}

func NewBookingRepository(client database.IMongoConnection) interfaces.IBookingRepository {
	return &BookingRepository{Client: client}
}

func (r *BookingRepository) Add(b *entities.Booking) (booking *entities.Booking, err error) {
	collection, err := r.Client.GetCollection(environments.BookingDatabase, environments.BookingCollection)
	if err != nil {
		log.Println(err)
		return nil, errors.New("Ocorreu um erro inesperado!")
	}

	_, err = collection.InsertOne(context.Background(), &b)
	if err != nil {
		log.Println(err)
		return nil, errors.New("Ocorreu um erro inesperado!")
	}

	return b, nil
}

func (r *BookingRepository) Get() (bookings []entities.Booking, err error) {
	return nil, nil
}
