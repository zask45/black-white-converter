package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func postHandler(c *gin.Context) {
	file, _ := c.FormFile("file")
	c.SaveUploadedFile(file, "uploads/"+file.Filename)
	c.JSON(http.StatusOK, gin.H{"message": "Upload sukses"})
}

func main() {
	router := gin.Default()
	router.POST("/upload", postHandler)
	fmt.Println("http://localhost:8080")
	router.Run(":8080")
}
