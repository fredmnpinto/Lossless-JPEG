package lossy

import (
	"image"
	"image/jpeg"
	"os"
)

// EncodeJpeg writes the given image encoded in regular JPEG to the given file.
func EncodeJpeg(ycbcr *image.YCbCr, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	err = jpeg.Encode(file, ycbcr, &jpeg.Options{Quality: 100})
	if err != nil {
		return err
	}

	return nil
}

// DecodeJpeg reads a regular JPEG image from the file
func DecodeJpeg(filename string) (image.Image, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	img, err := jpeg.Decode(file)
	if err != nil {
		return nil, err
	}

	return img, nil
}
