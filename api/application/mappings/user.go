package mappings

import (
	"github.com/jailtonjunior94/bookings/api/application/dtos/requests"
	"github.com/jailtonjunior94/bookings/api/application/dtos/responses"
	"github.com/jailtonjunior94/bookings/api/domain/entities"
)

func ToUserEntity(r *requests.Register, password string) (e *entities.User) {
	entity := new(entities.User)
	entity.NewUser(r.Name, r.Email, password)

	return entity
}

func ToUserResponse(e *entities.User) (r *responses.UserResponse) {
	return &responses.UserResponse{
		ID:    e.ID.Hex(),
		Name:  e.Name,
		Email: e.Email,
	}
}
