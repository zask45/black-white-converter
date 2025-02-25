package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func rootHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"Name":    "Yuta Okkotsu",
		"Age":     20,
		"Address": "Tokyo",
	})
}

func main() {
	router := gin.Default()
	router.GET("/", rootHandler)
	fmt.Println("http://localhost:8080")
	router.Run(":8080")
}
