package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func do(input string) (int, error) {
	ss, err := parse(input)
	if err != nil {
		return 0, fmt.Errorf("parsing input: %v", err)
	}

	var (
		sum      int
		dividend int
		divisor  int
	)
	for i, s := range ss {
		dividend, divisor, err = findEvenlyDiv(s)
		if err != nil {
			return 0, fmt.Errorf("in line %d: %v", i+1, err)
		}
		sum += dividend / divisor
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

func findEvenlyDiv(nn []int) (dividend int, divisor int, err error) {
	for _, dividend = range nn {
		for _, divisor = range nn {
			if divisor == dividend {
				continue
			}
			if dividend%divisor == 0 {
				return dividend, divisor, nil
			}
		}
	}
	return 0, 0, errors.New("no evenly divisible values")
}
