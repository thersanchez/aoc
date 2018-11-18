package main

import (
	"fmt"

	"github.com/thersanchez/aoc/2017/day03b/memory"
)

func main() {
	fmt.Println(Do(368078))
}

// Do solves day03b.
func Do(target int) int {
	for p := 1; ; p++ {
		v := memory.CalculateValue(p)
		if v > target {
			return v
		}
	}
}
