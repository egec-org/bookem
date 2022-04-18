package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.POST("/appointments", func(c *gin.Context) {
		c.JSON(201, gin.H{
			"message": "Your appointment is booked.",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
