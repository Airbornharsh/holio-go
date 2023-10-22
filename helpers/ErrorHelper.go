package helpers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func ErrorResponse(c *gin.Context, err error) bool {
	if err != nil {
		fmt.Println(err)
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return true
	}
	return false
}
