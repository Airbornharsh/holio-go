package models

type HotelImage struct {
	HotelImageID int    `json:"hotel_image_id"`
	HotelID      int    `json:"hotel_id"`
	ImageURL     string `json:"image_url"`
	Description  string `json:"description"`
}
