package controllers

import (
	"fmt"
	"strconv"

	"github.com/airbornharsh/holio-go/database"
	"github.com/airbornharsh/holio-go/helpers"
	"github.com/airbornharsh/holio-go/models"
	"github.com/gin-gonic/gin"
)

func CreateHotelHandler(c *gin.Context) {
	tempuser, exists := c.Get("user")

	if !exists || (exists && tempuser != nil && tempuser.(models.User).UserType != "owner") {
		c.JSON(401, gin.H{
			"message": "Unauthorized",
		})
		return
	}

	var hotel models.Hotel

	c.BindJSON(&hotel)

	query := `INSERT INTO hotels (owner_user_id, name, description, address, phone_number, website_url, email, latitude, longitude, star_rating, avg_rating, avg_price) VALUES ('` + fmt.Sprintf("%d", tempuser.(models.User).UserID) + `', '` + hotel.Name + `', '` + hotel.Description + `', '` + hotel.Address + `', '` + hotel.PhoneNumber + `', '` + hotel.WebsiteURL + `', '` + hotel.Email + `', '` + fmt.Sprintf("%.8f", hotel.Latitude) + `', '` + fmt.Sprintf("%.8f", hotel.Longitude) + `', '` + fmt.Sprintf("%.1f", hotel.StarRating) + `', '` + fmt.Sprintf("%.2f", hotel.AvgRating) + `', '` + fmt.Sprintf("%.2f", hotel.AvgPrice) + `');`

	DB, _ := database.GetDB()
	_, err := DB.Exec(query)

	if helpers.ErrorResponse(c, err) {

		return
	}

	c.JSON(200, gin.H{
		"message": "Hotel Created SuccessFully",
	})
}

func SearchHotelsHandler(c *gin.Context) {
	q := c.Request.URL.Query()

	hotelName := q.Get("hotel_name")
	// latitude := q.Get("latitude")
	// longitude := q.Get("longitude")
	// startDate := q.Get("start_date")
	// endDate := q.Get("end_date")
	priceStart := q.Get("price_start")
	priceEnd := q.Get("price_end")
	// startRating := q.Get("start_rating")

	if hotelName == "" || priceStart == "" || priceEnd == "" {
		c.JSON(400, gin.H{
			"message": "Invalid Query Parameters",
		})
		return
	}

	query := `SELECT * FROM hotels WHERE name ILIKE '%' || '` + hotelName + `'|| '%' AND avg_price BETWEEN '` + priceStart + `' AND '` + priceEnd + `';`

	DB, _ := database.GetDB()

	var hotelRows []models.Hotel

	rows, err := DB.Query(query)

	if helpers.ErrorResponse(c, err) {
		return
	}

	for rows.Next() {
		var hotel models.Hotel

		err := rows.Scan(&hotel.HotelID, &hotel.OwnerUserId, &hotel.Name, &hotel.Description, &hotel.Address, &hotel.PhoneNumber, &hotel.WebsiteURL, &hotel.Email, &hotel.Latitude, &hotel.Longitude, &hotel.StarRating, &hotel.AvgRating, &hotel.AvgPrice)

		if helpers.ErrorResponse(c, err) {
			return
		}

		hotelRows = append(hotelRows, hotel)
	}

	c.JSON(200, gin.H{
		"message": "Search Hotels",
		"data":    hotelRows,
	})
}

func GetPopularHotelsHandler(c *gin.Context) {
	q := c.Request.URL.Query()

	minStar := q.Get("min_star")

	var query string

	if minStar == "" {
		query = `SELECT * FROM hotels ORDER BY star_rating ASC`
	} else {
		query = `SELECT * FROM hotels WHERE star_rating >= '` + minStar + `' ORDER BY star_rating ASC`
	}

	DB, _ := database.GetDB()
	row, err := DB.Query(query)

	if helpers.ErrorResponse(c, err) {
		return
	}

	defer row.Close()

	var hotelRows []models.Hotel

	for row.Next() {
		var hotel models.Hotel

		err := row.Scan(&hotel.HotelID, &hotel.OwnerUserId, &hotel.Name, &hotel.Description, &hotel.Address, &hotel.PhoneNumber, &hotel.WebsiteURL, &hotel.Email, &hotel.Latitude, &hotel.Longitude, &hotel.StarRating, &hotel.AvgRating, &hotel.AvgPrice)

		if helpers.ErrorResponse(c, err) {
			return
		}

		hotelRows = append(hotelRows, hotel)
	}

	c.JSON(200, gin.H{
		"message": "Popular Hotels",
		"data":    hotelRows,
	})
}

func GetHotelHandler(c *gin.Context) {
	hotelId := c.Param("id")

	if hotelId == "" {
		c.JSON(400, gin.H{
			"message": "Invalid Hotel Id",
		})
		return
	}

	query := "SELECT * FROM hotels WHERE hotel_id = '" + string(hotelId) + "';"

	var hotel models.Hotel

	DB, _ := database.GetDB()
	err := DB.QueryRow(query).Scan(&hotel.HotelID, &hotel.OwnerUserId, &hotel.Name, &hotel.Description, &hotel.Address, &hotel.PhoneNumber, &hotel.WebsiteURL, &hotel.Email, &hotel.Latitude, &hotel.Longitude, &hotel.StarRating, &hotel.AvgRating, &hotel.AvgPrice)

	if helpers.ErrorResponse(c, err) {
		return
	}

	c.JSON(200, gin.H{
		"message": "Found Hotel",
		"data":    hotel,
	})
}

func UpdateHotelHandler(c *gin.Context) {
	tempUser, exists := c.Get("user")

	if !exists || (exists && tempUser != nil && tempUser.(models.User).UserType != "owner") {
		c.JSON(401, gin.H{
			"message": "Unauthorized",
		})
		return
	}

	hotelId := c.Param("id")

	if hotelId == "" {
		c.JSON(400, gin.H{
			"message": "Invalid Hotel Id",
		})
		return
	}

	var tempHotel models.Hotel

	c.Bind(&tempHotel)

	var hotel models.Hotel

	query := `SELECT * FROM hotels WHERE owner_user_id = '` + fmt.Sprintf("%d", tempUser.(models.User).UserID) + `' AND hotel_id = '` + string(hotelId) + `';`

	DB, _ := database.GetDB()
	err := DB.QueryRow(query).Scan(&hotel.HotelID, &hotel.OwnerUserId, &hotel.Name, &hotel.Description, &hotel.Address, &hotel.PhoneNumber, &hotel.WebsiteURL, &hotel.Email, &hotel.Latitude, &hotel.Longitude, &hotel.StarRating, &hotel.AvgRating, &hotel.AvgPrice)

	if helpers.ErrorResponse(c, err) {
		return
	}

	hotelPtr := helpers.ReplaceHotel(&hotel, &tempHotel)

	query = `UPDATE hotels SET name = '` + hotelPtr.Name + `' , description = '` + hotelPtr.Description + `' , address = '` + hotelPtr.Address + `' , phone_number = '` + hotelPtr.PhoneNumber + `' , website_url = '` + hotelPtr.WebsiteURL + `' , email = '` + hotelPtr.Email + `' , latitude = '` + fmt.Sprintf("%.8f", hotelPtr.Latitude) + `' , longitude = '` + fmt.Sprintf("%.8f", hotelPtr.Longitude) + `' , star_rating = '` + fmt.Sprintf("%.1f", hotelPtr.StarRating) + `' , avg_rating = 	'` + fmt.Sprintf("%.2f", hotelPtr.AvgRating) + `' , avg_price = '` + fmt.Sprintf("%.2f", hotelPtr.AvgPrice) + `' WHERE hotel_id = '` + strconv.Itoa(hotelPtr.HotelID) + `';`

	_, err = DB.Exec(query)

	if helpers.ErrorResponse(c, err) {
		return
	}

	c.JSON(200, gin.H{
		"message": "Updated Hotel",
		"data":    hotel,
	})
}

func DeleteHotelHandler(c *gin.Context) {
	tempUser, exists := c.Get("user")

	if !exists || (exists && tempUser != nil && tempUser.(models.User).UserType != "owner") {
		c.JSON(401, gin.H{
			"message": "Unauthorized",
		})
		return
	}

	hotelId := c.Param("id")

	if hotelId == "" {
		c.JSON(400, gin.H{
			"message": "Invalid Hotel Id",
		})
		return
	}

	query := `DELETE FROM hotels WHERE owner_user_id = '` + fmt.Sprintf("%d", tempUser.(models.User).UserID) + `' AND hotel_id = '` + string(hotelId) + `';`

	DB, _ := database.GetDB()
	_, err := DB.Exec(query)

	if helpers.ErrorResponse(c, err) {
		return
	}

	c.JSON(200, gin.H{
		"message": "Hotel Deleted Successfully",
	})
}

func AddHotelImagesHandler(c *gin.Context) {
	tempUser, exists := c.Get("user")

	if !exists || (exists && tempUser != nil && tempUser.(models.User).UserType != "owner") {
		c.JSON(401, gin.H{
			"message": "Unauthorized",
		})
		return
	}

	hotelId := c.Param("id")

	if hotelId == "" {
		c.JSON(400, gin.H{
			"message": "Invalid Hotel Id",
		})
		return
	}

	var hotelImage models.HotelImage

	err := c.ShouldBindJSON(&hotelImage)

	if helpers.ErrorResponse(c, err) {
		return
	}

	DB, _ := database.GetDB()

	query := "SELECT EXISTS (SELECT 1 FROM hotels WHERE hotel_id = '" + hotelId + "' AND owner_user_id = '" + fmt.Sprintf("%d", tempUser.(models.User).UserID) + "');"

	var hotelExists bool

	err = DB.QueryRow(query).Scan(&hotelExists)

	if helpers.ErrorResponse(c, err) {
		return
	}

	if !hotelExists {
		c.JSON(400, gin.H{
			"message": "Hotel Does Not Exist",
		})
		return
	}

	//insert photo model here
	query = "INSERT INTO HotelImages (hotel_id, image_url, description) VALUES ('" + string(hotelId) + "' , '" + hotelImage.ImageURL + "' , '" + hotelImage.Description + "');"

	_, err = DB.Exec(query)

	if helpers.ErrorResponse(c, err) {
		return
	}

	c.JSON(200, gin.H{
		"message": "Add Hotel Photos Handler",
	})
}

func GetAllImagesHandler(c *gin.Context) {
	tempUser, exists := c.Get("user")

	if !exists || (exists && tempUser == nil) {
		c.JSON(401, gin.H{
			"message": "Unauthorized",
		})
		return
	}

	hotelId := c.Param("id")

	if hotelId == "" {
		c.JSON(400, gin.H{
			"message": "Invalid Hotel Id",
		})
		return
	}

	query := "SELECT EXISTS (SELECT 1 FROM hotels WHERE hotel_id = '" + hotelId + "' AND owner_user_id = '" + fmt.Sprintf("%d", tempUser.(models.User).UserID) + "');"

	var hotelExists bool

	DB, _ := database.GetDB()

	err := DB.QueryRow(query).Scan(&hotelExists)

	if helpers.ErrorResponse(c, err) {
		return
	}

	if !hotelExists {
		c.JSON(400, gin.H{
			"message": "Hotel Does Not Exist",
		})
		return
	}

	query = "SELECT * FROM HotelImages WHERE hotel_id = '" + hotelId + "';"

	rows, err := DB.Query(query)

	if helpers.ErrorResponse(c, err) {
		return
	}

	defer rows.Close()

	var hotelImages []models.HotelImage

	for rows.Next() {
		var hotelImage models.HotelImage

		err := rows.Scan(&hotelImage.HotelImageID, &hotelImage.HotelID, &hotelImage.ImageURL, &hotelImage.Description)

		if helpers.ErrorResponse(c, err) {
			return
		}

		hotelImages = append(hotelImages, hotelImage)
	}

	c.JSON(200, gin.H{
		"message": "All Photos",
		"images":  hotelImages,
	})
}

func ChangeHotelFacilitiesHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "ChangeHotelFacilitiesHandler",
	})
}

func ChangeHotelRoomsHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "ChangeHotelRoomsHandler",
	})
}

func ChangeHotelAmenitiesHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "ChangeHotelAmenitiesHandler",
	})
}
