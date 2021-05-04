package services

import (
	"github.com/jailtonjunior94/bookings/api/application/dtos/requests"
	"github.com/jailtonjunior94/bookings/api/application/dtos/responses"
	"github.com/jailtonjunior94/bookings/api/domain/interfaces"
	"github.com/jailtonjunior94/bookings/api/infrastructure/adapters"
)

type AuthService struct {
	UserRepository interfaces.IUserRepository
	HashAdapter    adapters.IHashAdapter
	JwtAdapter     adapters.IJwtAdapter
}

func NewAuthService(r interfaces.IUserRepository, h adapters.IHashAdapter, j adapters.IJwtAdapter) interfaces.IAuthService {
	return &AuthService{UserRepository: r, HashAdapter: h, JwtAdapter: j}
}

func (a *AuthService) Authenticate(request *requests.Login) *responses.HttpResponse {
	user, err := a.UserRepository.GetByEmail(request.Email)
	if err != nil {
		return responses.ServerError()
	}

	if user == nil {
		return responses.BadRequest("Usuário e/ou senha inválidos")
	}

	if isValid := a.HashAdapter.CheckHash(user.Password, request.Password); !isValid {
		return responses.BadRequest("Usuário e/ou senha inválidos")
	}

	token, err := a.JwtAdapter.GenerateTokenJWT(user.ID.Hex(), user.Email)
	if err != nil {
		return responses.ServerError()
	}

	return responses.Ok(responses.NewAuthResponse(token))
}
