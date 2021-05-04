package environments

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/viper"
)

var (
	MongoConnectionString = ""
	BookingDatabase       = ""
	UsersCollection       = ""
	BookingCollection     = ""
	RabbitMQConnection    = ""
	RabbitMQQueue         = ""
	JwtSecret             = ""
)

func New() {
	var err error

	viper.SetConfigName(fmt.Sprintf("appsettings.%s", os.Getenv("ENVIRONMENT")))
	viper.SetConfigType("json")
	viper.AddConfigPath(".")

	err = viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}

	MongoConnectionString = viper.GetString("mongo.connection")
	BookingDatabase = viper.GetString("mongo.database")
	UsersCollection = viper.GetString("mongo.userCollection")
	BookingCollection = viper.GetString("mongo.bookingCollection")
	RabbitMQConnection = viper.GetString("rabbitMQ.connection")
	RabbitMQQueue = viper.GetString("rabbitMQ.queue")
	JwtSecret = viper.GetString("security.jwtSecret")
}
