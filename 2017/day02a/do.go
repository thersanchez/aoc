package main

import (
	"fmt"
	"strconv"
	"strings"
)

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

// returns an error if s has:
// - non-integers
// - empty lines
// - lines without data
// - no lines at all
func parse(s string) ([][]int, error) {
	lines := strings.Split(s, "\n")
	if len(lines) == 0 {
		return nil, fmt.Errorf("no data")
	}
	ret := make([][]int, len(lines))
	for ln, l := range lines {
		ww := strings.Fields(l)
		if len(ww) == 0 {
			return nil, fmt.Errorf("line %d: no data", ln+1)
		}
		ret[ln] = make([]int, len(ww))
		var err error
		for wn, w := range ww {
			ret[ln][wn], err = strconv.Atoi(w)
			if err != nil {
				return nil, fmt.Errorf("line %d, word %d: %v", ln+1, wn+1, err)
			}
		}
	}
	return ret, nil
}

// Assumes len(nn)>0
func min(nn []int) int {
	r := nn[0]
	for _, n := range nn[1:] {
		if n < r {
			r = n
		}
	}
	return r
}

// Assumes len(nn)>0
func max(nn []int) int {
	r := nn[0]
	for _, n := range nn[1:] {
		if n > r {
			r = n
		}
	}
	return r
}
