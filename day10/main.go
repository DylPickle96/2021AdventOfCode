package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	contents, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Printf("cannot read input.txt. Error: %v", err)
		return
	}
	lines := strings.Split(string(contents), "\n")
	badChars := []string{}
	lineCount := 0
	for _, line := range lines {
		l := strings.Split(line, "")
		badChar := checkLine(l)
		if badChar != "" {
			badChars = append(badChars, badChar)
		}
		lineCount += 1
		fmt.Println("line count: ", lineCount)
	}
	fmt.Println(makeScore(badChars))
}

func makeScore(badChars []string) int {
	parenthesesCount := 0
	squareCount := 0
	curlyCount := 0
	greaterThanCount := 0
	for _, badChar := range badChars {
		if badChar == ")" {
			parenthesesCount += 1
		} else if badChar == "]" {
			squareCount += 1
		} else if badChar == "}" {
			curlyCount += 1
		} else if badChar == ">" {
			greaterThanCount += 1
		}
	}
	return (3 * parenthesesCount) + (57 * squareCount) + (1197 * curlyCount) + (25137 * greaterThanCount)
}

func checkLine(line []string) string {
	openingBrackets := []string{}
	for _, char := range line {
		if char == "(" || char == "[" || char == "{" || char == "<" {
			openingBrackets = append(openingBrackets, char)
		} else {
			if char == ")" && openingBrackets[len(openingBrackets)-1] == "(" {
				openingBrackets = sliceDelete(openingBrackets)
				continue
			}
			if char == "]" && openingBrackets[len(openingBrackets)-1] == "[" {
				openingBrackets = sliceDelete(openingBrackets)
				continue
			}
			if char == "}" && openingBrackets[len(openingBrackets)-1] == "{" {
				openingBrackets = sliceDelete(openingBrackets)
				continue
			}
			if char == ">" && openingBrackets[len(openingBrackets)-1] == "<" {
				openingBrackets = sliceDelete(openingBrackets)
				continue
			}
			return char
		}
	}
	return ""
}

func sliceDelete(openingBrackets []string) []string {
	copy(openingBrackets[len(openingBrackets)-1:], openingBrackets[len(openingBrackets):])
	openingBrackets[len(openingBrackets)-1] = ""
	openingBrackets = openingBrackets[:len(openingBrackets)-1]
	return openingBrackets
}
