package main

import (
	"io"

	"github.com/thersanchez/aoc/2017/day05a/mem"
)

// Do count the number of steps to reach the exit.
func do(r io.Reader) int {
	var m mem.AutoInc
	load(m, r)
	return cpu(m)
}

func load(m mem.AutoInc, r io.Reader) {

}

func cpu(m mem.AutoInc) int {
	var pc, ir, steps int
	var err error

	for {
		// Stores the Mem Value at PC Address into IR
		ir, err = m.Read(pc)
		if err != nil {
			return steps
		}

		// Jump
		pc = pc + ir
	}
}
