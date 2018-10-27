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
	mem := []int{1}
	for p := 1; ; p++ {
		v := memory.CalculateValue(p)
		mem = append(mem, v)
		if v > target {
			return v
		}
	}
}
