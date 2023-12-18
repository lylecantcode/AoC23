package main

import (
	"aoc23/myLib"
	"flag"
	"fmt"
	"log"
	"slices"
	"sort"
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
	log.Println("<-start time")
	log.Println(Solution(input))
}

func Solution(input []string) int {
	total := 0
	for i := 0; i < len(input)-1; i++ {
		groups := myLib.StringToIntArrayFunc(input[i], func(r rune) bool {
			return !unicode.IsNumber(r)
		})
		// fixedState := myLib.IndexAll([]byte(input[i]), '.')
		damagedState := myLib.IndexAll([]byte(input[i]), '#')
		unknownState := myLib.IndexAll([]byte(input[i]), '?')
	
		numMissing := func() int {
			tot := -len(damagedState)
			for i := 0; i < len(groups); i++ {
				tot += groups[i]
			}
			return tot
		}()
		if numMissing == len(unknownState) {
			total += 1
			continue
		}

		if slices.Compare(groupedCount(damagedState), groups) == 0 {
			total++
			continue
		}

		track := map[string]interface{}{}
		sol := recursiveSolve(damagedState, unknownState, groups, track)
		// fmt.Println(sol, groups, "\n\n")
		total += sol
	}
	return total
}

func groupedCount(damagedState []int) []int {
	count := []int{1}
	index := 0
	for i := 0; i < len(damagedState)-1; i++ {
		if damagedState[i+1]-damagedState[i] != 1 {
			count = append(count, 1)
			index++
		} else {
			count[index]++
		}

	}
	return count
}

func recursiveSolve(s1, s2, pattern []int, track map[string]interface{}) int {
	var comp bool
	var val int
	uniqueCheck := fmt.Sprint(s1)
	_, exists := track[uniqueCheck]
	if !exists {
		comp = slices.Equal(groupedCount(s1), pattern)
		// fmt.Println(s1, pattern, s2, comp)
		track[uniqueCheck] = nil
	}

	if comp {
		val = 1
	}
	if len(s2) == 0 {
		return val
	}
	// skipped ?
	sol1 := recursiveSolve(s1, s2[1:], pattern, track)
	s3 := appendToSorted(s1[:], s2[0])

	// check if groupedCount(s1) == pattern
	sol2 := recursiveSolve(s3, s2[1:], pattern, track)

	return sol1 + sol2 + val // ideally do a check of right side and only check the pattern if new value is not adjacent to a damaged one
}

func appendToSorted(in []int, add int) []int {
	out := make([]int, len(in)+1)
	var added bool
	// fmt.Printf("in:%v, add:%v\n", in, add)
	if sort.IntsAreSorted(in) {
		iterator := 0
		for i := 0; i < len(in); i++ {
			if add < in[i] && !added {
				out[iterator] = add
				iterator++
				added = true
			}
			out[iterator] = in[i]
			iterator++
		}
	}
	if !added {
		out[len(out)-1] = add
	}
	return out
}
