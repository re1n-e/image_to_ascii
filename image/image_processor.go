package image

import (
	"fmt"
	"image"
)

func loadImageAsMatrix(img image.Image) [][][3]uint8 {
	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y

	matrix := make([][][3]uint8, height)
	for y := 0; y < height; y++ {
		row := make([][3]uint8, width)
		for x := 0; x < width; x++ {
			r, g, b, _ := img.At(x, y).RGBA()
			row[x] = [3]uint8{
				uint8(r >> 8),
				uint8(g >> 8),
				uint8(b >> 8),
			}
		}
		matrix[y] = row
	}
	return matrix
}

func createBrightnessMatrix(matrix [][][3]uint8) [][]uint8 {
	height := len(matrix)
	width := len(matrix[0])
	brightness := make([][]uint8, height)

	for y := 0; y < height; y++ {
		row := make([]uint8, width)
		for x := 0; x < width; x++ {
			r, g, b := float32(matrix[y][x][0]), float32(matrix[y][x][1]), float32(matrix[y][x][2])
			row[x] = uint8(r*0.299 + g*0.587 + b*0.114)
		}
		brightness[y] = row
	}

	return brightness
}

func imageToAscii(matrix [][]uint8) {
	asciiChars := "`^\",:;Il!i~+_-?][}{1)(|\\/tfjrxnuvczXYUJCLQ0OZmwqpdbkhao*#MW&8%B@$"
	levels := len(asciiChars)

	for _, row := range matrix {
		for _, val := range row {
			index := int(val) * (levels - 1) / 255 // Scale 0–255 to 0–len(asciiChars)-1
			fmt.Print(string(asciiChars[index]))
		}
		fmt.Println()
	}
}
