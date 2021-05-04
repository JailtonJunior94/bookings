package responses

import "time"

type BookingResponse struct {
	ID   string        `json:"id"`
	Date time.Time     `json:"date"`
	User *UserResponse `json:"user"`
}
