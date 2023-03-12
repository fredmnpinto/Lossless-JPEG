package lossless

import (
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

// DPCM predictor to predict the pixel based on it's neighbors
func predictPixel(a color.YCbCr, b color.YCbCr, c color.YCbCr) color.YCbCr {

	predictor := predictors[predictorToUse()]

	y := predictor(a.Y, b.Y, c.Y)
	cb := predictor(a.Cb, b.Cb, c.Cb)
	cr := predictor(a.Cr, b.Cr, c.Cr)

	return color.YCbCr{Y: y, Cb: cb, Cr: cr}
}

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
