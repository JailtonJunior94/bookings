package interfaces

import (
	"github.com/jailtonjunior94/bookings/api/application/dtos/requests"
	"github.com/jailtonjunior94/bookings/api/application/dtos/responses"
)

type IAuthService interface {
	Authenticate(request *requests.Login) *responses.HttpResponse
}
