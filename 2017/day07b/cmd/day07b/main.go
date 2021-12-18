package main

import (
	"fmt"

	"github.com/thersanchez/aoc/2017/day07a"
	"github.com/thersanchez/aoc/2017/day07b"
)

func main() {
	n := day07b.Node{
		Node: day07a.Node{
			Weight: 42,
		},
		TotalWeight: 30,
	}

	fmt.Println(n.Weight)
	fmt.Println(n.TotalWeight)
}
