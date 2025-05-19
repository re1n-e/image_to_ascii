package image

import (
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"os"

	"github.com/nfnt/resize"
)

func resizeImage(img image.Image, newWidth uint) image.Image {
	bounds := img.Bounds()
	ratio := float64(bounds.Dy()) / float64(bounds.Dx())
	newHeight := uint(float64(newWidth) * ratio)

	resized := resize.Resize(newWidth, newHeight, img, resize.Lanczos3)
	return resized
}

func validate_image(filePath string) image.Image {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		fmt.Printf("Invalid Image file: %v\n", err)
		os.Exit(1)
	}

	return img
}

func Image_processor() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage: %s <Image file>\n", os.Args[0])
		return
	}
	imageFile := os.Args[1]
	img := resizeImage(validate_image(imageFile), 50)
	MatrixImg := loadImageAsMatrix(img)
	BrightImg := createBrightnessMatrix(MatrixImg)
	imageToAscii(BrightImg)
}
