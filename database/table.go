package database

import (
	"database/sql"
	"fmt"
	"log"
)

func MakeTable(db *sql.DB) {
	fmt.Println("Creating tables")
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to database")

	makeUsersTable(db)
	makeHotelsTable(db)
	makeHotelFacilitiesTable(db)
	makeRoomsTable(db)
	makeAmenitiesTable(db)
	makeRoomAmenitiesTable(db)
	makeBookingsTable(db)
	makeReviewsTable(db)
	makePaymentsTable(db)
	makeHotelImagesTable(db)

	fmt.Println("Tables created")
}

func makeUsersTable(db *sql.DB) {
	c := `
		CREATE TABLE IF NOT EXISTS Users (
			user_id SERIAL PRIMARY KEY,
			username VARCHAR(255) NOT NULL,
			password VARCHAR(255) NOT NULL,
			user_type VARCHAR(50) NOT NULL,
			email VARCHAR(255) NOT NULL,
			full_name VARCHAR(255),
			address TEXT,
			phone_number VARCHAR(20)
		);
	`
	createTableIfNot(db, c, "Users")
}

func makeHotelsTable(db *sql.DB) {
	c := `
		CREATE TABLE IF NOT EXISTS Hotels (
			hotel_id SERIAL PRIMARY KEY,
			name VARCHAR(255) NOT NULL,
			description TEXT,
			address TEXT,
			phone_number VARCHAR(20),
			website_url TEXT,
			email VARCHAR(255) NOT NULL,
			latitude DECIMAL(10, 8) NOT NULL,
			longitude DECIMAL(11, 8) NOT NULL,
			star_rating DECIMAL(2, 1) NOT NULL,
			avg_rating DECIMAL(3, 2) NOT NULL,
			avg_price DECIMAL(10, 2) NOT NULL
		);
	`
	createTableIfNot(db, c, "Hotels")
}

func makeHotelFacilitiesTable(db *sql.DB) {
	c := `
		CREATE TABLE IF NOT EXISTS HotelFacilities (
			facility_id SERIAL PRIMARY KEY,
			hotel_id INT REFERENCES Hotels(hotel_id),
			name VARCHAR(255) NOT NULL,
			description TEXT
		);
	`
	createTableIfNot(db, c, "HotelFacilities")
}

func makeRoomsTable(db *sql.DB) {
	c := `
		CREATE TABLE IF NOT EXISTS Rooms (
			room_id SERIAL PRIMARY KEY,
			hotel_id INT REFERENCES Hotels(hotel_id),
			room_number INT NOT NULL,
			room_type VARCHAR(100) NOT NULL,
			description TEXT,
			price DECIMAL(10, 2) NOT NULL,
			rating DECIMAL(3, 2) NOT NULL,
			availability BOOLEAN NOT NULL
		);
	`
	createTableIfNot(db, c, "Rooms")
}

func makeAmenitiesTable(db *sql.DB) {
	c := `
		CREATE TABLE IF NOT EXISTS Amenities (
			amenity_id SERIAL PRIMARY KEY,
			Hotel_id INT REFERENCES Hotels(hotel_id),
      name VARCHAR(255) NOT NULL,
      description TEXT
		);
	`
	createTableIfNot(db, c, "Amenities")
}

func makeRoomAmenitiesTable(db *sql.DB) {
	c := `
		CREATE TABLE IF NOT EXISTS RoomAmenities (
			room_id INT REFERENCES Rooms(room_id),
			amenity_id INT REFERENCES Amenities(amenity_id),
			PRIMARY KEY (room_id, amenity_id)
		);
	`
	createTableIfNot(db, c, "RoomAmenities")
}

func makeBookingsTable(db *sql.DB) {
	c := `
		CREATE TABLE IF NOT EXISTS Bookings (
			booking_id SERIAL PRIMARY KEY,
      user_id INT REFERENCES Users(user_id),
      room_id INT REFERENCES Rooms(room_id),
      check_in_date DATE NOT NULL,
      check_out_date DATE NOT NULL,
      total_price DECIMAL(10, 2) NOT NULL,
      booking_status VARCHAR(50) NOT NULL
		);
	`
	createTableIfNot(db, c, "Bookings")
}

func makeReviewsTable(db *sql.DB) {
	c := `
		CREATE TABLE IF NOT EXISTS Reviews (
			review_id SERIAL PRIMARY KEY,
			user_id INT REFERENCES Users(user_id),
			hotel_id INT REFERENCES Hotels(hotel_id),
			rating DECIMAL(3, 2) NOT NULL,
			review_text TEXT,
			review_date DATE NOT NULL
		);
	`
	createTableIfNot(db, c, "Reviews")
}

func makePaymentsTable(db *sql.DB) {
	c := `
		CREATE TABLE IF NOT EXISTS Payments (
			payment_id SERIAL PRIMARY KEY,
			user_id INT REFERENCES Users(user_id),
			booking_id INT REFERENCES Bookings(booking_id),
			payment_date DATE NOT NULL,
			amount DECIMAL(10, 2) NOT NULL,
			payment_status VARCHAR(50) NOT NULL
		);
	`
	createTableIfNot(db, c, "Payments")
}

func makeHotelImagesTable(db *sql.DB) {
	c := `
		CREATE TABLE IF NOT EXISTS HotelImages (
			image_id SERIAL PRIMARY KEY,
			hotel_id INT REFERENCES Hotels(hotel_id),
			image_url TEXT NOT NULL,
			description TEXT
		);
	`
	createTableIfNot(db, c, "HotelImages")
}

func createTableIfNot(DB *sql.DB, c string, tableName string) error {
	_, err := DB.Exec(c)
	if err != nil {
		fmt.Println(tableName)
		log.Fatal(err)
		return err
	}
	return err
}

func DropTables(db *sql.DB) {
	dropTables := `DROP TABLE IF EXISTS HotelImages;
		DROP TABLE IF EXISTS Payments;
		DROP TABLE IF EXISTS Reviews;
		DROP TABLE IF EXISTS Bookings;
		DROP TABLE IF EXISTS RoomAmenities;
		DROP TABLE IF EXISTS Amenities;
		DROP TABLE IF EXISTS Rooms;
		DROP TABLE IF EXISTS HotelFacilities;
		DROP TABLE IF EXISTS Hotels;
		DROP TABLE IF EXISTS Users;
	`
	_, err := db.Exec(dropTables)
	if err != nil {
		fmt.Println("Error dropping tables")
		log.Fatal(err)
	}

	fmt.Println("Tables dropped")

}
