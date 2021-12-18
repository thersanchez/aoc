package day07a

import (
	"fmt"
	"strconv"
	"strings"
)

func ParseLine(line string) (
	name string,
	weight int,
	children []string,
	err error,
) {
	chunks := strings.SplitN(line, " ", 3)
	if len(chunks) == 1 { // no spaces in line
		return "", 0, nil, fmt.Errorf("invalid line")
	}

	// if there are any empty strings in chunks,
	// that means there were a space at the
	// beginning or a duplicated space somewhere
	for _, c := range chunks {
		if c == "" {
			return "", 0, nil, fmt.Errorf("invalid line")
		}
	}

	name = chunks[0]

	weight, err = parseWeight(chunks[1])
	if err != nil {
		return "", 0, nil, fmt.Errorf("invalid line")
	}

	if len(chunks) == 2 {
		return name, weight, nil, nil
	}

	children, err = parseChildren(chunks[2])
	if err != nil {
		return "", 0, nil, fmt.Errorf("invalid line")
	}

	return name, weight, children, nil
}

// parseWeight returns the weight in s. It panics if s is the empty string.
func parseWeight(s string) (int, error) {
	if s[0] != '(' {
		return 0, fmt.Errorf("invalid line")
	}
	if s[len(s)-1] != ')' {
		return 0, fmt.Errorf("invalid line")
	}
	if s == "()" {
		return 0, fmt.Errorf("invalid line")
	}

	number := s[1 : len(s)-1]
	n, err := strconv.Atoi(number)
	if err != nil {
		return 0, err
	}
	if n <= 0 {
		return 0, fmt.Errorf("weight must be >0, was %d", n)
	}
	return n, nil
}

func parseChildren(s string) ([]string, error) {
	if !strings.HasPrefix(s, "-> ") {
		return nil, fmt.Errorf("invalid line")
	}

	tail := strings.TrimPrefix(s, "-> ")
	children := strings.Split(tail, ", ")

	for _, c := range children {
		if c == "" {
			return nil, fmt.Errorf("invalid line")
		}
	}

	return children, nil
}
