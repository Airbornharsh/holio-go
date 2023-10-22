package middlewares

import (
	"github.com/airbornharsh/holio-go/database"
	"github.com/airbornharsh/holio-go/helpers"
	"github.com/airbornharsh/holio-go/models"

	"github.com/gin-gonic/gin"
)

func TokenVerifyMiddleWare(c *gin.Context) {
	Auth := c.Request.Header.Get("Authorization")

	if Auth == "" {
		c.JSON(400, gin.H{
			"message": "Authorization Header Not Found",
		})
		c.Set("user", nil)

		c.Next()
	}

	// tokenString := strings.Split(Auth, " ")[1]

	// if tokenString == "" {
	// 	c.JSON(400, gin.H{
	// 		"message": "Token not found",
	// 	})
	// 	c.Set("user", nil)

	// 	c.Next()
	// }

	userId, err := helpers.GetClaims(c, Auth)

	if err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		c.Set("user", nil)

		c.Next()
		return
	}

	if userId == "" {
		c.JSON(400, gin.H{
			"message": "Invalid Token",
		})
		c.Set("user", nil)

		c.Next()
		return
	}

	var user models.User

	DB, err := database.GetDB()

	if err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		c.Set("user", nil)

		c.Next()
		return
	}

	query := `SELECT * FROM users WHERE user_id = '` + userId + `';`

	err = DB.QueryRow(query).Scan(&user.UserID, &user.Username, &user.Password, &user.UserType, &user.Email, &user.FullName, &user.Address, &user.PhoneNumber)

	if err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		c.Set("user", nil)

		c.Next()
		return
	}

	c.Set("user", user)

	c.Next()
}
