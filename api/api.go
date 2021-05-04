package main

import (
	"fmt"
	"log"
	"os"

	"github.com/jailtonjunior94/bookings/api/infrastructure/database"
	"github.com/jailtonjunior94/bookings/api/infrastructure/environments"
	"github.com/jailtonjunior94/bookings/api/infrastructure/ioc"
	"github.com/jailtonjunior94/bookings/api/presentation/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()

	app.Use(cors.New())
	app.Use(logger.New())

	environments.New()

	mongoConnection := database.NewConnection()
	defer mongoConnection.Disconnect()

	ioc.SetupDependencyInjection(mongoConnection)

	routes.SetupRoutes(app)

	port := os.Getenv("PORT")
	fmt.Printf("🚀 API is running on http://localhost:%v", port)
	log.Fatal(app.Listen(fmt.Sprintf(":%v", port)))
}
