package main

import (
	"errors"
	"fmt"
	"math"
)

// Do solves aoc 2017 day03a.
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

// sidesCenters returns the values of the centers at the four sides of an spiral memory array with side n.
// It returns an error if n<3 or even.
//
// example:
// n=3
// spiral array:
// 5 4 3
// 6 1 2
// 7 8 9
// centers: 2 4 6 8
func sidesCenters(n int) ([4]int, error) {
	if n < 3 {
		return [4]int{}, fmt.Errorf("invalid array side, must be >=3, was %d", n)
	}
	if n%2 == 0 {
		return [4]int{}, fmt.Errorf("invalid array side, must be odd, was %d", n)
	}

	last := n * n
	var centers [4]int
	centers[3] = last - (n / 2)
	centers[2] = centers[3] - (n - 1)
	centers[1] = centers[2] - (n - 1)
	centers[0] = centers[1] - (n - 1)

	return centers, nil
}
