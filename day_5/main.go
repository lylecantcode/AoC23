package main

import (
	// "fmt"
	"log"
	"os"
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
	var seeds []int
	var seedToSoil, soilToFert, fertToWater, waterToLight, lightToTemp, tempToHumid, humidToLoc []conv
	// change which conversion for the headings
	// if between min and max (loop), then convStart, otherwise 1:1

}
