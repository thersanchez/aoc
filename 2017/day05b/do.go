package main

import (
	"io"

	"github.com/thersanchez/aoc/2017/day05a/mem"
	"github.com/thersanchez/aoc/2017/day05b/cpu"
	"github.com/thersanchez/aoc/2017/day05b/readchanger"
)

// Do count the number of steps to reach the exit.
func do(r io.Reader) (int, error) {
	m, err := mem.NewFromReader(r)
	if err != nil {
		return 0, err
	}
	return cpu.Run(readchanger.New(m)), nil
}
