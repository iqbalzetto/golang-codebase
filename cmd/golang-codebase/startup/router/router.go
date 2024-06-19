package router

import (
	"github.com/gin-gonic/gin"
	"github.com/iqbalzetto/golang-codebase/internal/user"
)

/*
|--------------------------------------------------------------------------
| A. Set up Router
|--------------------------------------------------------------------------
|
| The router is responsible for handling incoming HTTP requests and sending
| responses back to the client. We will define the routes and their handlers
| in this file, allowing us to keep our main file clean and organized.
|
*/

func GinRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.GET("/user", user.GetUser)

	return router

}
