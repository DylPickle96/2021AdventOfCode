package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	contents, err := os.ReadFile("input.txt")
	if err != nil {
		log.Printf("cannot read input.txt. Error: %v", err)
		return
	}
	lines := strings.Split(string(contents), "\n")
	total := 0
	// for i := 1; i < len(lines); i++ {
	// 	current, err := strconv.ParseInt(lines[i], 10, 64)
	// 	if err != nil {
	// 		log.Printf("cannot parse line. Error: %v", err)
	// 		break
	// 	}
	// 	previous, err := strconv.ParseInt(lines[i-1], 10, 64)
	// 	if err != nil {
	// 		log.Printf("cannot parse line. Error: %v", err)
	// 		break
	// 	}
	// 	if current > previous {
	// 		total += 1
	// 	}
	// }

	for index, _ := range lines {
		cSet, currentTotal := countForward(lines, int64(index))
		nSet, nextTotal := countForward(lines, int64(index+1))
		if currentTotal < nextTotal {
			total += 1
		}
		fmt.Println("current set: ", cSet, "current Total: ", currentTotal, "next set: ", nSet, "next total: ", nextTotal, "total: ", total)
		fmt.Println(index+4, len(lines))
		if index+4 == len(lines) {
			break
		}
	}
	log.Println(total)
}

func countForward(lines []string, index int64) ([]int64, int64) {
	total := int64(0)
	set := []int64{}
	for i := index; i <= index+2; i++ {
		current, err := strconv.ParseInt(lines[i], 10, 64)
		if err != nil {
			log.Printf("cannot parse line. Error: %v", err)
			break
		}
		set = append(set, current)
		total += current

	}
	return set, total
}
