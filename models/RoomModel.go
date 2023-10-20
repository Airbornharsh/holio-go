package models

type Room struct {
	RoomID       int     `json:"room_id"`
	HotelID      int     `json:"hotel_id"`
	RoomNumber   int     `json:"room_number"`
	RoomType     string  `json:"room_type"`
	Description  string  `json:"description"`
	Price        float64 `json:"price"`
	Rating       float64 `json:"rating"`
	Availability bool    `json:"availability"`
}
