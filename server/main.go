package main

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

func postHandler(c *gin.Context) {
	file, _ := c.FormFile("file")

	filename := file.Filename
	folder := "uploads/"
	c.SaveUploadedFile(file, folder+filename)

	buf := processImage(c, folder, filename)
	c.Data(http.StatusOK, "image/png", buf.Bytes())

	c.JSON(http.StatusOK, gin.H{"message": "Upload sukses"})
}

func processImage(c *gin.Context, folder string, filename string) *bytes.Buffer {
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
	buf, err := convertImageToBuffer(result)
	if err != nil {
		throwMessage(c, "Gagal mengconvert image ke buffer")
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
	return
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

	router.POST("/upload", postHandler)
	fmt.Println("http://localhost:8080")
	router.Run(":8080")
}
