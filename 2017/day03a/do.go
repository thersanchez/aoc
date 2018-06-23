package main

import (
	"fmt"
	"math"
)

// Do solves aoc 2017 day03a.
func Do(input int) (int, error) {
	if input == 1 {
		return 0, nil
	}
	side, err := minArraySideContaining(input)
	if err != nil {
		return 0, err
	}
	centers, err := sidesCenters(side)
	if err != nil {
		return 0, err
	}
	distToCenter := minDistance(input, centers)

	return manhattanDistance(distToCenter, side), nil
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
		asInt++
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

// minDistance calculates the shortest distance between a given number and 4 other numbers.
func minDistance(x int, others [4]int) int {
	dists := [4]int{}
	for i, o := range others {
		dists[i] = abs(x - o)
	}
	minDist := dists[0]
	for _, d := range dists[1:] {
		if d < minDist {
			minDist = d
		}
	}
	return minDist
}

// abs calculates the absolute value of an integer.
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func manhattanDistance(distToCenter, side int) int {
	return distToCenter + side/2
}
