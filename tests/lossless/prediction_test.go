package lossless_test

import (
	"lossless_jpeg/lossless"
	"testing"
)

func TestGetOffSetCategory_ReturnsCorrectCategory(t *testing.T) {
	offsetMock := lossless.PixelOffset{
		Y:  0,
		Cb: 32767,
		Cr: 6000,
	}

	offsetCategories := offsetMock.GetOffsetCategory()

	if offsetCategories.Y != 0 {
		t.Errorf("GetOffsetCategory() returned %v instead of 0", offsetCategories.Y)
	}

	if offsetCategories.Cb != 15 {
		t.Errorf("GetOffsetCategory() returned %v instead of %v", offsetCategories.Cb, 15)
	}

	if offsetCategories.Cr != 13 {
		t.Errorf("GetOffsetCategory() returned %v instead of %v", offsetCategories.Cr, 13)
	}
}
