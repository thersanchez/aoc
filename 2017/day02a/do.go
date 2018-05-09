package main

import "fmt"

func do(input string) (int, error) {
	ss, err := parse(input)
	if err != nil {
		return 0, fmt.Errorf("parsing input: %v", err)
	}
	var sum int
	for _, s := range ss {
		sum += max(s) - min(s)
	}
	return sum, nil
}

// returns a error if s has:
// - non-integers
// - negative integers
// - empty lines
func parse(s string) ([][]int, error) {
	return [][]int{{}}, nil
}

// Assumes len(s)>0
func min(s []int) int {
	return 0
}

// Assumes len(s)>0
func max(s []int) int {
	return 0
}
