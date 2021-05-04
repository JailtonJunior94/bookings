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
	TelegramBaseUrl    = ""
	BotKey             = ""
	ChatId             = int64(0)
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
	TelegramBaseUrl = viper.GetString("telegram.baseURL")
	BotKey = viper.GetString("telegram.botKey")
	ChatId = viper.GetInt64("telegram.chatID")
}
