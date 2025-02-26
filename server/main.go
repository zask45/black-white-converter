package main

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"path/filepath"

	"log"
	"os"

	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func postHandler(c *gin.Context) {
	file, _ := c.FormFile("file")

	if file == nil {
		log.Fatalf("file nil")
		return
	}

	os.MkdirAll("uploads", os.ModePerm) // Membuat folder jika belum ada
	filepath := filepath.Join("uploads", file.Filename)
	c.SaveUploadedFile(file, filepath)

	buf := processImage(filepath)
	if buf == nil {
		log.Fatalf("gagal memproses gambar")
	}
	c.Data(http.StatusOK, "image/png", buf.Bytes())

	c.JSON(http.StatusOK, gin.H{"message": "Upload sukses"})
}

func processImage(path string) *bytes.Buffer {
	// Open file
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("Gagal membuka file: %v", err)
	}
	defer file.Close()

	// Decode image
	img, _, _ := image.Decode(file)
	if img == nil {
		log.Fatalf("img nil")
	}

	// Convert ke black and white
	result := convertToBlackAndWhite(img, 128)

	buf, _ := convertImageToBuffer(result)
	if buf == nil {
		log.Fatal("Gagal convert image ke black and white")
	}

	return buf

	// // Create new file to save result
	// result_filename := filename + "_black_and_white"
	// outFile, err := os.Create(result_filename)
	// if err != nil {
	// 	throwMessage(c, err.Error())
	// }
	// defer outFile.Close()

	// // Save result to PNG
	// err = png.Encode(outFile, result)
	// if err != nil {
	// 	throwMessage(c, "Gagal mengencoding gambar")
	// }

	// fmt.Println("Gambar berhasil diproses")
	// return result_filename
	// return nil
}

func convertToBlackAndWhite(img image.Image, threshold uint8) *image.Gray {
	bounds := img.Bounds()
	bwImg := image.NewGray(bounds)

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			if color.GrayModel.Convert(img.At(x, y)).(color.Gray).Y > threshold {
				bwImg.Set(x, y, color.Gray{255}) // Putih
			} else {
				bwImg.Set(x, y, color.Gray{0}) // Hitam
			}
		}
	}

	return bwImg
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
}

func main() {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"POST", "GET", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// router.GET("/", func(c *gin.Context) {
	// 	c.String(http.StatusOK, "Server berjalan!")
	// })

	router.POST("/upload", postHandler)
	fmt.Println("http://localhost:8080")
	router.Run(":8080")
}
