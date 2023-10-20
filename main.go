package main

import (
	"os"

	"github.com/gin-gonic/gin"

	"github.com/airbornharsh/holio-go/database"
	"github.com/airbornharsh/holio-go/routes"
)

type widgets struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func main() {
	database.DBInit()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r := gin.Default()
	r.Use(gin.Logger())

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to Holio",
		})
	})

	routes.Router(r)

	r.Run(":" + port)

	// c := "CREATE TABLE IF NOT EXISTS widgets (id serial PRIMARY KEY, name text NOT NULL)"
	// _, err = conc.Exec(c)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// defer conc.Close()

	// c = "SELECT name FROM widgets WHERE id = " + "1"

	// var name string

	// rows := conc.QueryRow(c)
	// err = rows.Scan(&name)
	// if err != nil {
	// 	fmt.Println("Error in query statement")
	// 	log.Fatal(err)
	// }

	// defer conc.Close()

	// fmt.Println(name)
}
