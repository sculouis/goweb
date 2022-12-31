package main

import (
	. "GoWeb/src"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.SetTrustedProxies([]string{"localhost"})
	v1 := router.Group("/v1")
	AddUserRouter(v1)

	// router.GET("/", func(c *gin.Context) {
	// 	// If the client is 192.168.1.2, use the X-Forwarded-For
	// 	// header to deduce the original client IP from the trust-
	// 	// worthy parts of that header.
	// 	// Otherwise, simply return the direct client IP
	// 	c.JSON(200, gin.H{
	// 		"message":  "ping",
	// 		"message2": "Sucess!",
	// 	})
	// })

	// router.POST("/ping/:id", func(c *gin.Context) {
	// 	id := c.Param("id")
	// 	c.JSON(200, gin.H{
	// 		"id": id,
	// 	})
	// })

	router.Run(":8080")
}
