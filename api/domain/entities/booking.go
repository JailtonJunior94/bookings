package entities

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Booking struct {
	ID   primitive.ObjectID `bson:"_id"`
	Date time.Time          `bson:"date"`
	User User               `bson:"user"`
}

func (b *Booking) NewBooking(date time.Time, user User) {
	b.ID = primitive.NewObjectID()
	b.Date = date
	b.User = user
}
