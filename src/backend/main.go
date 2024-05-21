package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusOK)
			return
		}

		c.Next()
	}
}
func main() {

	r := gin.Default()
	r.Use(corsMiddleware())
	r.OPTIONS("/api/usr/login", func(c *gin.Context) {

		fmt.Println("Request Headers: ", c.Request.Header)
		fmt.Println("Request Body: ", c.Request.Body)
		c.JSON(200, gin.H{"message": "Options Request!"})

	})
	r.POST("/api/usr/login", func(c *gin.Context) {
		fmt.Println("Request Headers: ", c.Request.Header)
		fmt.Println("Request Body: ", c.Request.Body)
		c.JSON(200, gin.H{"message": "Post Request!"})
	})

	r.Run(":8081")
}
