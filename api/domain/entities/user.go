package entities

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID       primitive.ObjectID `bson:"_id"`
	Name     string             `bson:"name"`
	Email    string             `bson:"email"`
	Password string             `bson:"password"`
}

func (u *User) NewUser(name, email, password string) {
	u.ID = primitive.NewObjectID()
	u.Name = name
	u.Email = email
	u.Password = password
}
