package user

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	// Get user from database
	fmt.Printf("%v", c)
}
