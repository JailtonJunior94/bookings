package requests

import (
	"errors"
	"time"
)

type Booking struct {
	Date time.Time `json:"date"`
}

func (b *Booking) IsValid() error {
	if b.Date.IsZero() {
		return errors.New("A Data é obrigatória")
	}

	return nil
}
