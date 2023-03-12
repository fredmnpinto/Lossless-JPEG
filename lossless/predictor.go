package lossless

import (
	"image"
	"image/color"
	"log"
	"os"
	"strconv"
)

var predictors = []func(a uint8, b uint8, c uint8) uint8{
	func(a uint8, b uint8, c uint8) uint8 { return 0 },
	func(a uint8, b uint8, c uint8) uint8 { return a },
	func(a uint8, b uint8, c uint8) uint8 { return b },
	func(a uint8, b uint8, c uint8) uint8 { return c },
	func(a uint8, b uint8, c uint8) uint8 { return a + b - c },
	func(a uint8, b uint8, c uint8) uint8 { return a + (b-c)/2 },
	func(a uint8, b uint8, c uint8) uint8 { return b + (a-c)/2 }, // 6
	func(a uint8, b uint8, c uint8) uint8 { return (a + b) / 2 },
}

const defaultPredictor = 6 // This is the most commonly used predictor

// CalculateOffsets calculates the offsets from the predicted pixel
// and the actual pixel for each component of each pixel
func CalculateOffsets(image image.YCbCr) [][]color.YCbCr {
	var offsets [][]color.YCbCr
	for y := 0; y < image.Rect.Max.Y; y++ {
		row := make([]color.YCbCr, image.Rect.Max.X)
		for x := 0; x < image.Rect.Max.X; x++ {
			pixel := image.YCbCrAt(x, y)
			neighbor_a := image.YCbCrAt(x-1, y)
			neighbor_b := image.YCbCrAt(x, y-1)
			neighbor_c := image.YCbCrAt(x-1, y-1)

			predicted := predictPixel(neighbor_a, neighbor_b, neighbor_c)

			pixel_offset := color.YCbCr{
				Y:  pixel.Y - predicted.Y,
				Cb: pixel.Cb - predicted.Cb,
				Cr: pixel.Cr - predicted.Cr,
			}

			row = append(row, pixel_offset)
		}
		offsets = append(offsets, row)
	}
	return offsets
}

// predictPixel uses a DPCM predictor to predict the pixel based on it's neighbors
func predictPixel(a color.YCbCr, b color.YCbCr, c color.YCbCr) color.YCbCr {

	predictor := predictors[predictorToUse()]

	y := predictor(a.Y, b.Y, c.Y)
	cb := predictor(a.Cb, b.Cb, c.Cb)
	cr := predictor(a.Cr, b.Cr, c.Cr)

	return color.YCbCr{Y: y, Cb: cb, Cr: cr}
}

// predictorToUse returns the predictor number to be used
func predictorToUse() int {
	numStr := os.Getenv("PREDICTOR")

	if numStr == "" {
		return defaultPredictor
	}

	// Convert the string to an int
	num, err := strconv.Atoi(numStr)
	if err != nil {
		log.Println("Error:", err)
		return defaultPredictor
	}

	// Print the numeric value
	log.Println("Using predictor:", num)

	return num
}
