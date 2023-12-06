package main

import (
	"aoc23/myLib"
	"fmt"
	"strings"
)

func main() {
	partTwo(myLib.ErrHandledReadConv("input.txt"))
}

type conv struct {
	min       int
	max       int
	convStart int
}

func partTwo(input []string) {
	// map for each, use map address to store all values,
	conversionSlice := make([][]conv, 7)

	seeds := strings.Fields(input[0])
	// fmt.Println(seeds[1:])
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
		min := myLib.ErrHandledAtoi(inputSlice[1])
		convStart := myLib.ErrHandledAtoi(inputSlice[0])
		length := myLib.ErrHandledAtoi(inputSlice[2])
		max := min + length

		conversionSlice[section] = append(conversionSlice[section], conv{min, max, convStart})

	}
	// fmt.Println(conversionSlice)
	const MaxUint = ^uint(0)
	var location int = int(MaxUint >> 1)
	for k := 1; k < len(seeds); k += 2 {
		seedStart := myLib.ErrHandledAtoi(seeds[k])
		seedLen := myLib.ErrHandledAtoi(seeds[k+1])
		for seed := seedStart; seed < seedStart+seedLen; seed++ {
			current := seed
			// fmt.Println("current seed:", current)
			for i := 0; i < len(conversionSlice); i++ {
				currentConv := conversionSlice[i]
				for j := 0; j < len(currentConv); j++ {
					if current >= currentConv[j].min && current < currentConv[j].max {
						current = current - currentConv[j].min + currentConv[j].convStart
						// fmt.Println("converted", current, currentConv[j])
						break
					}
				}
			}

			if current < location {
				location = current
			}
		}
	}
	fmt.Println(location)
}
