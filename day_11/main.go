package main

import (
	"aoc23/myLib"
	"fmt"
	"log"
	"math"
)

func main() {
	inputBytes := myLib.ErrHandledRead("input.txt")
	input := [][]byte{{}}
	row := 0
	for i := 0; i < len(inputBytes)-1; i++ {
		if inputBytes[i] == '\n' {
			input = append(input, []byte{})
			row++
		} else {
			input[row] = append(input[row], inputBytes[i])
		}
	}
	fmt.Println(input)
	log.Println(partOne(input))
	log.Println(partTwo(input))
}

func partTwo(input [][]byte) int {
	return 0
}

type galaxy struct {
	x, y float64
}

func partOne(input [][]byte) int {
	galaxyColCounter := make([]float64, len(input[0]))
	var galaxyPos []galaxy
	// galaxyScale := 0

	// add new rows vertically
	for i := 0; i < len(input); i++ {
		pos := myLib.IndexAll(input[i], '#')
		if len(pos) == 0 {
			input = append(input[:i], append([][]byte{input[i]}, input[i:]...)...)
			i++
		} else {
			for j := 0; j < len(pos); j++ {
				galaxyColCounter[pos[j]]++
				fj := float64(pos[j])
				fi := float64(i)
				if galaxyPos == nil {
					galaxyPos = []galaxy{{x: fj, y: fi}}
				} else {
					galaxyPos = append(galaxyPos, galaxy{x: fj, y: fi})
				}
				pos[j] = pos[j] + i
			}
		}
	}
	var rowI int = 0
	// add new rows horizontally
	for i := 0; i < len(galaxyColCounter); i++ {
		if galaxyColCounter[i] == 0 {
			for j := 0; j < len(input); j++ {
				input[j] = append(input[j][:rowI], append([]byte{input[j][rowI]}, input[j][rowI:]...)...)

			}
			floatRowI := float64(rowI)
			for k := 0; k < len(galaxyPos); k++ {

				if galaxyPos[k].x > floatRowI {
					galaxyPos[k].x += 1
				}
			}
			rowI++
		}
		rowI++
	}
	var total float64
	for i := 0; i < len(galaxyPos); i++ {
		for j := i + 1; j < len(galaxyPos); j++ {
			steps := math.Abs(galaxyPos[j].x-galaxyPos[i].x) + math.Abs(galaxyPos[j].y-galaxyPos[i].y)
			total += steps
			log.Printf("%v -> %v = %v", galaxyPos[i], galaxyPos[j], steps)
		}
	}

	// fmt.Println(input, galaxyPos)
	return int(total)
}
