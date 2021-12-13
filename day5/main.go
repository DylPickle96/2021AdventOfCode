package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// represents the raw data from the input file
type dataSet struct {
	lines []line
}

type line struct {
	start dot
	end   dot
}

type dot struct {
	x int64
	y int64
}

func main() {
	contents, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Printf("cannot read input.txt. Error: %v", err)
		return
	}
	lines := strings.Split(string(contents), "\n")
	dataSet := createDataSet(lines)
	createLineMap(dataSet)
}

func createDataSet(lines []string) dataSet {
	dataSet := dataSet{}
	for _, l := range lines {
		sp := strings.Split(l, " ")
		line := line{}
		for i, s := range sp {
			p := strings.Split(s, ",")
			x, _ := strconv.ParseInt(p[0], 10, 64)
			y, _ := strconv.ParseInt(p[1], 10, 64)
			dot := dot{
				x: x,
				y: y,
			}
			if i == 0 {
				line.start = dot
			} else if i == 1 {
				line.end = dot
			}
		}
		dataSet.lines = append(dataSet.lines, line)
	}
	return dataSet
}

func createLineMap(dataSet dataSet) {
	lineData := make(map[int64]map[int64]int64)
	for _, line := range dataSet.lines {
		// we only care about horizontal lines
		if line.start.x == line.end.x {
			if line.start.y < line.end.y {
				for i := line.start.y; i <= line.end.y; i++ {
					if lineData[line.start.x] == nil {
						lineData[line.start.x] = make(map[int64]int64)
					}
					lineData[line.start.x][i] += 1
				}
			} else {
				for i := line.end.y; i <= line.start.y; i++ {
					if lineData[line.start.x] == nil {
						lineData[line.start.x] = make(map[int64]int64)
					}
					lineData[line.start.x][i] += 1
				}
			}
			// we only care about horizontal lines
		} else if line.start.y == line.end.y {
			if line.start.x < line.end.x {
				for i := line.start.x; i <= line.end.x; i++ {
					if lineData[i] == nil {
						lineData[i] = make(map[int64]int64)
					}
					lineData[i][line.start.y] += 1
				}
			} else {
				for i := line.end.x; i <= line.start.x; i++ {
					if lineData[i] == nil {
						lineData[i] = make(map[int64]int64)
					}
					lineData[i][line.start.y] += 1
				}
			}
			// we only care about diagonal lines
		} else {
			// these values track how far we are from the orginial position
			var x, y int64
			for {
				// check if map is nil before assignment
				if lineData[line.start.x+x] == nil {
					lineData[line.start.x+x] = make(map[int64]int64)
				}
				// mark the data point
				lineData[line.start.x+x][line.start.y+y] += 1
				// check our exit condition
				if line.start.x+x == line.end.x && line.start.y+y == line.end.y {
					break
				}
				// add or subtract for the next cycle
				if line.start.x < line.end.x {
					x += 1
				} else {
					x -= 1
				}
				// add or subtract for the next cycle
				if line.start.y < line.end.y {
					y += 1
				} else {
					y -= 1
				}

			}
		}
	}
	fmt.Println(lineData)
	var overlapTotal int64
	for _, yValues := range lineData {
		for _, v := range yValues {
			if v > 1 {
				overlapTotal += 1
			}
		}
	}
	log.Println(overlapTotal)
}
