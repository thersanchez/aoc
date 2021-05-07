package main

import (
	"fmt"
	"strings"
)

func parseLine(line string) (
	name string, weight int, children []string, err error) {
	chunks := strings.SplitN(line, " ", 3)
	if len(chunks) == 1 { // no spaces in line
		return "", 0, nil, fmt.Errorf("invalid line")
	}

	// if there are any empty strings in chunks,
	// that means there were a space at the
	// beginning or at the end
	for _, c := range chunks {
		if c == "" {
			return "", 0, nil, fmt.Errorf("invalid line")
		}
	}

	name = chunks[0]
	weight, err = decodeWeight(chunks[1])
	if err != nil {
		return "", 0, nil, fmt.Errorf("invalid line")
	}

	return name, weight, nil, nil
}

func decodeWeight(s string) (int, error) {
	return 42, nil
}
