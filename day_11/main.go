package main

import (
	"aoc23/myLib"
	"flag"
	"log"
	"math"
)

func main() {
	test := flag.Bool("test", false, "controls which input to use")
	flag.Parse()
	inputFile := "input.txt"
	if *test {
		inputFile = "test_input.txt"
	}
	inputBytes := myLib.ErrHandledRead(inputFile)
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
	const scale = 1000000
	// add new rows vertically
	var floatCol int
	for i := 0; i < len(input); i++ {
		pos := myLib.IndexAll(input[i], '#')
		if len(pos) == 0 {
			floatCol += scale - 1
		}
		for j := 0; j < len(pos); j++ {
			// log.Printf("{%v,%v}", pos[j], i)
			fj := float64(pos[j])
			fi := float64(i + floatCol)
			if galaxyPos == nil {
				galaxyPos = []galaxy{{x: fj, y: fi}}
			} else {
				galaxyPos = append(galaxyPos, galaxy{x: fj, y: fi})
			}
			galaxyColCounter[pos[j]]++
		}
	}
	var rowI int = 0
	// add new rows horizontally
	log.Println(galaxyColCounter)
	for i := 0; i < len(galaxyColCounter); i++ {
		if galaxyColCounter[i] == 0 {
			// log.Println(i, rowI)
			floatRowI := float64(rowI + i)
			for k := 0; k < len(galaxyPos); k++ {
				if galaxyPos[k].x > floatRowI {
					galaxyPos[k].x += scale - 1
				}
			}
			rowI += scale - 1
		}
	}
	var total float64
	for i := 0; i < len(galaxyPos); i++ {
		for j := i + 1; j < len(galaxyPos); j++ {
			steps := math.Abs(galaxyPos[j].x-galaxyPos[i].x) + math.Abs(galaxyPos[j].y-galaxyPos[i].y)
			total += steps
		}
	}

	log.Println(galaxyPos)
	return int(total)
}
