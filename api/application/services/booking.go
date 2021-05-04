package services

import (
	"encoding/json"

	"github.com/jailtonjunior94/bookings/api/application/dtos/requests"
	"github.com/jailtonjunior94/bookings/api/application/dtos/responses"
	"github.com/jailtonjunior94/bookings/api/application/mappings"
	"github.com/jailtonjunior94/bookings/api/domain/interfaces"
	"github.com/jailtonjunior94/bookings/api/infrastructure/environments"
)

type BookingService struct {
	RabbitMQ          interfaces.IRabbitMQ
	BookingRepository interfaces.IBookingRepository
	UserRepository    interfaces.IUserRepository
}

func NewBookingService(r interfaces.IBookingRepository, b interfaces.IRabbitMQ, u interfaces.IUserRepository) interfaces.IBookingService {
	return &BookingService{BookingRepository: r, RabbitMQ: b, UserRepository: u}
}

func (s *BookingService) Bookings() *responses.HttpResponse {
	return nil
}

func (s *BookingService) CreateBooking(userId string, request *requests.Booking) *responses.HttpResponse {
	user, err := s.UserRepository.GetById(userId)
	if err != nil {
		return responses.ServerError()
	}

	newBooking := mappings.ToBookingEntity(request, user)
	booking, err := s.BookingRepository.Add(newBooking)
	if err != nil {
		return responses.ServerError()
	}

	body, err := json.Marshal(booking)
	if err != nil {
		return responses.ServerError()
	}

	if err := s.RabbitMQ.SendMessage(environments.RabbitMQQueue, string(body)); err != nil {
		return responses.ServerError()
	}

	return responses.Created(mappings.ToBookingResponse(booking))
}
