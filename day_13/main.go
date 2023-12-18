package main

import (
	"aoc23/myLib"
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

var verbose bool

func main() {
	test := flag.Bool("test", false, "controls which input to use")
	flag.BoolVar(&verbose, "v", false, "print statements")
	flag.Parse()
	inputFile := "input.txt"
	if *test {
		inputFile = "test_input.txt"
	}
	// normal method was struggling with the linebreaks#

	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	patterns := [][]string{{}}

	for scanner.Scan() {
		if len(scanner.Text()) == 0 {
			patterns = append(patterns, []string{})
		} else {
			patterns[len(patterns)-1] = append(patterns[len(patterns)-1], scanner.Text())
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	log.Printf("answer: [%v] for [%v]\n", partOne(patterns), inputFile)
}

func partOne(patterns [][]string) int {
	total := 0
	for _, pattern := range patterns {
		sym := vertSymCheck(pattern)
		if sym > 0 {
			if verbose {
				fmt.Println("vert")
			}
			total += sym
		} else {
			if verbose {
				fmt.Println("hor")
			}
			hor := HorSymCheck(pattern)
			if hor <= 0 {
				fmt.Println("no matches at all", pattern)
			}
			total += hor
		}
	}
	return total
}

func vertSymCheck(pattern []string) int {
	transPattern := [][]byte{}
	for i := 0; i < len(pattern); i++ {
		transPattern = append(transPattern, []byte(pattern[i]))
	}
	transPattern = myLib.Transpose(transPattern)
	stringTransPattern := []string{}
	for i := 0; i < len(transPattern); i++ {
		stringTransPattern = append(stringTransPattern, string(transPattern[i]))
	}
	return HorSymCheck(stringTransPattern) / 100
}

type mirror struct {
	pattern       []string
	reflectSmudge map[int]bool
}

func HorSymCheck(pattern []string) int {
	// change to include smudged and not smudged as struct  per
	m := mirror{pattern, map[int]bool{}}
	for i := 0; i < len(pattern)-1; i++ {
		compare := sliceCompare([]byte(pattern[i]), []byte(pattern[i+1]))
		// check if smudged
		if compare == 1 {
			m.reflectSmudge[i] = false
		} else if compare == 0 {
			m.reflectSmudge[i] = true
		}
	}
	if verbose {
		log.Println("sym lines", m)
	}
	// checking for smudged lines

	for line, smudge := range m.reflectSmudge {
		if !smudge {
			fmt.Printf("already smudged - possible sym line %v\n", line)
			sym := true
			// check for reflection around the line, if hits boundaries,
			for i := 1; i < len(pattern) && sym; i++ {
				// reflected pairs
				lineOne, lineTwo := line-i, line+1+i
				if lineOne < 0 || lineTwo >= len(pattern) {
					break
				}
				comparison := sliceCompare([]byte(pattern[lineOne]), []byte(pattern[lineTwo]))
				// already used the smudge
				// if comparison == 1 {
				// 	smudge = true
				// }
				if comparison >= 1 {
					sym = false
				}
			}
			if sym && !smudge {
				return (line + 1) * 100
			}
		} else {
			// check around reflection line, and if smudged at end, continue
			// otherwise return (line + 1)*100
			sym := true
			fmt.Printf("for possible sym line %v\n", line)
			for i := 1; i < len(pattern) && sym; i++ {
				// reflected pairs
				lineOne, lineTwo := line-i, line+1+i
				if lineOne < 0 || lineTwo >= len(pattern) {
					break
				}
				comparison := sliceCompare([]byte(pattern[lineOne]), []byte(pattern[lineTwo]))

				if comparison == 1 {
					smudge = false
					fmt.Printf("found a difference between lines %v and %v (i := %v)\n", lineOne, lineTwo, i)
				}
				if comparison > 1 {
					fmt.Printf("too many difs between lines %v and %v\n", lineOne, lineTwo)
					sym = false
				}
			}
			if sym && !smudge {
				return (line + 1) * 100
			}
		}

	}

	if verbose {
		fmt.Println("no matches")
	}
	return -1
}

func sliceCompare[S ~[]E, E comparable](s1, s2 S) int {
	var dif int
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			dif++
		}
	}

	if verbose {
		// fmt.Printf("%v diff: \n\t%v\n\t%v\n", dif, s1, s2)
	}
	return dif
}
