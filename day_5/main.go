package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal("failed to read input")
	}
	inputLines := strings.Split(string(input), "\n")
	partOne(inputLines)
	partTwo(inputLines)
}

func partTwo(input []string) {
}

type conv struct {
	min       int
	max       int
	convStart int
}

func partOne(input []string) {
	// map for each, use map address to store all values,
	conversionSlice := make([][]conv, 7)

	seeds := strings.Fields(input[0])
	fmt.Println(seeds[1:])
	// can loop through using if line break
	section := 0
	for i := 2; i < len(input); i++ {
		if len(input[i]) == 0 {
			section++
			continue
		}
		inputSlice := strings.Fields(input[i])
		if len(inputSlice) != 3 {
			continue
		}
		min, _ := strconv.Atoi(inputSlice[1])
		convStart, _ := strconv.Atoi(inputSlice[0])
		length, _ := strconv.Atoi(inputSlice[2])
		max := min + length

		conversionSlice[section] = append(conversionSlice[section], conv{min, max, convStart})

	}

	var locations sort.IntSlice
	for _, seed := range seeds[1:] {
		current, _ := strconv.Atoi(seed)
		for i := 0; i < len(conversionSlice); i++ {
			currentConv := conversionSlice[i]
		conversionChart:
			for j := 0; j < len(currentConv); j++ {
				if current > currentConv[j].min && current < currentConv[j].max {
					current = current - currentConv[j].min + currentConv[j].convStart
					fmt.Println("converted", current, currentConv[j])
					break conversionChart
				}
			}
			// fmt.Println(current)

		}
		locations = append(locations, current)
	}
	sort.Sort(locations)
	fmt.Println(locations[0])
}
