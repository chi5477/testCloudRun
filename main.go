package main

import (
	"github.com/gin-gonic/gin"
)

func main() {

	a := 2
	b := 2
	if a != b {
		panic("a is not equal to b")
	}

	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello Cloud Run! We make a change to the code",
		})
	})
	router.Run(":8080")
}
