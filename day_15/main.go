package main

import (
	"aoc23/myLib"
	"flag"
	"fmt"
	"log"
	"strings"
	"unicode"
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

type lens struct {
	key   string
	value int
}

func partTwo(input []string) int {
	in := strings.Split(input[0], ",")
	outSlice := make([][]lens, 256)
startLoop:
	for _, val := range in {
		out := strings.FieldsFunc(val, func(r rune) bool {
			return !unicode.IsLetter(r) && !unicode.IsNumber(r) // number added for the test data
		})
		key := out[0]
		hash := partOne([]string{key})
		if len(out) == 1 {
			for i := 0; i < len(outSlice[hash]); i++ {
				if outSlice[hash][i].key == key {
					outSlice[hash] = append(outSlice[hash][:i], outSlice[hash][i+1:]...)
					break
				}
			}
		}
		if len(out) == 2 {
			for i := 0; i < len(outSlice[hash]); i++ {
				if outSlice[hash][i].key == key {
					outSlice[hash][i].value = myLib.ErrHandledAtoi(out[1])
					continue startLoop
				}
			}
			outSlice[hash] = append(outSlice[hash], lens{key: key, value: myLib.ErrHandledAtoi(out[1])})
		}
	}

	total := 0
	for i := 0; i < len(outSlice); i++ {
		for j := 0; j < len(outSlice[i]); j++ {
			total += (i + 1) * (j + 1) * (outSlice[i][j].value)
			fmt.Println((i + 1), (j + 1), (outSlice[i][j].value), (outSlice[i][j].key))
		}
	}

	return total
}

func partOne(input []string) int {
	total := 0
	in := strings.Split(input[0], ",")
	for _, val := range in {
		bVal := []byte(val)
		valTot := 0
		for i := 0; i < len(bVal); i++ {
			valTot += int(bVal[i])
			valTot *= 17
			valTot = valTot % 256
		}
		total += valTot
	}
	return total
}
