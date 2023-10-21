package controllers

import (
	"fmt"

	"github.com/airbornharsh/holio-go/database"
	"github.com/airbornharsh/holio-go/helpers"
	"github.com/airbornharsh/holio-go/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func SignupHandler(c *gin.Context) {
	var user models.User

	//Bind JSON to Struct
	if err := c.BindJSON(&user); err != nil {
		fmt.Println(err)
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
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
	if err != nil {
		fmt.Println(err)
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return

	}

	//Check if User Already Exists
	s := `SELECT EXISTS(SELECT 1 FROM users WHERE username = '` + user.Username + `' OR email = '` + user.Email + `' OR phone_number = '` + user.PhoneNumber + `');`

	var exists bool
	rows, err := DB.Query(s)
	if err != nil {
		fmt.Println(err)
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return

	}

	if rows.Next() {
		err = rows.Scan(&exists)
		if err != nil {
			fmt.Println(err)
			c.JSON(500, gin.H{
				"message": err.Error(),
			})
			return
		}
		c.JSON(200, gin.H{
			"message": "User Already Exists",
		})
		return
	}

	//Hashing Password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	if err != nil {
		fmt.Println(err)
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	user.Password = string(hashedPassword)
	user.UserType = "user"

	//Inserting User
	s = "INSERT INTO Users (username,password,user_type,email,full_name,address,phone_number) VALUES('" + user.Username + "','" + user.Password + "','" + user.UserType + "','" + user.Email + "','" + user.FullName + "','" + user.Address + "','" + user.PhoneNumber + "');"

	_, err = DB.Exec(s)
	if err != nil {
		fmt.Println(err)
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	token, err := helpers.GenerateToken(&user)

	if err != nil {
		fmt.Println(err)
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
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
	if err := c.BindJSON(&tempuser); err != nil {
		fmt.Println(err)
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	//Getting DB
	DB, err := database.GetDB()
	if err != nil {
		fmt.Println(err)
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	//Check if User Already Exists
	s := `SELECT * FROM users WHERE username = '` + tempuser.Username + `' OR email = '` + tempuser.Email + `' OR phone_number = '` + tempuser.PhoneNumber + `';`

	rows, err := DB.Query(s)
	if err != nil {
		fmt.Println(err)
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	if rows.Next() {
		err = rows.Scan(&user.UserID, &user.Username, &user.Password, &user.UserType, &user.Email, &user.FullName, &user.Address, &user.PhoneNumber)
		if err != nil {
			fmt.Println(err)
			c.JSON(500, gin.H{
				"message": err.Error(),
			})
			return
		}

		//Comparing Password
		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(tempuser.Password))
		if err != nil {
			fmt.Println(err)
			c.JSON(500, gin.H{
				"message": err.Error(),
			})
			return
		}

		//Generate Token
		token, err := helpers.GenerateToken(&user)

		if err != nil {
			fmt.Println(err)
			c.JSON(500, gin.H{
				"message": err.Error(),
			})
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
	c.JSON(200, gin.H{
		"message": "ChangePasswordHandler",
	})
}

func ChangeEmailHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "ChangeEmailHandler",
	})
}

func ChangePhoneHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "ChangePhoneHandler",
	})
}
