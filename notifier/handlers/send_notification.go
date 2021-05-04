package handlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/jailtonjunior94/bookings/api/application/dtos/responses"
	"github.com/jailtonjunior94/bookings/notifier/dtos"
	"github.com/jailtonjunior94/bookings/notifier/infrastructure"
)

type INotificationHandler interface {
	SendMessage(data []byte) error
}

type NotificationHandler struct{}

func NewNotificationHandler() INotificationHandler {
	return &NotificationHandler{}
}

func (t *NotificationHandler) SendMessage(data []byte) error {
	var booking responses.BookingResponse
	if err := json.Unmarshal(data, &booking); err != nil {
		return err
	}

	text := `
			 Olá: %s

			Sua reserva para a data: %s
			
			Está confirmada, obrigado!
	`

	request := dtos.NewSendMessage(infrastructure.ChatId, fmt.Sprintf(text, booking.User.Name, booking.Date.Format("January 2, 2006")))
	reqBytes, err := json.Marshal(&request)
	if err != nil {
		return err
	}

	url := fmt.Sprintf("%s%s/sendMessage", infrastructure.TelegramBaseUrl, infrastructure.BotKey)
	res, err := http.Post(url, "application/json", bytes.NewBuffer(reqBytes))

	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusOK {
		return errors.New("Não foi possível enviar mensagem")
	}

	return nil
}
