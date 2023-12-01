package main

import (
	// "fmt"
	"log"
	"os"
	// "strings"
)

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal("failed to read input")
	}

}
