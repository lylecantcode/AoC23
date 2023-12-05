package main

import (
	"fmt"
	"log"
	"os"
	"slices"
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

func partOne(input []string) {
	size := len(input)
	schematic := make([][]*int, size)
	symbols := make([][]bool, size)
	for i := 0; i < size; i++ {
		width := len(input[i])
		schematic[i] = make([]*int, width)
		symbols[i] = make([]bool, width)
		for j := 0; j < width; j++ {
			val := input[i][j]
			if val >= 48 && val <= 57 {
				schematic[i][j] = intPtr(int(val) - 48)
			} else if val != '.' {
				symbols[i][j] = true
			}
		}
	}

	for i := 0; i < len(schematic); i++ {
		val := 0
		length := 0
		for j := 0; j < len(schematic[i]); j++ {
			current := schematic[i][j]
			if current != nil {
				val = val*10 + *current
				length++
			} else {
				ptrVal := intPtr(val)
				for k := length; k > 0; k-- {
					schematic[i][j-k] = ptrVal
				}
				val = 0
				length = 0
			}
			if j == len(schematic[i])-1 && val != 0 {
				ptrVal := intPtr(val)
				for k := length; k > 0; k-- {
					schematic[i][j-k] = ptrVal
				}
			}
		}
	}
	var addressSlice []*int
	for i := 0; i < len(symbols); i++ {
		for j := 0; j < len(symbols[i]); j++ {
			if symbols[i][j] {
				// check the 9 squares around
				// var symbolTracker []*int
				for k := i - 1; k <= i+1; k++ {
					for l := j - 1; l <= j+1; l++ {
						address := checkProximities(k, l, schematic)
						if address != nil && !slices.Contains(addressSlice, address) {
							addressSlice = append(addressSlice, address)
							// symbolTracker = append(symbolTracker, address)
						}
					}
				}
			}
		}
	}
	total := 0
	for i := 0; i < len(addressSlice); i++ {
		fmt.Println(*addressSlice[i])
		total += *addressSlice[i]
	}
	fmt.Println(total)
}

func checkProximities(i, j int, slice [][]*int) *int {
	if i < 0 || j < 0 || i >= len(slice) || j >= len(slice[i]) || slice[i][j] == nil {
		return nil
	}
	return slice[i][j]
}

func partTwo(input []string) {
	size := len(input)
	schematic := make([][]*int, size)
	symbols := make([][]bool, size)
	for i := 0; i < size; i++ {
		width := len(input[i])
		schematic[i] = make([]*int, width)
		symbols[i] = make([]bool, width)
		for j := 0; j < width; j++ {
			val := input[i][j]
			if val >= 48 && val <= 57 {
				schematic[i][j] = intPtr(int(val) - 48)
			} else if val == '*' {
				symbols[i][j] = true
			}
		}
	}

	for i := 0; i < len(schematic); i++ {
		val := 0
		length := 0
		for j := 0; j < len(schematic[i]); j++ {
			current := schematic[i][j]
			if current != nil {
				val = val*10 + *current
				length++
			} else {
				ptrVal := intPtr(val)
				for k := length; k > 0; k-- {
					schematic[i][j-k] = ptrVal
				}
				val = 0
				length = 0
			}
			if j == len(schematic[i])-1 && val != 0 {
				ptrVal := intPtr(val)
				for k := length; k > 0; k-- {
					schematic[i][j-k] = ptrVal
				}
			}
		}
	}
	var total int
	var addressSlice []*int
	for i := 0; i < len(symbols); i++ {
		for j := 0; j < len(symbols[i]); j++ {
			if symbols[i][j] {
				// check the 9 squares around
				var symbolTracker []*int
				for k := i - 1; k <= i+1; k++ {
					for l := j - 1; l <= j+1; l++ {
						address := checkProximities(k, l, schematic)
						if address != nil && !slices.Contains(addressSlice, address) {
							addressSlice = append(addressSlice, address)
							symbolTracker = append(symbolTracker, address)
						}
					}
				}
				if len(symbolTracker) == 2 {
					total += *symbolTracker[0] * *symbolTracker[1]
				}
			}
		}
	}
	fmt.Println(total)
}

func intPtr(i int) *int {
	return &i
}
