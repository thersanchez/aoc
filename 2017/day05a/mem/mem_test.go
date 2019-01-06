package mem_test

import (
	"fmt"
	"testing"

	"github.com/thersanchez/aoc/2017/day05a/mem"
)

// Checks that negative memory sizes are invalid.
func TestMemNewNegativeSizeError(t *testing.T) {
	t.Parallel()
	for _, size := range []int{
		-1, -3, -1000,
	} {
		size := size
		t.Run(fmt.Sprintf("%d", size), func(t *testing.T) {
			t.Parallel()
			_, err := mem.New(size)
			if err == nil {
				t.Fail()
			}
		})
	}
}

func TestMemNewInitializesToZero(t *testing.T) {
	t.Parallel()
	for _, size := range []int{
		1, 3, 1000,
	} {
		size := size
		t.Run(fmt.Sprintf("%d", size), func(t *testing.T) {
			t.Parallel()
			m, err := mem.New(size)
			if err != nil {
				t.Fatal()
			}
			for i := 0; i < size; i++ {
				v, err := m.Read(i)
				if err != nil {
					t.Errorf("read error: %v", err)
				}
				if v != 0 {
					t.Errorf("want 0, got %d", v)
				}

			}
		})
	}
}

// checks that zero sized memories always returns errors when reading or writing.
func TestMemNewZeroSizeMemory(t *testing.T) {
	t.Parallel()
	m, err := mem.New(0)
	if err != nil {
		t.Fatal()
	}
	_, err = m.Read(0)
	if err == nil {
		t.Errorf("unexpected success reading from 0")
	}
	_, err = m.Read(1)
	if err == nil {
		t.Errorf("unexpected success reading from 1")
	}
	_, err = m.Read(100)
	if err == nil {
		t.Errorf("unexpected success reading from 100")
	}
	err = m.Write(0, 1)
	if err == nil {
		t.Errorf("unexpected success writing to 0")
	}
	err = m.Write(1, 1)
	if err == nil {
		t.Errorf("unexpected success writing to 1")
	}
	err = m.Write(100, 1)
	if err == nil {
		t.Errorf("unexpected success writing to 100")
	}
}

// checks that read returns the values inserted previously by write.
func TestMemReadReturnsDataWrittenWithWrite(t *testing.T) {
	t.Parallel()

	size := 3
	want := 42

	for _, addr := range []int{
		0, 1, 2,
	} {
		addr := addr
		t.Run(fmt.Sprintf("%d", addr), func(t *testing.T) {
			t.Parallel()
			m, err := mem.New(size)
			if err != nil {
				t.Fatal()
			}
			if err = m.Write(addr, want); err != nil {
				t.Fatal()
			}
			got, err := m.Read(addr)
			if err != nil {
				t.Fatal()
			}
			if got != want {
				t.Errorf("want %d, got %d", want, got)
			}
		})
	}
}

// checks that read detects when the given addr is out of bounds.
func TestMemReadDetectsInvalidAddr(t *testing.T) {
	t.Parallel()

	size := 10

	for _, addr := range []int{
		-100, -1, 10, 11, 1000,
	} {
		addr := addr
		t.Run(fmt.Sprintf("%d", addr), func(t *testing.T) {
			t.Parallel()
			m, err := mem.New(size)
			if err != nil {
				t.Fatal()
			}
			if _, err := m.Read(addr); err == nil {
				t.Errorf("unexpected success")
			}
		})
	}
}

// checks that write detects when the given addr is out of bounds.
func TestMemWriteDetectsInvalidAddr(t *testing.T) {
	t.Parallel()

	size := 10

	for _, addr := range []int{
		-100, -1, 10, 11, 1000,
	} {
		addr := addr
		t.Run(fmt.Sprintf("%d", addr), func(t *testing.T) {
			t.Parallel()
			m, err := mem.New(size)
			if err != nil {
				t.Fatal()
			}
			if err := m.Write(addr, 42); err == nil {
				t.Errorf("unexpected success")
			}
		})
	}
}
