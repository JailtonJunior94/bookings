package interfaces

import (
	"github.com/jailtonjunior94/bookings/api/application/dtos/requests"
	"github.com/jailtonjunior94/bookings/api/application/dtos/responses"
)

type IUserService interface {
	CreateUser(request *requests.Register) *responses.HttpResponse
}
