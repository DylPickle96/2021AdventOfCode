package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type lowPoint struct {
	x int
	y int
}

func main() {
	contents, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Printf("cannot read input.txt. Error: %v", err)
		return
	}
	dataSet := makeDataSet(strings.Split(string(contents), "\n"))
	lowPoints := checkHeatMap(dataSet)
	determineBasins(lowPoints, dataSet)
}

func determineBasins(lowPoints []lowPoint, dataSet [][]int64) {
	basinSizes := []int{}
	for _, l := range lowPoints {
		returnPoints := []lowPoint{l}
		checkedLowPoints := []lowPoint{}
		basinSize := 1
		findBasinSize(dataSet, l, returnPoints, checkedLowPoints, &basinSize, true)
		fmt.Println("basinSize:", basinSize)
		basinSizes = append(basinSizes, basinSize)
	}
	findTheThreeLargestBasins(basinSizes)
}

func findTheThreeLargestBasins(basinSizes []int) {
	var first, second, third int
	for _, size := range basinSizes {
		if size > first {
			third = second
			second = first
			first = size
		} else if size > second {
			third = second
			second = size
		} else if size > third {
			third = size
		}
	}
	fmt.Printf("largest basins. First: %d. Second: %d. Third: %d\n", first, second, third)
	fmt.Println("Product of the largest basins:", first*second*third)
}

func findBasinSize(dataSet [][]int64, currentLowPoint lowPoint, returnPoints, checkedLowPoints []lowPoint, basinSize *int, firstRun bool) {
	fmt.Println("value of low point:", dataSet[currentLowPoint.y][currentLowPoint.x])
	fmt.Printf("currentLowPoint: %+v\n", currentLowPoint)
	fmt.Println("checkedLowPoints:", checkedLowPoints)
	// sanity check that we don't cause an infinite loop and out of bound check
	if currentLowPoint.x-1 >= 0 && !haveWeCheckedThisLowPoint(checkedLowPoints, lowPoint{x: currentLowPoint.x - 1, y: currentLowPoint.y}) {
		// left and it does not equal 9
		if dataSet[currentLowPoint.y][currentLowPoint.x-1] != 9 && dataSet[currentLowPoint.y][currentLowPoint.x-1] >= dataSet[currentLowPoint.y][currentLowPoint.x] {
			*basinSize += 1
			if !firstRun {
				returnPoints = append(returnPoints, currentLowPoint)
			}
			checkedLowPoints = append(checkedLowPoints, currentLowPoint)
			findBasinSize(dataSet, lowPoint{
				x: currentLowPoint.x - 1,
				y: currentLowPoint.y,
			}, returnPoints, checkedLowPoints, basinSize, false)
			return
		}
	}
	// sanity check that we don't cause an infinite loop and out of bound check
	if currentLowPoint.x+1 < len(dataSet[currentLowPoint.y]) && !haveWeCheckedThisLowPoint(checkedLowPoints, lowPoint{x: currentLowPoint.x + 1, y: currentLowPoint.y}) {
		// right and it does not equal 9
		if dataSet[currentLowPoint.y][currentLowPoint.x+1] != 9 && dataSet[currentLowPoint.y][currentLowPoint.x+1] >= dataSet[currentLowPoint.y][currentLowPoint.x] {
			*basinSize += 1
			if !firstRun {
				returnPoints = append(returnPoints, currentLowPoint)
			}
			checkedLowPoints = append(checkedLowPoints, currentLowPoint)
			findBasinSize(dataSet, lowPoint{
				x: currentLowPoint.x + 1,
				y: currentLowPoint.y,
			}, returnPoints, checkedLowPoints, basinSize, false)
			return
		}
	}
	// sanity check that we don't cause an infinite loop and out bound check
	if currentLowPoint.y-1 >= 0 && !haveWeCheckedThisLowPoint(checkedLowPoints, lowPoint{x: currentLowPoint.x, y: currentLowPoint.y - 1}) {
		// up and it does not equal 9
		if dataSet[currentLowPoint.y-1][currentLowPoint.x] != 9 && dataSet[currentLowPoint.y-1][currentLowPoint.x] >= dataSet[currentLowPoint.y][currentLowPoint.x] {
			*basinSize += 1
			if !firstRun {
				returnPoints = append(returnPoints, currentLowPoint)
			}
			checkedLowPoints = append(checkedLowPoints, currentLowPoint)
			findBasinSize(dataSet, lowPoint{
				x: currentLowPoint.x,
				y: currentLowPoint.y - 1,
			}, returnPoints, checkedLowPoints, basinSize, false)
			return
		}
	}
	// sanity check that we don't cause an infinite loop and out bound check
	if currentLowPoint.y+1 < len(dataSet) && !haveWeCheckedThisLowPoint(checkedLowPoints, lowPoint{currentLowPoint.x, currentLowPoint.y + 1}) {
		// down and it does not equal 9
		if dataSet[currentLowPoint.y+1][currentLowPoint.x] != 9 && dataSet[currentLowPoint.y+1][currentLowPoint.x] >= dataSet[currentLowPoint.y][currentLowPoint.x] {
			*basinSize += 1
			if !firstRun {
				returnPoints = append(returnPoints, currentLowPoint)
			}
			checkedLowPoints = append(checkedLowPoints, currentLowPoint)
			findBasinSize(dataSet, lowPoint{
				x: currentLowPoint.x,
				y: currentLowPoint.y + 1,
			}, returnPoints, checkedLowPoints, basinSize, false)
			return
		}
	}
	// exit condition
	if len(returnPoints) == 0 {
		fmt.Println("returning")
		return
	}
	returnLowPoint := returnPoints[len(returnPoints)-1]
	fmt.Println("returnLowPoint:", returnLowPoint)
	// slice delete to preserve order
	copy(returnPoints[len(returnPoints)-1:], returnPoints[len(returnPoints):])
	returnPoints[len(returnPoints)-1] = lowPoint{}
	returnPoints = returnPoints[:len(returnPoints)-1]
	checkedLowPoints = append(checkedLowPoints, currentLowPoint)
	fmt.Println("returnPoints:", returnPoints)
	fmt.Println()
	findBasinSize(dataSet, returnLowPoint, returnPoints, checkedLowPoints, basinSize, false)
}

func haveWeCheckedThisLowPoint(checkedLowPoints []lowPoint, l lowPoint) bool {
	for _, checkedLowPoint := range checkedLowPoints {
		if checkedLowPoint.x == l.x && checkedLowPoint.y == l.y {
			return true
		}
	}
	return false
}
func checkHeatMap(heatmap [][]int64) []lowPoint {
	lowPoints := []lowPoint{}
	for i := 0; i < len(heatmap); i++ {
		for j := 0; j < len(heatmap[i]); j++ {
			// out of bound check
			if i-1 >= 0 {
				// if the current value is larger than above it continue
				if heatmap[i][j] >= heatmap[i-1][j] {
					continue
				}
			}
			// out of bound check
			if i+1 < len(heatmap) {
				if heatmap[i][j] >= heatmap[i+1][j] {
					continue
				}
			}
			if j-1 >= 0 {
				if heatmap[i][j] >= heatmap[i][j-1] {
					continue
				}
			}
			if j+1 < len(heatmap[i]) {
				if heatmap[i][j] >= heatmap[i][j+1] {
					continue
				}
			}
			lowPoints = append(lowPoints, lowPoint{
				x: j,
				y: i,
			})
		}
	}
	return lowPoints
}

func makeDataSet(lines []string) [][]int64 {
	heatmap := [][]int64{}
	for _, line := range lines {
		values := strings.Split(line, "")
		heatmapLine := []int64{}
		for _, value := range values {
			v, _ := strconv.ParseInt(value, 10, 64)
			heatmapLine = append(heatmapLine, v)
		}
		heatmap = append(heatmap, heatmapLine)
	}
	return heatmap
}
