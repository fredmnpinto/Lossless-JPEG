package main

import (
	"image"
	"lossless_jpeg/lossy"
)

func errCheck(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	img, err := lossy.DecodeJpeg("input.jpeg")
	errCheck(err)

	ycbcr := convertRGBToYCbCr(img)

	err = lossy.EncodeJpeg(ycbcr, "output.jpeg")
	errCheck(err)
}

func convertRGBToYCbCr(img image.Image) *image.YCbCr {
	bounds := img.Bounds()
	ycbcr := image.NewYCbCr(bounds, image.YCbCrSubsampleRatio420)
	for y := 0; y < bounds.Max.Y; y++ {
		for x := 0; x < bounds.Max.X; x++ {
			r, g, b, _ := img.At(x, y).RGBA()
			ycbcr.Y[y*ycbcr.YStride+x] = uint8(0.299*float64(r>>8) + 0.587*float64(g>>8) + 0.114*float64(b>>8))
			ycbcr.Cb[(y/2)*ycbcr.CStride+x/2] = uint8(-0.169*float64(r>>8) - 0.331*float64(g>>8) + 0.5*float64(b>>8) + 128)
			ycbcr.Cr[(y/2)*ycbcr.CStride+x/2] = uint8(0.5*float64(r>>8) - 0.419*float64(g>>8) - 0.081*float64(b>>8) + 128)
		}
	}
	return ycbcr
}
