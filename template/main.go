package main

import (
	"aoc23/myLib"
	"flag"
	"log"
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
	return 0
}
