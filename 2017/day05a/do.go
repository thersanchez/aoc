package main

import (
	"io"

	"github.com/thersanchez/aoc/2017/day05a/cpu"
	"github.com/thersanchez/aoc/2017/day05a/mem"
)

// Do count the number of steps to reach the exit.
func do(r io.Reader) (int, error) {
	m, err := mem.NewFromReader(r)
	if err != nil {
		return 0, err
	}
	return cpu.Run(mem.NewAutoInc(m)), nil
}
