package interfaces

import "github.com/jailtonjunior94/bookings/api/domain/entities"

type IUserRepository interface {
	Add(p *entities.User) (user *entities.User, err error)
	GetByEmail(email string) (user *entities.User, err error)
	GetById(id string) (user *entities.User, err error)
}
