package infrastructure

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/viper"
)

var (
	RabbitMQConnection = ""
	RabbitMQQueue      = ""
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

	RabbitMQConnection = viper.GetString("rabbitMQ.connection")
	RabbitMQQueue = viper.GetString("rabbitMQ.queue")
}
