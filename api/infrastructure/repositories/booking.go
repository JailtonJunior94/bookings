package repositories

import (
	"context"
	"errors"
	"log"

	"github.com/jailtonjunior94/bookings/api/domain/entities"
	"github.com/jailtonjunior94/bookings/api/domain/interfaces"
	"github.com/jailtonjunior94/bookings/api/infrastructure/database"
	"github.com/jailtonjunior94/bookings/api/infrastructure/environments"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (r *BookingRepository) Get(userId string) (bookings []entities.Booking, err error) {
	collection, err := r.Client.GetCollection(environments.BookingDatabase, environments.BookingCollection)
	if err != nil {
		log.Println(err)
		return nil, errors.New("Ocorreu um erro inesperado!")
	}

	objectID, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		log.Println(err)
		return nil, errors.New("Erro ao converter ID")
	}

	cursor, err := collection.Find(context.Background(), bson.M{"user._id": objectID})
	if err != nil {
		log.Println(err)
		return nil, err
	}

	if err = cursor.All(context.Background(), &bookings); err != nil {
		log.Println(err)
		return nil, err
	}

	return bookings, nil
}
