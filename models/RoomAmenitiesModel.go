package models

type RoomAmenities struct {
	RoomAmenityID int `json:"room_amenity_id"`
	RoomID        int `json:"room_id"`
	AmenityID     int `json:"amenity_id"`
}