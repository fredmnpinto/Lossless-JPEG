package lossless

import (
	"image"
	"image/color"
	"log"
	"math"
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

type PixelOffset struct {
	Y  int16
	Cb int16
	Cr int16
}

type PixelOffsetCategories struct {
	Y  uint8
	Cb uint8
	Cr uint8
}

// CalculatePredictionOffsets calculates the offsets from the predicted pixel
// and the actual pixel for each component of each pixel
func CalculatePredictionOffsets(image image.YCbCr) ([][]PixelOffset, [][]PixelOffsetCategories) {
	var offsets [][]PixelOffset
	var offsetCategories [][]PixelOffsetCategories

	for y := 0; y < image.Rect.Max.Y; y++ {
		row := make([]PixelOffset, image.Rect.Max.X)
		categoriesRow := make([]PixelOffsetCategories, image.Rect.Max.X)

		for x := 0; x < image.Rect.Max.X; x++ {
			pixel := image.YCbCrAt(x, y)

			predicted := predictPixel(getRelevantNeighbors(image, x, y))

			pixelOffset := PixelOffset{
				Y:  int16(pixel.Y) - int16(predicted.Y),
				Cb: int16(pixel.Cb) - int16(predicted.Cb),
				Cr: int16(pixel.Cr) - int16(predicted.Cr),
			}

			row = append(row, pixelOffset)
			categoriesRow = append(categoriesRow, PixelOffsetCategories{
				Y:  getOffsetCategory(pixelOffset.Y),
				Cb: getOffsetCategory(pixelOffset.Cb),
				Cr: getOffsetCategory(pixelOffset.Cr),
			})
		}

		offsets = append(offsets, row)
		offsetCategories = append(offsetCategories, categoriesRow)
	}
	return offsets, offsetCategories
}

// getRelevantNeighbors returns the neighbors relevant to the predictor for given pixel position
func getRelevantNeighbors(image image.YCbCr, x int, y int) (neighborA color.YCbCr, neighborB color.YCbCr, neighborC color.YCbCr) {
	neighborA = image.YCbCrAt(x-1, y)
	neighborB = image.YCbCrAt(x, y-1)
	neighborC = image.YCbCrAt(x-1, y-1)

	return neighborA, neighborB, neighborC
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

// getOffsetCategory returns the offset category (from 0 to 16) based on the offset (from 0 to 32768)
func getOffsetCategory(offset int16) uint8 {
	if offset == 0 {
		return 0
	}

	category := uint8(math.Log2(float64(offset))) + 1

	return category
}

func (offset *PixelOffset) GetOffsetCategory() PixelOffsetCategories {
	yCategory := getOffsetCategory(offset.Y)
	cbCategory := getOffsetCategory(offset.Cb)
	crCategory := getOffsetCategory(offset.Cr)

	return PixelOffsetCategories{
		Y:  yCategory,
		Cb: cbCategory,
		Cr: crCategory,
	}
}
