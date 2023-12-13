package main

import (
	"aoc23/myLib"
	"flag"
	"log"
	"slices"
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

func partTwo(input []string) int {
	return 0
}

type springGroup struct {
	length, damagedCount                                                int
	brokenGroups, fixedIndex, notfixedIndex, damagedIndex, unknownIndex []int
	variations                                                          [][]int
}

func partOne(input []string) int {
	var springs []springGroup
	for i := 0; i < len(input)-1; i++ {
		damagedSpringGroups := myLib.StringToIntArrayFunc(input[i], func(r rune) bool {
			return !unicode.IsNumber(r)
		})
		damagedCount := func() int {
			tot := 0
			for i := 0; i < len(damagedSpringGroups); i++ {
				tot += damagedSpringGroups[i]
			}
			return tot
		}()
		fixedState := myLib.IndexAll([]byte(input[i]), '.')
		damagedState := myLib.IndexAll([]byte(input[i]), '#')
		unknownState := myLib.IndexAll([]byte(input[i]), '?')
		notFixedState := append(damagedState, unknownState...)
		slices.Sort(notFixedState)
		springs = append(springs, springGroup{length: len(input[i]) - 2*len(damagedSpringGroups), damagedCount: damagedCount, fixedIndex: fixedState,
			notfixedIndex: notFixedState, brokenGroups: damagedSpringGroups, damagedIndex: damagedState, unknownIndex: unknownState})
	}
	arrangements := 0
	for _, spring := range springs {
		arrangements += springElim(spring)

	}
	return arrangements
}

/* func springFirstTry(spring springGroup) int {
	var arrangements int
	if len(spring.unknownIndex) == 0 || spring.length == spring.damagedCount {
		arrangements += 1
	} else {
		// has to start at most (length - damagedcount) away but also has to account for the split between damaged count
	nextSpring:
		for i := 0; i < spring.length-(spring.damagedCount-1+len(spring.brokenGroups)-1); i++ {
			// is spring potentially broken?
			if slices.Contains(spring.notfixedIndex, i) {
				pos := i
				var variation []int
				// broken groups
				for _, v := range spring.brokenGroups {
					// check if it matches a spring group
					for j := 0; j < v; j++ {
						if !slices.Contains(spring.notfixedIndex, pos) {
							continue
						}
						variation = append(variation, pos)
						pos++
					}
					// the space between groups:
					pos++
				}
				if len(variation) == spring.damagedCount {
					log.Println("values", variation, "expected", spring.damagedCount)
					spring.variations = append(spring.variations, variation)
				} else {
					log.Println("too few values", variation, "expected", spring.damagedCount)
				}
				for k := 0; k < len(spring.damagedIndex); k++ {
					if !slices.Contains(variation, spring.damagedIndex[k]) {
						continue nextSpring
					}
				}
				// add a check to make sure all the fixed springs are counted before adding them
				// this means storing another array
				log.Println(pos, spring.notfixedIndex, spring.brokenGroups, spring.variations)
				arrangements += len(spring.variations)
			}

		}
	}
 }*/

func springCheck(sg springGroup) int {
	total := 0
	for i := 0; i < sg.length-(sg.damagedCount+len(sg.brokenGroups)-1); i++ {
		sg.variations = append(sg.variations, []int{})
		indexCheck := slices.Index(sg.notfixedIndex, i)
		if indexCheck >= 0 {
			for _, v := range sg.brokenGroups {
				max := i + v
				if max < len(sg.notfixedIndex) && sg.notfixedIndex[indexCheck+v] != max {
					continue
				} else {
					total++
					sg.variations[i] = append(sg.variations[i], i)
					log.Println("woo", total, i, sg.variations)
				}
			}

		}
	}
	return total
}

func springElim(sg springGroup) int {
	group := group(sg.notfixedIndex)
	log.Println(group)
	// output := []int{}
	// 1, 1, 3
	// 7
	// 1, 1, 3
	// 3, 5
	// 1, 1, 3
	// 4, 4, 4
	current := 0
	for i := 0; i < len(sg.brokenGroups); i++ {
		numBroken := sg.brokenGroups[i]
		if current > len(group)-1 {
			// out of groups
			break
		}
		if numBroken > group[current] {
			// try again with a different group
			current++
			i--
		} else {
			// group[current] = group[current]-numBroken
			group[current] -= numBroken + 1
		}
	}
	total := 1
	for i := 0; i < len(group); i++ {
		if group[i] >= 0 {
			total *= group[i] + 2
		}
	}
	log.Println("elim", group, sg.brokenGroups, total)
	return 0
}

func group(slice []int) []int {
	returnSlice := []int{1}
	log.Println(slice)
	for i := 0; i < len(slice); i++ {
		// log.Println(slice[i], returnSlice)
		if i >= len(slice)-2 {
			if slice[i]-slice[i-1] == 1 {
				returnSlice[len(returnSlice)-1]++
			} else {
				returnSlice = append(returnSlice, 1)
			}
			break
		} else if slice[i+1]-slice[i] == 1 {
			returnSlice[len(returnSlice)-1]++
		} else {
			returnSlice = append(returnSlice, 1)
		}
	}
	// log.Println(slice[i], returnSlice)
	return returnSlice
}
