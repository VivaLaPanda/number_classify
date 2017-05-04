package features

import "math"

// // Setting globals for white and black colors
// var white = color.Gray{Y: math.MaxUint8}
// var black = color.Gray{Y: 0}
//
// func preprocess(img image.Image) (retImg image.Gray) {
// 	// Doing some basic processing on the original image
// 	img = imaging.Grayscale(img)
// 	img = imaging.AdjustContrast(img, 30)
// 	img = imaging.Sharpen(img, 2)
//
// 	// Converting to Gray and binarizing
// 	grayImg := grayscale.Convert(img, grayscale.ToGrayLuma)
// 	threshold := grayscale.Otsu(grayImg)
// 	grayscale.Threshold(grayImg, threshold, 0, 255)
//
// 	// Get the bounds of the region the number is in
// 	imgSlice := grayImg.Bounds()
// 	var minPoint image.Point
// 	var maxPoint image.Point
//
// 	for y := imgSlice.Min.Y; y < imgSlice.Max.Y; y++ {
// 		for x := imgSlice.Min.X; x < imgSlice.Max.X; x++ {
// 			if grayImg.GrayAt(x, y) == white {
// 				if minPoint != image.ZP {
// 					minPoint = image.Point{X: x, Y: y}
// 				} else if maxPoint != image.ZP {
// 					maxPoint = image.Point{X: x, Y: y}
// 				}
// 			}
// 		}
// 	}
//
// 	// Crop around those bounds
// 	grayImg = grayImg.SubImage(image.Rectangle{Min: minPoint, Max: maxPoint}).(*image.Gray)
//
// 	return *grayImg
// }

// Density
func Density(imgArray [][]int) float64 {
	fgPx := 0.0
	bgPx := 0.0
	for _, row := range imgArray {
		for _, elem := range row {
			if elem == 0 {
				bgPx++
			} else {
				fgPx++
			}
		}
	}

	return fgPx / (bgPx + fgPx)
}

func VertSymmetry(imgArray [][]int) float64 {
	// Get total number of pixels
	// Assumes that array isn't ragged
	totalPx := float64(len(imgArray) + len(imgArray[0]))
	vertSymmetricPx := 0.0

	for i, row := range imgArray[:len(imgArray)/2] {
		for j, topElement := range row {
			bottomElement := imgArray[(len(imgArray)/2)+i][j]

			if topElement == bottomElement {
				vertSymmetricPx++
			}
		}
	}

	return vertSymmetricPx / totalPx
}

func HorizontalSymmetry(imgArray [][]int) float64 {
	// Get total number of pixels
	// Assumes that array isn't ragged
	totalPx := float64(len(imgArray) + len(imgArray[0]))
	horizontalSymmetricPx := 0.0

	for _, row := range imgArray {
		for j, leftElement := range row[:len(row)/2] {
			rightElement := row[(len(row)/2)+j]

			if leftElement == rightElement {
				horizontalSymmetricPx++
			}
		}
	}

	return horizontalSymmetricPx / totalPx
}

func HorizontalIntercepts(imgArray [][]int) (minIntercepts int, maxIntercepts int) {
	minIntercepts = math.MaxInt64
	maxIntercepts = math.MinInt64
	prevPx := 0

	for _, row := range imgArray {
		numIntercepts := 0

		for _, element := range row {
			if prevPx == 1 && element == 0 {
				numIntercepts++
			}

			prevPx = element
		}

		if numIntercepts < minIntercepts && numIntercepts != 0 {
			minIntercepts = numIntercepts
		}

		if numIntercepts > maxIntercepts {
			maxIntercepts = numIntercepts
		}
	}

	return minIntercepts, maxIntercepts
}

func VertIntercepts(imgArray [][]int) (minIntercepts int, maxIntercepts int) {
	minIntercepts = math.MaxInt64
	maxIntercepts = math.MinInt64
	prevPx := 0

	for i, row := range imgArray {
		numIntercepts := 0

		for j, _ := range row {
			if prevPx == 1 && imgArray[j][i] == 0 {
				numIntercepts++
			}

			prevPx = imgArray[j][i]
		}

		if numIntercepts < minIntercepts && numIntercepts != 0 {
			minIntercepts = numIntercepts
		}

		if numIntercepts > maxIntercepts {
			maxIntercepts = numIntercepts
		}
	}

	return minIntercepts, maxIntercepts
}
