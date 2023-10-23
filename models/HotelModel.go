package models

type Hotel struct {
	HotelID     int     `json:"hotel_id"`
	OwnerUserId int     `json:"owner_user_id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Address     string  `json:"address"`
	PhoneNumber       string  `json:"phone_number"`
	WebsiteURL  string  `json:"website_url"`
	Email       string  `json:"email"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
	StarRating  float64 `json:"star_rating"`
	AvgRating   float64 `json:"avg_rating"`
	AvgPrice    float64 `json:"avg_price"`
}
