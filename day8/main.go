package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var correctDigitScore = make(map[int]int)

func init() {
	correctDigitScore[42] = 0
	correctDigitScore[17] = 1
	correctDigitScore[34] = 2
	correctDigitScore[39] = 3
	correctDigitScore[30] = 4
	correctDigitScore[37] = 5
	correctDigitScore[41] = 6
	correctDigitScore[25] = 7
	correctDigitScore[49] = 8
	correctDigitScore[45] = 9
}

func main() {
	contents, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Printf("cannot read input.txt. Error: %v", err)
		return
	}
	lines := strings.Split(string(contents), "\n")
	sum := int64(0)
	for _, line := range lines {
		inputOutput := strings.Split(line, "|")
		inputDigits := strings.Split(inputOutput[0], " ")
		currentDigitsScore := make(map[string]int)
		for _, digit := range inputDigits {
			segments := strings.Split(digit, "")
			for _, segment := range segments {
				currentDigitsScore[segment] += 1
			}
		}
		outDigits := strings.Split(inputOutput[1], " ")
		outputValue := ""
		for _, digit := range outDigits {
			segments := strings.Split(digit, "")
			scoreForDigit := 0
			for _, segment := range segments {
				scoreForDigit += currentDigitsScore[segment]
			}
			outputValue = outputValue + strconv.FormatInt(int64(correctDigitScore[scoreForDigit]), 10)
		}
		s, _ := strconv.ParseInt(outputValue, 10, 64)
		sum += s
	}
	fmt.Println(sum)
}

// // a flat structure which defines the correct positions of digits in a working submarine
// var correctPositions = make(map[int]string)
// var newPositions = make(map[int]string)
// var knownPositionsOfNumbers = make(map[int][]int)

// func init() {
// 	correctPositions[0] = "a"
// 	correctPositions[1] = "b"
// 	correctPositions[2] = "c"
// 	correctPositions[3] = "d"
// 	correctPositions[4] = "e"
// 	correctPositions[5] = "f"
// 	correctPositions[6] = "g"

// 	knownPositionsOfNumbers[0] = []int{0, 1, 2, 4, 5, 6}
// 	knownPositionsOfNumbers[1] = []int{2, 5}
// 	knownPositionsOfNumbers[2] = []int{0, 2, 3, 4, 6}
// 	knownPositionsOfNumbers[3] = []int{0, 2, 3, 5, 6}
// 	knownPositionsOfNumbers[4] = []int{1, 2, 3, 5}
// 	knownPositionsOfNumbers[5] = []int{0, 1, 3, 5, 6}
// 	knownPositionsOfNumbers[6] = []int{0, 1, 3, 4, 5, 6}
// 	knownPositionsOfNumbers[7] = []int{0, 2, 5}
// 	knownPositionsOfNumbers[8] = []int{0, 1, 2, 3, 4, 5, 6}
// 	knownPositionsOfNumbers[9] = []int{0, 1, 2, 3, 5, 6}
// }

// func main() {
// 	contents, err := os.ReadFile("input.txt")
// 	if err != nil {
// 		fmt.Printf("cannot read input.txt. Error: %v", err)
// 		return
// 	}
// 	lines := strings.Split(string(contents), "\n")
// 	sum := int64(0)
// 	for _, line := range lines {
// 		displayValues := strings.Split(line, "|")[0]
// 		outputValues := strings.Split(line, "|")[1]
// 		determineMapping(displayValues)
// 		mappingForNumbers := determineValue(displayValues)
// 		ret := calulateOutput(outputValues, mappingForNumbers)
// 		fmt.Println(ret)
// 	}
// 	fmt.Println(sum)
// }

// func determineValue(displaValues string) map[int]string {
// 	digits := strings.Split(displaValues, " ")
// 	var mappingForNumbers = make(map[int]string)
// 	for _, digit := range digits {
// 		positions := []int{}
// 		// look at each "segment" in the "segments"
// 		segments := strings.Split(digit, "")
// 		for _, segment := range segments {
// 			for key, position := range newPositions {
// 				if position == segment {
// 					positions = append(positions, key)
// 				}
// 			}
// 		}
// 		for key, n := range knownPositionsOfNumbers {
// 			if sameIntSlice(n, positions) {
// 				mappingForNumbers[key] = digit
// 				break
// 			}
// 		}
// 	}
// 	return mappingForNumbers
// }

// func calulateOutput(outputValues string, mappingForNumbers map[int]string) int64 {
// 	values := strings.Split(outputValues, " ")
// 	outputNumbers := []int{}
// 	for _, value := range values {
// 		splitValue := strings.Split(value, "")
// 		for key, mapping := range mappingForNumbers {
// 			m := strings.Split(mapping, "")
// 			if sameStringSlice(splitValue, m) {
// 				outputNumbers = append(outputNumbers, key)
// 			}
// 		}
// 	}
// 	stringNumber := ""
// 	for _, number := range outputNumbers {
// 		str := strconv.FormatInt(int64(number), 10)
// 		stringNumber = stringNumber + str
// 	}
// 	ret, _ := strconv.ParseInt(stringNumber, 10, 64)
// 	return ret
// }

// func sameIntSlice(x, y []int) bool {
// 	if len(x) != len(y) {
// 		return false
// 	}
// 	// create a map of string -> int
// 	diff := make(map[int]int, len(x))
// 	for _, _x := range x {
// 		// 0 value for int is 0, so just increment a counter for the string
// 		diff[_x]++
// 	}
// 	for _, _y := range y {
// 		// If the string _y is not in diff bail out early
// 		if _, ok := diff[_y]; !ok {
// 			return false
// 		}
// 		diff[_y] -= 1
// 		if diff[_y] == 0 {
// 			delete(diff, _y)
// 		}
// 	}
// 	return len(diff) == 0
// }

// func sameStringSlice(x, y []string) bool {
// 	if len(x) != len(y) {
// 		return false
// 	}
// 	// create a map of string -> int
// 	diff := make(map[string]int, len(x))
// 	for _, _x := range x {
// 		// 0 value for int is 0, so just increment a counter for the string
// 		diff[_x]++
// 	}
// 	for _, _y := range y {
// 		// If the string _y is not in diff bail out early
// 		if _, ok := diff[_y]; !ok {
// 			return false
// 		}
// 		diff[_y] -= 1
// 		if diff[_y] == 0 {
// 			delete(diff, _y)
// 		}
// 	}
// 	return len(diff) == 0
// }

// func determineMapping(displaValues string) {
// 	digits := strings.Split(displaValues, " ")
// 	findOne(digits)
// 	findSeven(digits)
// 	findFour(digits)
// 	findEight(digits)
// }

// func findOne(digits []string) {
// 	for _, digit := range digits {
// 		// number of segments which make up one
// 		if len(digit) == 2 {
// 			// assign the segments to the known positions of one
// 			segments := strings.Split(digit, "")
// 			if correctPositions[2] == segments[0] {
// 				newPositions[5] = segments[0]
// 				newPositions[2] = segments[1]
// 			} else {
// 				newPositions[2] = segments[0]
// 				newPositions[5] = segments[1]
// 			}
// 			break
// 		}
// 	}
// }

// func findSeven(digits []string) {
// 	for _, digit := range digits {
// 		// number of segments which make up seven
// 		if len(digit) == 3 {
// 			segments := strings.Split(digit, "")
// 			for _, segment := range segments {
// 				// seven will contain the two known segments from one
// 				// use this to determine the last segment postion of seven
// 				if segment != newPositions[2] && segment != newPositions[5] {
// 					newPositions[0] = segment
// 					break
// 				}
// 			}
// 			break
// 		}
// 	}
// }

// func findFour(digits []string) {
// 	for _, digit := range digits {
// 		// number of segments which make up four
// 		if len(digit) == 4 {
// 			segments := strings.Split(digit, "")
// 			newSegments := []string{}
// 			for _, segment := range segments {
// 				// four will contain the two known segments from one
// 				if segment != newPositions[2] && segment != newPositions[5] {
// 					newSegments = append(newSegments, segment)
// 				}
// 			}
// 			if correctPositions[1] == newSegments[0] {
// 				newPositions[1] = newSegments[1]
// 				newPositions[3] = newSegments[0]
// 			} else if correctPositions[1] == newSegments[1] {
// 				newPositions[1] = newSegments[0]
// 				newPositions[3] = newSegments[1]
// 			} else if correctPositions[3] == newSegments[0] {
// 				newPositions[1] = newSegments[0]
// 				newPositions[3] = newSegments[1]

// 			} else if correctPositions[3] == newSegments[1] {
// 				newPositions[1] = newSegments[1]
// 				newPositions[3] = newSegments[0]
// 			} else {
// 				// if either don't match just assign in linear order
// 				newPositions[1] = newSegments[1]
// 				newPositions[3] = newSegments[0]
// 			}
// 			break
// 		}
// 	}
// }

// func findEight(digits []string) {
// 	for _, digit := range digits {
// 		// number of segments which make up eight
// 		if len(digit) == 7 {
// 			segments := strings.Split(digit, "")
// 			newSegments := []string{}
// 			for _, segment := range segments {
// 				// eight will contain all 5 known positions
// 				if segment != newPositions[0] &&
// 					segment != newPositions[1] &&
// 					segment != newPositions[2] &&
// 					segment != newPositions[3] &&
// 					segment != newPositions[5] {
// 					newSegments = append(newSegments, segment)
// 				}
// 			}
// 			if correctPositions[4] == newSegments[0] {
// 				newPositions[4] = newSegments[1]
// 				newPositions[6] = newSegments[0]
// 			} else if correctPositions[4] == newSegments[1] {
// 				newPositions[4] = newSegments[0]
// 				newPositions[6] = newSegments[1]
// 			} else if correctPositions[6] == newSegments[0] {
// 				newPositions[4] = newSegments[0]
// 				newPositions[6] = newSegments[1]

// 			} else if correctPositions[6] == newSegments[1] {
// 				newPositions[4] = newSegments[1]
// 				newPositions[6] = newSegments[0]
// 			} else {
// 				// if either don't match just assign in linear order
// 				newPositions[4] = newSegments[1]
// 				newPositions[6] = newSegments[0]
// 			}
// 			break
// 		}

// 	}
// }

// // part 1
// func determineUniqueCharacters(inputDigits string) int {
// 	digits := strings.Split(inputDigits, " ")
// 	count := 0
// 	for _, digit := range digits {
// 		switch len(digit) {
// 		case 2:
// 			count += 1
// 		case 3:
// 			count += 1
// 		case 4:
// 			count += 1
// 		case 7:
// 			count += 1
// 		default:
// 			continue
// 		}
// 	}
// 	return count
// }
