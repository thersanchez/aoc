package cpu_test

import (
	"errors"
	"testing"

	"github.com/thersanchez/aoc/2017/day05a/cpu"
)

// mock satisfies cpu.AutoIncMem by running the function
// stored in readAndInc.
type mock struct {
	readAndInc func(int) (int, error)
}

func (m mock) ReadAndInc(addr int) (int, error) {
	return m.readAndInc(addr)
}

func TestRunHaltsWhenJumpingOutOfMem(t *testing.T) {
	// an AutoIncMem that returns that you are out of mem and complains
	// if you read again.
	callCount := 0
	mem := mock{
		readAndInc: func(int) (int, error) {
			if callCount != 0 {
				t.Error("CPU didn't halt")
			}
			callCount++
			return 0, errors.New("any address is out of mem")
		},
	}
	_ = cpu.Run(mem)
}
