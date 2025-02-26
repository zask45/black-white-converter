package main

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"image/png"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
)

func postHandler(c *gin.Context) {
	file, _ := c.FormFile("file")

	filename := file.Filename
	folder := "uploads/"
	c.SaveUploadedFile(file, folder+filename)

	result_file := processImage(c, folder, filename)
	c.JSON(http.StatusOK, gin.H{"message": "Upload sukses"})
}

func processImage(c *gin.Context, folder string, filename string) string {
	// Open file
	file, err := os.Open(folder + filename)
	if err != nil {
		throwMessage(c, "Gagal membuka gambar")
	}
	defer file.Close()

	// Decode image
	img, _, err := image.Decode(file)
	if err != nil {
		throwMessage(c, "Format gambar tidak didukung")
	}

	// Convert ke black and white
	result := convertToBlackAndWhite(img, 128)

	// Create new file to save result
	result_filename := filename + "_black_and_white"
	outFile, err := os.Create(result_filename)
	if err != nil {
		throwMessage(c, err.Error())
	}
	defer outFile.Close()

	// Save result to PNG
	err = png.Encode(outFile, result)
	if err != nil {
		throwMessage(c, "Gagal mengencoding gambar")
	}

	fmt.Println("Gambar berhasil diproses")
	return result_filename
}

func convertToBlackAndWhite(img image.Image, threshold uint8) *image.Gray {

}

func convertImageToBuffer(img image.Image) (*bytes.Buffer, error) {
	var buf bytes.Buffer
	err := png.Encode(&buf, img)
	if err != nil {
		return nil, err
	}

	return &buf, nil
}

func throwMessage(c *gin.Context, message string) {
	c.JSON(400, gin.H{"message": message})
	return
}

func main() {
	router := gin.Default()
	router.POST("/upload", postHandler)
	fmt.Println("http://localhost:8080")
	router.Run(":8080")
}
