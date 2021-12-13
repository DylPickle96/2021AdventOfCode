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
	horizontal := int64(0)
	depth := int64(0)
	aim := int64(0)
	for _, line := range lines {
		instruction := strings.Split(line, " ")
		value, err := strconv.ParseInt(instruction[1], 10, 64)
		if err != nil {
			log.Printf("cannot parse line. Error: %v", err)
		}
		switch instruction[0] {
		case "up":
			aim = aim - value
		case "down":
			aim = aim + value
		case "forward":
			horizontal = horizontal + value
			depth = depth + (value * aim)
		}
	}
	fmt.Println("horizontal: ", horizontal, "depth: ", depth)
	fmt.Println(horizontal * depth)
}
