package models

type Review struct {
	ReviewID   int     `json:"review_id"`
	UserID     int     `json:"user_id"`
	HotelID    int     `json:"hotel_id"`
	Rating     float64 `json:"rating"`
	ReviewText string  `json:"review_text"`
	ReviewDate string  `json:"review_date"`
}
