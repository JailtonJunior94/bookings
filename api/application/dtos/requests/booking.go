package requests

import (
	"errors"
	"time"
)

type Booking struct {
	Date time.Time `json:"date"`
}

func (b *Booking) IsValid() error {
	if b.Date == time.Now() {
		return errors.New("O Nome é obrigatório")
	}

	return nil
}
