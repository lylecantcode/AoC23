package main

import (
	"aoc23/myLib"
	"flag"
	"fmt"
	"log"
	"strings"
)

func main() {
	fmt.Println("rn=1", []byte("rn=1"))
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
	/*Determine the ASCII code for the current character of the string.
	Increase the current value by the ASCII code you just determined.
	Set the current value to itself multiplied by 17.
	Set the current value to the remainder of dividing itself by 256.*/
	return total
}
