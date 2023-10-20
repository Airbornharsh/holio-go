package models

type User struct {
	UserID      int    `json:"user_id"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	UserType    string `json:"user_type"`
	Email       string `json:"email"`
	FullName    string `json:"full_name"`
	Address     string `json:"address"`
	PhoneNumber string `json:"phone_number"`
}
