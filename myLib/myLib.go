package myLib

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"math"
	"os"
	"reflect"
	"strconv"
	"strings"
)

func ErrHandledAtoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return i
}

func ErrHandledReadConv(s string) []string {
	a, err := os.ReadFile(s)
	if err != nil {
		log.Fatal(err)
	}
	return strings.Split(string(a), "\n")
}

func ErrHandledRead(s string) [][]byte {
	output := [][]byte{{}}
	file, err := os.Open(s)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		output = append(output, scanner.Bytes())

	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("failed to read %v, error: %v", s, err)
	}
	return output
}

func CheckSurroundings(i, j int, slice [][]*int) *int {
	if i < 0 || j < 0 || i >= len(slice) || j >= len(slice[i]) || slice[i][j] == nil {
		return nil
	}
	return slice[i][j]
}

func IntPtr(val int) *int {
	return &val
}

func BoolPtr(val bool) *bool {
	return &val
}

type set map[interface{}]interface{}

func (s set) exists(input interface{}) bool {
	_, e := s[input]
	return e
}

func (s set) Add(input interface{}) error {
	val := reflect.ValueOf(input)
	switch reflect.TypeOf(input).Kind() {
	// any single
	case reflect.Struct, reflect.Bool, reflect.String, reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Float32, reflect.Float64, reflect.Complex64, reflect.Complex128, reflect.Uintptr:
		s[val] = nil
		return nil
	case reflect.Slice, reflect.Map:
		for i := 0; i < val.Len(); i++ {
			s[val.Index(i)] = nil
		}
		return nil
	}
	return errors.New(fmt.Sprintf("Cannot add this type %v to the set", reflect.TypeOf(input)))
}

func (s set) AddMultiple(input ...interface{}) {
	for _, v := range input {
		s.Add(v)
	}
}

// try implementing with generics?
func (s set) Remove(input interface{}) {
	val := reflect.ValueOf(input)
	switch reflect.TypeOf(input).Kind() {
	// any single
	case reflect.Bool, reflect.String, reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Float32, reflect.Float64, reflect.Complex64, reflect.Complex128, reflect.Uintptr:
		delete(s, val)
	case reflect.Slice, reflect.Map:
		for i := 0; i < val.Len(); i++ {
			delete(s, val)
		}
	case reflect.Struct:
		fmt.Println("struct", val, reflect.TypeOf(input).Kind())
		s[input] = reflect.Value{}
	default:
		fmt.Println(val, reflect.TypeOf(input).Kind())
	}
}

func (s set) RemoveMultiple(input ...interface{}) {
	for _, v := range input {
		s.Remove(v)
	}
}

func (s set) Clear() {
	for k := range s {
		delete(s, k)
	}
}

func (s set) Update(ns set) {
	for k := range ns {
		s[k] = nil
	}
}

func (s set) Union(ns set) set {
	newSet := set{}
	for k := range ns {
		newSet[k] = nil
	}
	for k := range s {
		newSet[k] = nil
	}
	return newSet
}

func (s set) Intersection(ns set) set {
	newSet := set{}
	for k := range s {
		_, exists := ns[k]
		if exists {
			newSet[k] = nil
		}
	}
	return newSet
}

func (s set) Difference(ns set) set {
	newSet := set{}
	for k := range s {
		_, exists := ns[k]
		if !exists {
			newSet[k] = nil
		}
	}
	return newSet
}

func (s set) Print() {
	var str strings.Builder
	for k := range s {
		fmt.Fprintf(&str, "%v ", k)
	}
	fmt.Println(str.String())
}

func (s set) Copy() set {
	ns := set{}
	ns.Update(s)
	return ns
}

func (s set) IsIntersected(ns set) bool {
	if len(s.Intersection(ns)) > 0 {
		return true
	}
	return false
}

func StringToIntArray(input string) []int {
	row := strings.Fields(input)
	output := []int{}
	for i := 0; i < len(row); i++ {
		output = append(output, ErrHandledAtoi(row[i]))
	}
	return output
}

type fn func(rune) bool

func StringToIntArrayFunc(input string, f fn) []int {
	row := strings.FieldsFunc(input, f)
	output := []int{}
	for i := 0; i < len(row); i++ {
		output = append(output, ErrHandledAtoi(row[i]))
	}
	return output
}

// any reason to implement pop?

func Biggest(integers ...int) int {
	max := math.MinInt
	for _, i := range integers {
		if i > max {
			max = i
		}
	}
	return max
}

func IndexAll[S ~[]E, E comparable](s S, v E) []int {
	indices := []int{}
	for i := range s {
		if v == s[i] {
			indices = append(indices, i)
		}
	}

	if len(indices) == 0 {
		return nil
	}
	return indices
}

func Transpose[S ~[][]E, E comparable](input S) S {
	outSlice := make(S, len(input[0]))
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[i]); j++ {
			if j == len(outSlice) {
				outSlice = append(outSlice, []E{})
			}
			outSlice[j] = append(outSlice[j], input[i][j])
		}
	}
	return outSlice
}
