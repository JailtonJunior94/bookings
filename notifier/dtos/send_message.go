package dtos

type SendMessageReqBody struct {
	ChatID int64  `json:"chat_id"`
	Text   string `json:"text"`
}

func NewSendMessage(chatID int64, text string) *SendMessageReqBody {
	return &SendMessageReqBody{
		ChatID: chatID,
		Text:   text,
	}
}
