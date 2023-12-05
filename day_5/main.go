package main

import (
	"cmp"
	"fmt"
	"log"
	"os"
	"slices"
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
	var seedToSoil, soilToFert, fertToWater, waterToLight, lightToTemp, tempToHumid, humidToLoc []conv
	conversionSlice := make([]conv, 7)

	seeds := strings.Fields(input[0])
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
		min, _ := strconv.Atoi(inputSlice[0])
		convStart, _ := strconv.Atoi(inputSlice[1])
		length, _ := strconv.Atoi(inputSlice[2])
		max := min + length

		switch section {
		case 0:
			seedToSoil = append(seedToSoil, conv{min, max, convStart})
		case 1:
			soilToFert = append(soilToFert, conv{min, max, convStart})
		case 2:
			fertToWater = append(fertToWater, conv{min, max, convStart})
		case 3:
			waterToLight = append(waterToLight, conv{min, max, convStart})
		case 4:
			lightToTemp = append(lightToTemp, conv{min, max, convStart})
		case 5:
			tempToHumid = append(tempToHumid, conv{min, max, convStart})
		case 6:
			humidToLoc = append(humidToLoc, conv{min, max, convStart})
		}
	}
	slices.SortFunc(seedToSoil, sortConv)
	slices.SortFunc(soilToFert, sortConv)
	slices.SortFunc(fertToWater, sortConv)
	slices.SortFunc(waterToLight, sortConv)
	slices.SortFunc(lightToTemp, sortConv)
	slices.SortFunc(tempToHumid, sortConv)
	slices.SortFunc(humidToLoc, sortConv)
	for _, seed := range seeds {
		for i := 0; i < len(seedToSoil); i++ {

		}
	}

	fmt.Println(seeds, humidToLoc)
}

func sortConv(a, b conv) int {
	return cmp.Compare(a.min, b.min)
}
