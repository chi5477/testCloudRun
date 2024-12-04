package main

import (
	"github.com/gin-gonic/gin"
)

func main() {

	a := 1
	b := 2
	if a != b {
		panic("a is not equal to b")
	}

	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello Cloud Run!",
		})
	})
	router.Run(":8080")
}
