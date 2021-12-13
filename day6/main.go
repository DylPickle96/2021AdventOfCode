package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type lanternfish struct {
	timer int64
}

func main() {
	contents, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Printf("cannot read input.txt. Error: %v", err)
		return
	}
	values := strings.Split(string(contents), ",")
	school := []lanternfish{}
	for _, value := range values {
		v, _ := strconv.ParseInt(value, 10, 64)
		school = append(school, lanternfish{
			timer: v,
		})
	}
	// I'm sorry CPU
	for i := 1; i <= 256; i++ {
		newFish := []lanternfish{}
		for i := 0; i < len(school); i++ {
			// if the lanternFish's timer hits zero it spawns a new fish and resets its timer to 6
			if school[i].timer == 0 {
				newFish = append(newFish, lanternfish{
					timer: 8,
				})
				school[i].timer = 6
				continue
			}
			school[i].timer -= 1
		}
		school = append(school, newFish...)
	}
	fmt.Println(len(school))
}
