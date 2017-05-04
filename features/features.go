package features

import (
	"github.com/disintegration/imaging"
	"github.com/harrydb/go/img/grayscale"
	"image"
	"image/color"
	"math"
)

// Setting globals for white and black colors
var white = color.Gray{Y: math.MaxUint8}
var black = color.Gray{Y: 0}

func preprocess(img image.Image) (retImg image.Gray) {
	// Doing some basic processing on the original image
	img = imaging.Grayscale(img)
	img = imaging.AdjustContrast(img, 30)
	img = imaging.Sharpen(img, 2)

	// Converting to Gray and binarizing
	grayImg := grayscale.Convert(img, grayscale.ToGrayLuma)
	threshold := grayscale.Otsu(grayImg)
	grayscale.Threshold(grayImg, threshold, 0, 255)

	// Get the bounds of the region the number is in
	imgSlice := grayImg.Bounds()
	var minPoint image.Point
	var maxPoint image.Point

	for y := imgSlice.Min.Y; y < imgSlice.Max.Y; y++ {
		for x := imgSlice.Min.X; x < imgSlice.Max.X; x++ {
			if grayImg.GrayAt(x, y) == white {
				if minPoint != image.ZP {
					minPoint = image.Point{X: x, Y: y}
				} else if maxPoint != image.ZP {
					maxPoint = image.Point{X: x, Y: y}
				}
			}
		}
	}

	// Crop around those bounds
	grayImg = grayImg.SubImage(image.Rectangle{Min: minPoint, Max: maxPoint}).(*image.Gray)

	return *grayImg
}

// Density
func Density(inImg image.Image) int {
	img := preprocess(inImg)

	return 0
}
