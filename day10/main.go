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
		if strings.Contains(line, " ") {
			fmt.Println("Contains space")
			continue
		}
		l := strings.Split(line, "")
		badChar := checkLine(l)
		if badChar != "" {
			badChars = append(badChars, badChar)
		}
		lineCount += 1
		fmt.Println("line count: ", lineCount)
	}
	fmt.Println(badChars)
}

func checkLine(line []string) string {
	openingBrackets := []string{}
	for _, char := range line {
		if char == "(" || char == "[" || char == "{" || char == "<" {
			openingBrackets = append(openingBrackets, char)
		} else {
			if char == ")" && openingBrackets[len(openingBrackets)-1] == "(" {
				openingBrackets = sliceDelete(openingBrackets)
			} else {
				fmt.Println("line:", line, "char:", char)
				return char
			}
			if char == "]" && openingBrackets[len(openingBrackets)-1] == "[" {
				openingBrackets = sliceDelete(openingBrackets)
			} else {
				fmt.Println("line:", line, "char:", char)
				return char
			}
			if char == "}" && openingBrackets[len(openingBrackets)-1] == "{" {
				openingBrackets = sliceDelete(openingBrackets)
			} else {
				fmt.Println("line:", line, "char:", char)
				return char
			}
			if char == ">" && openingBrackets[len(openingBrackets)-1] == "<" {
				openingBrackets = sliceDelete(openingBrackets)
			} else {
				fmt.Println("line:", line, "char:", char)
				return char
			}
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
