package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/airbornharsh/holio-go/database"
	"github.com/airbornharsh/holio-go/routes"
)

func main() {
	database.DBInit()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r := gin.New()
	// r.Use(gin.Logger())
	gin.SetMode(gin.ReleaseMode)

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to Holio",
		})
	})

	routes.Router(r)

	fmt.Println("Server Started at Port " + port)
	r.Run(":" + port)

}
