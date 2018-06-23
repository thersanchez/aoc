package main

import (
	"errors"
	"fmt"
	"math"
)

func Do(input int) (int, error) {
	return 0, errors.New("not implemented")
}

// minArraySideContaining returns the side of the minimum array
// containing the number x or an error if x<1.
func minArraySideContaining(x int) (int, error) {
	if x < 1 {
		return 0, fmt.Errorf("invalid input, must be >=1, was %d", x)
	}
	root := math.Sqrt(float64(x))
	ceil := math.Ceil(root)
	asInt := int(math.Round(ceil))
	if asInt%2 == 0 {
		asInt += 1
	}
	return asInt, nil
}
