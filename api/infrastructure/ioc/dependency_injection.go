package ioc

import (
	"github.com/jailtonjunior94/bookings/api/application/services"
	"github.com/jailtonjunior94/bookings/api/domain/interfaces"
	"github.com/jailtonjunior94/bookings/api/infrastructure/adapters"
	"github.com/jailtonjunior94/bookings/api/infrastructure/bus"
	"github.com/jailtonjunior94/bookings/api/infrastructure/database"
	"github.com/jailtonjunior94/bookings/api/infrastructure/environments"
	"github.com/jailtonjunior94/bookings/api/infrastructure/repositories"
	"github.com/jailtonjunior94/bookings/api/presentation/controllers"
)

var (
	MongoConnection   database.IMongoConnection
	HashAdapter       adapters.IHashAdapter
	JwtAdapter        adapters.IJwtAdapter
	UserRepository    interfaces.IUserRepository
	BookingRepository interfaces.IBookingRepository
	RabbitMQ          interfaces.IRabbitMQ
	UserService       interfaces.IUserService
	AuthService       interfaces.IAuthService
	BookingService    interfaces.IBookingService
	UserController    *controllers.UserController
	AuthController    *controllers.AuthController
	BookingController *controllers.BookingController
)

func SetupDependencyInjection(mongoConnection database.IMongoConnection) {
	/* Database */
	MongoConnection = mongoConnection

	/* Adapters */
	HashAdapter = adapters.NewHashAdapter()
	JwtAdapter = adapters.NewJwtAdapter()

	/* Bus */
	RabbitMQ = bus.New(environments.RabbitMQConnection)

	/* Repositories */
	UserRepository = repositories.NewUserRepository(MongoConnection)
	BookingRepository = repositories.NewBookingRepository(MongoConnection)

	/* Services */
	UserService = services.NewUserService(UserRepository, HashAdapter)
	AuthService = services.NewAuthService(UserRepository, HashAdapter, JwtAdapter)
	BookingService = services.NewBookingService(BookingRepository, RabbitMQ, UserRepository)

	/* Controllers */
	UserController = controllers.NewUserController(UserService)
	AuthController = controllers.NewAuthController(AuthService)
	BookingController = controllers.NewBookingController(BookingService, JwtAdapter)
}
