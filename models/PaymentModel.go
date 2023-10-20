package models

type Payment struct {
	PaymentID     int    `json:"payment_id"`
	UserID        int    `json:"user_id"`
	BookingID     int    `json:"booking_id"`
	Amount        int    `json:"amount"`
	PaymentDate   string `json:"payment_date"`
	PaymentStatus string `json:"payment_status"`
}
