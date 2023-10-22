package controllers

import (
	"strconv"

	"github.com/airbornharsh/holio-go/database"
	"github.com/airbornharsh/holio-go/helpers"
	"github.com/airbornharsh/holio-go/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func SignupHandler(c *gin.Context) {
	var user models.User

	//Bind JSON to Struct
	err := c.BindJSON(&user)
	if helpers.ErrorResponse(c, err) {
		return
	}

	//Check PassWord Valid or not
	if !helpers.IsValidPassword(user.Password) {
		c.JSON(400, gin.H{
			"message": "Password is not valid",
		})
		return
	}

	//Getting DB
	DB, err := database.GetDB()
	if helpers.ErrorResponse(c, err) {
		return
	}

	//Check if User Already Exists
	query := `SELECT EXISTS(SELECT 1 FROM users WHERE username = '` + user.Username + `' OR email = '` + user.Email + `' OR phone_number = '` + user.PhoneNumber + `');`

	var exists bool
	rows, err := DB.Query(query)
	if helpers.ErrorResponse(c, err) {
		return
	}

	if rows.Next() {
		err = rows.Scan(&exists)
		if helpers.ErrorResponse(c, err) {
			return
		}

		if exists {
			c.JSON(200, gin.H{
				"message": "User Already Exists",
			})
			return
		}
	}

	//Hashing Password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	if helpers.ErrorResponse(c, err) {
		return
	}

	user.Password = string(hashedPassword)
	user.UserType = "user"

	//Inserting User
	query = "INSERT INTO Users (username,password,user_type,email,full_name,address,phone_number) VALUES('" + user.Username + "','" + user.Password + "','" + user.UserType + "','" + user.Email + "','" + user.FullName + "','" + user.Address + "','" + user.PhoneNumber + "');"

	_, err = DB.Exec(query)
	if helpers.ErrorResponse(c, err) {
		return
	}

	token, err := helpers.GenerateToken(&user)

	if helpers.ErrorResponse(c, err) {
		return
	}

	c.Writer.Header().Set("Authorization", "Bearer "+token)
	c.JSON(200, gin.H{
		"message": "User Created",
		"token":   token,
		"userData": gin.H{
			"username":     user.Username,
			"email":        user.Email,
			"full_name":    user.FullName,
			"address":      user.Address,
			"phone_number": user.PhoneNumber,
			"user_type":    user.UserType,
		},
	})
}

func LoginHanlder(c *gin.Context) {
	var tempuser models.User
	var user models.User

	//Bind JSON to Struct
	err := c.BindJSON(&tempuser)
	if helpers.ErrorResponse(c, err) {
		return
	}

	//Getting DB
	DB, err := database.GetDB()
	if helpers.ErrorResponse(c, err) {
		return
	}

	//Check if User Already Exists
	query := `SELECT * FROM users WHERE username = '` + tempuser.Username + `' OR email = '` + tempuser.Email + `' OR phone_number = '` + tempuser.PhoneNumber + `';`

	rows, err := DB.Query(query)
	if helpers.ErrorResponse(c, err) {
		return
	}

	if rows.Next() {
		err = rows.Scan(&user.UserID, &user.Username, &user.Password, &user.UserType, &user.Email, &user.FullName, &user.Address, &user.PhoneNumber)
		if helpers.ErrorResponse(c, err) {
			return
		}
		//Comparing Password
		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(tempuser.Password))
		if helpers.ErrorResponse(c, err) {
			return
		}

		//Generate Token
		token, err := helpers.GenerateToken(&user)

		if helpers.ErrorResponse(c, err) {
			return
		}

		c.Writer.Header().Set("Authorization", "Bearer "+token)
		c.JSON(200, gin.H{
			"message": "Login Successful",
			"token":   token,
			"userData": gin.H{
				"username":     user.Username,
				"email":        user.Email,
				"full_name":    user.FullName,
				"address":      user.Address,
				"phone_number": user.PhoneNumber,
				"user_type":    user.UserType,
			},
		})
		return
	}

	c.JSON(400, gin.H{
		"message": "User Does Not Exists",
	})
}

func LogoutHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "LogoutHandler",
	})
}

func ForgotPasswordHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "ForgotPasswordHandler",
	})
}

func ResetPasswordHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "ResetPasswordHandler",
	})
}

func ChangePasswordHandler(c *gin.Context) {
	tempuser, exists := c.Get("user")

	if !(exists && tempuser != nil) {
		return
	}

	type PassWord struct {
		OldPassword string `json:"old_password"`
		NewPassword string `json:"new_password"`
	}

	var passWord PassWord

	c.BindJSON(&passWord)

	user := tempuser.(models.User)

	query := "SELECT password FROM users WHERE user_id = '" + strconv.Itoa(user.UserID) + "';"

	DB, err := database.GetDB()
	if helpers.ErrorResponse(c, err) {
		return
	}

	rows, err := DB.Query(query)
	if helpers.ErrorResponse(c, err) {
		return
	}
	defer rows.Close()

	var password string
	if rows.Next() {
		err = rows.Scan(&password)
		if helpers.ErrorResponse(c, err) {
			return
		}
	}

	//Comparing Password
	err = bcrypt.CompareHashAndPassword([]byte(password), []byte(passWord.OldPassword))
	if helpers.ErrorResponse(c, err) {
		return
	}

	//Hashing New Password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(passWord.NewPassword), 14)
	if helpers.ErrorResponse(c, err) {
		return
	}
	//Updating Password
	query = "UPDATE users SET password = '" + string(hashedPassword) + "' WHERE user_id = '" + strconv.Itoa(user.UserID) + "';"
	_, err = DB.Exec(query)

	if helpers.ErrorResponse(c, err) {
		return
	}

	c.JSON(200, gin.H{
		"message": "Password updated successfully",
	})
}

func ChangeEmailHandler(c *gin.Context) {
	tempUser, exits := c.Get("user")

	if !(exits && tempUser != nil) {
		return
	}

	type Email struct {
		NewEmail string `json:"new_email"`
	}

	var email Email

	c.BindJSON(&email)

	query := "UPDATE users SET email = '" + email.NewEmail + "' WHERE user_id = '" + strconv.Itoa(tempUser.(models.User).UserID) + "';"

	DB, err := database.GetDB()
	if helpers.ErrorResponse(c, err) {
		return
	}

	_, err = DB.Exec(query)
	if helpers.ErrorResponse(c, err) {
		return
	}

	c.JSON(200, gin.H{
		"message": "Email updated successfully",
	})
}

func ChangePhoneHandler(c *gin.Context) {
	tempUser, exits := c.Get("user")

	if !(exits && tempUser != nil) {
		return
	}

	type PhoneNumber struct {
		NewPhoneNumber string `json:"new_phone_number"`
	}

	var phoneNumber PhoneNumber

	c.BindJSON(&phoneNumber)

	query := "UPDATE users SET phone_number = '" + phoneNumber.NewPhoneNumber + "' WHERE user_id = '" + strconv.Itoa(tempUser.(models.User).UserID) + "';"

	DB, err := database.GetDB()
	if helpers.ErrorResponse(c, err) {
		return
	}

	_, err = DB.Exec(query)
	if helpers.ErrorResponse(c, err) {
		return
	}

	c.JSON(200, gin.H{
		"message": "Phone updated successfully",
	})
}
