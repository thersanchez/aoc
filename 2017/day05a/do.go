package main

import (
	"io"

	"github.com/thersanchez/aoc/2017/day05a/cpu"
	"github.com/thersanchez/aoc/2017/day05a/mem"
)

// Do count the number of steps to reach the exit.
func do(r io.Reader) int {
	var m mem.AutoInc
	load(m, r)
	return cpu.Run(m)
}

func load(m mem.AutoInc, r io.Reader) {

}
