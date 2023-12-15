package main

import (
	"aoc23/myLib"
	"flag"
	"fmt"
	"log"
)

func main() {
	testFlag := flag.Bool("test", false, "controls which input to use")
	flag.Parse()
	inputFile := "input.txt"
	if *testFlag {
		inputFile = "test_input.txt"
	}
	input := myLib.ErrHandledRead(inputFile)
	log.Println(partOne(input))
	log.Println(partTwo(input))
}

func partTwo(input [][]byte) int {
	return 0
}

func partOne(input [][]byte) int {
	// transpose and then just move all O next to #
	// rotates it so that North is now West, but same (0,0)
	totalRocks := 0
	placementArray := make([]int, len(input))
	transInput := myLib.Transpose(input)
	for i := 0; i < len(transInput); i++ {
		stops := myLib.IndexAll(transInput[i], '#')
		rocks := myLib.IndexAll(transInput[i], 'O')
		totalRocks += len(rocks)
		// all other rocks to stack up
		stops = append(stops, 999)
		// log.Println(rocks, stops)
		stoppedAt := 0
		for _, stop := range stops {
			for _, rock := range rocks {
				if rock >= stoppedAt && rock < stop {
					// fmt.Println(placementArray)
					placementArray[stoppedAt] += 1
					fmt.Println(placementArray)
					stoppedAt++
				}
			}
			stoppedAt = stop + 1
		}
	}
	fmt.Println(placementArray)
	total := 0
	scale := len(input)
	rockCheck := 0
	for i := 0; i < scale; i++ {
		total += placementArray[i] * (scale - 1 - i)
		rockCheck += placementArray[i]
	}
	log.Println(totalRocks, rockCheck)
	return total
}
