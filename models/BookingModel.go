package models

type Booking struct {
	BookingID    int    `json:"booking_id"`
	UserID       int    `json:"user_id"`
	RoomID       int    `json:"room_id"`
	CheckInDate  string `json:"check_in_date"`
	CheckOutDate string `json:"check_out_date"`
	TotalPrice   int    `json:"total_price"`
	Status       string `json:"status"`
}
