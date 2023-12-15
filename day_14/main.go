package main

import (
	"aoc23/myLib"
	"bytes"
	"flag"
	"fmt"
	"log"
	"strings"
)

func main() {
	testFlag := flag.Bool("test", false, "controls which input to use")
	flag.Parse()
	inputFile := "input.txt"
	if *testFlag {
		inputFile = "test_input.txt"
	}
	input := myLib.ErrHandledReadConv(inputFile)
	log.Println(partOne(input))
	log.Println(partTwo(input))
}

func partTwo(input []string) int {
	return 0
}

func partOne(input []string) int {
	for i := 0; i < len(input[0]); i++ {
		for j := len(input) - 2; j > 0; j-- {
			rOne := []byte(input[j])
			rTwo := []byte(input[j-1])
			// log.Printf("%v: %v\n", i, j)
			if rOne[i] == 'O' && rTwo[i] == '.' {
				rTwo[i] = 'O'
				rOne[i] = '.'
				input[j] = string(rOne)
				input[j-1] = string(rTwo)
				j = len(input) - 1
			}
		}
	}

	total := 0
	for i := 0; i < len(input); i++ {
		total += strings.Count(input[i], "O") * (len(input) - 1 - i)
	}
	return total
}

func partOneByte(input [][]byte) int {
	for i := 0; i < len(input[0]); i++ {
		for j := len(input) - 1; j > 0; j-- {
			if input[j][i] == 'O' && input[j-1][i] == '.' {
				input[j-1][i] = 'O'
				input[j][i] = '.'
				j = len(input)
			}
		}
	}
	myLib.PrintBytes(input)
	total := 0
	for i := 0; i < len(input); i++ {
		// fmt.Println(bytes.Count(input[i], []byte{'O'}), (len(input) - i))
		total += bytes.Count(input[i], []byte{'O'}) * (len(input) - i)
	}
	return total
}

func partOneTransposed(input [][]byte) int {
	// transpose and then just move all O next to #
	// rotates it so that North is now West, but same (0,0)
	// totalRocks := 0
	transInput := myLib.Transpose(input)
	// placementArray := make([]int, len(transInput[0]))
	myLib.PrintBytes(transInput)
	for _, row := range transInput {
		for j := len(row) - 1; j > 0; j-- {
			if row[j] == 'O' && row[j-1] == '.' {
				row[j-1] = 'O'
				row[j] = '.'
				j = len(row)
			}
		}
	}
	// total := 0
	// scale := len(input)
	// fmt.Println(scale)
	// for i := 0; i < scale; i++ {
	// 	total += placementArray[i] * (scale - i)
	// }
	revertedInput := myLib.Transpose(transInput)
	total := 0
	for i := 0; i < len(revertedInput); i++ {
		fmt.Println(bytes.Count(revertedInput[i], []byte{'O'}), (len(revertedInput) - i))
		total += bytes.Count(revertedInput[i], []byte{'O'}) * (len(revertedInput) - i)
	}
	return total
}
