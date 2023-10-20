package models

type HotelFacilities struct {
	FacilityID  int    `json:"facility_id"`
	HotelID     int    `json:"hotel_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
