package main

import (
	"log"
	"os"
	"strings"
)

func main() {
	contents, err := os.ReadFile("input.txt")
	if err != nil {
		log.Printf("cannot read input.txt. Error: %v", err)
		return
	}
	lines := strings.Split(string(contents), "\n")
	bitsValues := [][]string{}
	for _, line := range lines {
		bitsValues = append(bitsValues, strings.Split(line, ""))
	}
	lengthOfBitValue := len(bitsValues[0])

	for i := 0; i < lengthOfBitValue; i++ {
		zero := 0
		one := 0
		for _, value := range bitsValues {
			if value[i] == "0" {
				zero += 1
			}
			if value[i] == "1" {
				one += 1
			}
		}
		if zero > one {
			bitsValues = reduceBitValues(bitsValues, "one", i)
		} else if one > zero {
			bitsValues = reduceBitValues(bitsValues, "zero", i)
		} else if one == zero {
			bitsValues = reduceBitValues(bitsValues, "zero", i)
		}
		if len(bitsValues) == 1 {
			log.Println(bitsValues)
		}
	}

	// tracker := make(map[int]map[string]int64)
	// for _, bit := range bits {
	// 	for i := 0; i < lengthOfBitValue; i++ {
	// 		if tracker[i] == nil {
	// 			tracker[i] = make(map[string]int64)
	// 		}
	// 		if bit[i] == "0" {
	// 			tracker[i]["zero"] = tracker[i]["zero"] + 1
	// 		}
	// 		if bit[i] == "1" {
	// 			tracker[i]["one"] = tracker[i]["one"] + 1
	// 		}
	// 	}
	// }

}

func reduceBitValues(bitsValue [][]string, greatestBit string, index int) [][]string {
	reducedBitsValue := [][]string{}
	for _, value := range bitsValue {
		if greatestBit == "zero" && value[index] == "0" {
			reducedBitsValue = append(reducedBitsValue, value)
		}
		if greatestBit == "one" && value[index] == "1" {
			reducedBitsValue = append(reducedBitsValue, value)
		}
	}
	return reducedBitsValue
}
