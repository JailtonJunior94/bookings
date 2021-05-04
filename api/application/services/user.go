package services

import (
	"github.com/jailtonjunior94/bookings/api/application/dtos/requests"
	"github.com/jailtonjunior94/bookings/api/application/dtos/responses"
	"github.com/jailtonjunior94/bookings/api/application/mappings"
	"github.com/jailtonjunior94/bookings/api/domain/interfaces"
	"github.com/jailtonjunior94/bookings/api/infrastructure/adapters"
)

type UserService struct {
	UserRepository interfaces.IUserRepository
	HashAdapter    adapters.IHashAdapter
}

func NewUserService(r interfaces.IUserRepository, h adapters.IHashAdapter) interfaces.IUserService {
	return &UserService{UserRepository: r, HashAdapter: h}
}

func (u *UserService) CreateUser(request *requests.Register) *responses.HttpResponse {
	passwordCript, err := u.HashAdapter.GenerateHash(request.Password)
	if err != nil {
		return responses.BadRequest("Erro ao criar usuário")
	}

	newUser := mappings.ToUserEntity(request, passwordCript)
	user, err := u.UserRepository.Add(newUser)
	if err != nil {
		return responses.ServerError()
	}

	result := mappings.ToUserResponse(user)
	return responses.Created(result)
}
