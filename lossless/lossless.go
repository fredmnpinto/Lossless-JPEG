package lossless

import "image"

// Encode writes the image encoded in Lossless JPEG to the file
func Encode(image image.YCbCr, filename string) (err error) {

	return nil
}

// Decode reads the image encoded in Lossless JPEG from the file
func Decode(image image.YCbCr, filename string) (decodedImage image.YCbCr, err error) {

	return decodedImage, nil
}
