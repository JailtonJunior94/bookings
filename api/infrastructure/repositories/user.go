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

type UserRepository struct {
	Client database.IMongoConnection
}

func NewUserRepository(client database.IMongoConnection) interfaces.IUserRepository {
	return &UserRepository{Client: client}
}

func (r *UserRepository) Add(u *entities.User) (user *entities.User, err error) {
	collection, err := r.Client.GetCollection(environments.BookingDatabase, environments.UsersCollection)
	if err != nil {
		log.Println(err)
		return nil, errors.New("Ocorreu um erro inesperado!")
	}

	_, err = collection.InsertOne(context.Background(), &u)
	if err != nil {
		log.Println(err)
		return nil, errors.New("Ocorreu um erro inesperado!")
	}

	return u, nil
}

func (r *UserRepository) GetByEmail(email string) (user *entities.User, err error) {
	collection, err := r.Client.GetCollection(environments.BookingDatabase, environments.UsersCollection)
	if err != nil {
		log.Println(err)
		return nil, errors.New("Ocorreu um erro inesperado!")
	}

	if err := collection.FindOne(context.Background(), bson.M{"email": email}).Decode(&user); err != nil {
		log.Println(err)
		return nil, nil
	}

	return user, nil
}

func (r *UserRepository) GetById(id string) (user *entities.User, err error) {
	collection, err := r.Client.GetCollection(environments.BookingDatabase, environments.UsersCollection)
	if err != nil {
		log.Println(err)
		return nil, errors.New("Ocorreu um erro inesperado!")
	}

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println(err)
		return nil, errors.New("Erro ao converter ID")
	}

	if err := collection.FindOne(context.Background(), bson.M{"_id": objectID}).Decode(&user); err != nil {
		log.Println(err)
		return nil, nil
	}

	return user, nil
}
