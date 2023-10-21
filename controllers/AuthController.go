package controllers

import (
	"log"

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
		log.Fatal(err)
		c.JSON(500, gin.H{
			"message": "Error Binding JSON",
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
		log.Fatal("Error Getting DB:", err)
		c.JSON(500, gin.H{
			"message": "Error Getting DB",
		})
		return
		
	}
	
	//Check if User Already Exists
	s := `SELECT EXISTS(SELECT 1 FROM users WHERE username = '` + user.Username + `' OR email = '` + user.Email + `' OR phone_number = '` + user.PhoneNumber + `');`

	var exists bool
	rows, err := DB.Query(s)
	if err != nil {
		log.Fatal(err)
		c.JSON(500, gin.H{
			"message": "Error Querying DB",
		})
		return

	}

	if rows.Next() {
		err = rows.Scan(&exists)
		if err != nil {
			log.Fatal(err)
			c.JSON(500, gin.H{
				"message": "Error Scanning Rows",
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
		log.Fatal("Error While hashing the Password:", err)
		c.JSON(500, gin.H{
			"message": "Error While hashing the Password",
		})
		return
	}

	user.Password = string(hashedPassword)
	user.UserType = "user"

	//Inserting User
	s = "INSERT INTO Users (username,password,user_type,email,full_name,address,phone_number) VALUES('" + user.Username + "','" + user.Password + "','" + user.UserType + "','" + user.Email + "','" + user.FullName + "','" + user.Address + "','" + user.PhoneNumber + "');"

	_, err = DB.Exec(s)
	if err != nil {
		log.Fatal(err)
		c.JSON(500, gin.H{
			"message": "Error Executing Query",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "User Created",
	})
}

func LoginHanlder(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "LoginHanlder",
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
