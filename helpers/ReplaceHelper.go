package helpers

import "github.com/airbornharsh/holio-go/models"

func ReplaceHotel(hotel *models.Hotel, newHotel *models.Hotel) *models.Hotel {
	if newHotel.Name != "" {
		hotel.Name = newHotel.Name
	}
	if newHotel.Description != "" {
		hotel.Description = newHotel.Description
	}
	if newHotel.Address != "" {
		hotel.Address = newHotel.Address
	}
	if newHotel.PhoneNumber != "" {
		hotel.PhoneNumber = newHotel.PhoneNumber
	}
	if newHotel.WebsiteURL != "" {
		hotel.WebsiteURL = newHotel.WebsiteURL
	}
	if newHotel.Email != "" {
		hotel.Email = newHotel.Email
	}
	if newHotel.Latitude != 0 {
		hotel.Latitude = newHotel.Latitude
	}
	if newHotel.Longitude != 0 {
		hotel.Longitude = newHotel.Longitude
	}
	if newHotel.StarRating != 0 {
		hotel.StarRating = newHotel.StarRating
	}
	if newHotel.AvgRating != 0 {
		hotel.AvgRating = newHotel.AvgRating
	}
	if newHotel.AvgPrice != 0 {
		hotel.AvgPrice = newHotel.AvgPrice
	}

	return hotel
}

func ReplaceUser(user *models.User, newUser *models.User) *models.User {
	if newUser.Username != "" {
		user.Username = newUser.Username
	}
	if newUser.Address != "" {
		user.Address = newUser.Address
	}
	if newUser.FullName != "" {
		user.FullName = newUser.FullName
	}
	return user
}

func ReplaceRoom(room *models.Room, newRoom *models.Room) *models.Room {
	if newRoom.RoomNumber != 0 {
		room.RoomNumber = newRoom.RoomNumber
	}
	if newRoom.RoomType != "" {
		room.RoomType = newRoom.RoomType
	}
	if newRoom.RoomCapacity != 0 {
		room.RoomCapacity = newRoom.RoomCapacity
	}
	if newRoom.Description != "" {
		room.Description = newRoom.Description
	}
	if newRoom.Price != 0 {
		room.Price = newRoom.Price
	}
	if newRoom.Rating != 0 {
		room.Rating = newRoom.Rating
	}
	return room
}
