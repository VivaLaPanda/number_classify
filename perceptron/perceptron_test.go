package perceptron

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"testing"
)

func TestAll(t *testing.T) {
	perceptron := New(8, .1)

	trainArr, trainExptVal := parseFile("../train_data.txt")
	testArr, testExptVal := parseFile("../test_data.txt")

	perceptron = perceptron.Train(trainArr, trainExptVal)

	accuracy := perceptron.Test(testArr, testExptVal)
	fmt.Printf("Classified inputs with %% %v accuracy\n", accuracy*100)
}

func parseFile(filename string) (imgArrays [][][]int, expectedValues []int) {
	// Open the file
	file, err := os.Open(filename) // just pass the file name
	if err != nil {
		check(err)
	}
	defer file.Close()

	// Initailize the array of image bitmaps
	imgArrays = [][][]int{}

	// Start scanning the file line by line
	scanner := bufio.NewScanner(file)
	tempImgArray := [][]int{}
	tempExpectedValue := 0
	for scanner.Scan() {
		// Read in one line of the file
		tempLine := scanner.Text()

		// If we have a blank line then we finished a block of stuff
		// Save and scan until we get to then next block
		if tempLine == "" {
			imgArrays = append(imgArrays, tempImgArray)
			expectedValues = append(expectedValues, tempExpectedValue)

			tempImgArray = [][]int{}
			tempExpectedValue = 0

			// Skip one line and then go back to the beginning of the loop
			scanner.Scan()
			continue
		}

		// COnvert the row into an array of ints
		tempStrArray := strings.Split(tempLine, " ")

		// Line starts with a space, skip it
		if tempStrArray[0] == "" {
			tempStrArray = tempStrArray[1:]
		}

		// Convert the row to ints
		rowInts := make([]int, len(tempStrArray))
		for i, elem := range tempStrArray {
			rowInts[i], err = strconv.Atoi(elem)

			if err != nil {
				panic(fmt.Errorf("Failed to parse char into int: %v\n", elem))
			}
		}

		// We have 29 elements, we need to grab the expected value from the row
		if len(rowInts) == 29 {
			if rowInts[0] == 1 {
				tempExpectedValue = -1
			} else if rowInts[0] == 5 {
				tempExpectedValue = 1
			} else {
				panic(fmt.Errorf("Trying to identify number besides 1 and 5: %v\n", rowInts[0]))
			}

			rowInts = rowInts[1:]
		}

		// Ad this row to the current bitmap
		tempImgArray = append(tempImgArray, rowInts)
	}

	if err := scanner.Err(); err != nil {
		check(err)
	}

	return imgArrays, expectedValues
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
