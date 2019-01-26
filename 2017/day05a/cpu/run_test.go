package cpu_test

import (
	"errors"
	"fmt"
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
	t.Parallel()
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

func TestRunStartsAtZero(t *testing.T) {
	t.Parallel()
	want := 0
	mem := mock{
		readAndInc: func(got int) (int, error) {
			if got != want {
				t.Errorf("want %d, got %d", want, got)
			}
			return 0, errors.New("stop now")
		},
	}
	_ = cpu.Run(mem)
}

func TestRunJumpsCorrectly(t *testing.T) {
	t.Parallel()
	// the desired jump address
	want := 15
	numReads := 0
	mem := mock{
		readAndInc: func(got int) (int, error) {
			defer func() { numReads++ }()
			switch numReads {
			case 0:
				return want, nil
			case 1:
				if got != want {
					t.Fatalf("want %d, got %d", want, got)
				}
				return 0, errors.New("out of memory")
			default:
				t.Fatal("extra mem read")
				return 0, errors.New("unreachable")
			}
		},
	}
	_ = cpu.Run(mem)
}

func TestRunCountsJumpsCorrectly(t *testing.T) {
	t.Parallel()

	for _, numJumps := range []int{
		0, 1, 2, 10, 1000,
	} {
		numJumps := numJumps
		t.Run(fmt.Sprintf("%d", numJumps), func(t *testing.T) {
			t.Parallel()

			callCount := 0
			mem := mock{
				readAndInc: func(got int) (int, error) {
					callCount++
					// stop the CPU once it has done numJumps jumps
					if callCount > numJumps {
						return 0, errors.New("stop here")
					}
					// any value will do
					return 42, nil
				},
			}
			got := cpu.Run(mem)

			if got != numJumps {
				t.Errorf("want %d, got %d", numJumps, got)
			}
		})
	}
}
