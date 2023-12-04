package main

import (
	"fmt"
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

func partOne(input []string) {
	// last row has a new line on it
	size := len(input) - 1
	schematic := make([][]*int, size)
	for i := 0; i < size; i++ {
		width := len(input[i])
		schematic[i] = make([]*int, width)
		for j := 0; j < width; j++ {
			val := input[i][j]
			if val >= 48 && val <= 57 {
				schematic[i][j] = intPtr(int(val) - 48)
			} else if val != '.' {
				schematic[i][j] = intPtr(0)
			}
			_ = partCheck(i-1, j, schematic)
		}
		if i == size-1 {
			for k := 0; k < len(schematic[i]); k++ {
				_ = partCheck(i-1, k, schematic)
			}
		}
	}
	fmt.Println(*schematic[0][0], *schematic[0][1], *schematic[0][2])
}

func partTwo(input []string) {

}

func intPtr(i int) *int {
	return &i
}

func partCheck(i, j int, schematic [][]*int) int {
	if i < 0 {
		return 0
	}
	total := 0
	if schematic[i][j] != nil && *schematic[i][j] == 0 {
		if i > 0 {
			fmt.Println("case 1")
			if schematic[i-1][j] != nil {
				total += *coordinates(j, 0, schematic[i-1], total)

			}
		}
		if j != 0 {
			fmt.Println("case 2")
			if schematic[i][j-1] != nil {
				x := j - 1
				lastVal := 0
				for {
					if schematic[i][x] == nil {
						break
					}
					lastVal = *schematic[i][x]
					x++
				}
				total += lastVal
			}
		}
		if i != len(schematic)-2 {
			fmt.Println("case 3")
			if schematic[i+1][j] != nil {
			}
		}
		if j != len(schematic[i])-2 {
			fmt.Println("case 4")
			if schematic[i][j+1] != nil {
			}
		}
	}
	return total
}

func coordinates(j, dir int, schematicLine []*int, total int) *int {
	// if it would go out of bounds
	if schematicLine[j] == nil || j >= len(schematicLine) || j < 0 || *schematicLine[j] == 0 {
		return nil
	}
	// direction for recursion, right (1), left (-1), both (0)
	if dir == 0 {
		total = *schematicLine[j]
	}

	// if going left
	if dir < 0 || dir == 0 {
		fmt.Println("going left", total)
		if dir != 0 {
			total = *schematicLine[j]*10 + total
		}
		val := coordinates(j-1, -1, schematicLine, total)
		if val != nil && total/10 == *val {
			total = *val
		}
		fmt.Println("gone left", total)
	}
	// if going right
	if dir > 0 || dir == 0 {
		fmt.Println("going right", total)
		if dir != 0 {
			total = 10*(total) + *schematicLine[j]
		}
		val := coordinates(j+1, 1, schematicLine, total)
		if val != nil && *val/10 == total {
			total = *val
		}
	}
	fmt.Println("the total:\t\t", total)
	schematicLine[j] = nil
	return &total
}
