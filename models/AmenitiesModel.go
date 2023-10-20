package models

// Amenities struct is a row record of the amenities table in the ticketing database
type Amenities struct {
	AmenityID   int    `json:"amenity_id"`
	HotelID     int    `json:"hotel_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}