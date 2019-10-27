package readchanger_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/thersanchez/aoc/2017/day05b/readchanger"
)

type mockMem struct {
	read  func(addr int) (int, error)
	write func(addr int, value int) error
}

func (m mockMem) Read(addr int) (int, error) {
	return m.read(addr)
}

func (m mockMem) Write(addr int, value int) error {
	return m.write(addr, value)
}

func TestCanHandleReadError(t *testing.T) {
	t.Parallel()
	mm := mockMem{
		read: func(addr int) (int, error) {
			return 0, errors.New("some read error")
		},
	}
	rc := readchanger.New(mm)
	_, err := rc.ReadAndChange(10)
	if err == nil {
		t.Errorf("unexpected success")
	}
}

func TestCanHandleWriteError(t *testing.T) {
	t.Parallel()
	mm := mockMem{
		read: func(addr int) (int, error) {
			return 42, nil
		},
		write: func(addr, value int) error {
			return errors.New("some write error")
		},
	}
	rc := readchanger.New(mm)
	_, err := rc.ReadAndChange(10)
	if err == nil {
		t.Errorf("unexpected success")
	}
}

func TestReadsCorrectly(t *testing.T) {
	t.Parallel()
	want := struct {
		addr  int
		value int
	}{
		addr:  42,
		value: 10,
	}

	// a memory that complains if it is asked to read
	// from the wrong address.
	// It also counts the number of times it is read in nReads.
	nReads := 0
	mm := mockMem{
		read: func(addr int) (int, error) {
			nReads++
			if addr != want.addr {
				t.Errorf("wrong addr passed to mem, want %d, got %d",
					want.addr, addr)
			}
			return want.value, nil
		},
		write: func(addr, value int) error { return nil },
	}

	rc := readchanger.New(mm)
	got, err := rc.ReadAndChange(want.addr)
	if err != nil {
		t.Error(err)
	}
	if got != want.value {
		t.Errorf("want %d, got %d", want.value, got)
	}
	if nReads != 1 {
		t.Errorf("wrong number of read calls: want 1, got %d", nReads)
	}
}

func TestChangesCorrectly(t *testing.T) {
	t.Parallel()
	tests := []struct {
		addr     int
		original int
		changed  int
	}{
		{addr: 10, original: 0, changed: 1},
		{addr: 11, original: 1, changed: 2},
		{addr: 12, original: 2, changed: 3},
		{addr: 13, original: 3, changed: 2},
		{addr: 14, original: 4, changed: 3},
		{addr: 15, original: 5, changed: 4},
		{addr: 16, original: 6, changed: 5},
		{addr: 17, original: 7, changed: 6},
	}

	for _, test := range tests {
		desc := fmt.Sprintf("%d", test.original)
		test := test
		t.Run(desc, func(t *testing.T) {
			t.Parallel()
			// a memory that complains if it is asked to write
			// the wrong value or on the wrong address.
			// It also counts the number of times it is written in nWrites.
			nWrites := 0
			mm := mockMem{
				read: func(_ int) (int, error) {
					return test.original, nil
				},
				write: func(addr, value int) error {
					nWrites++
					if addr != test.addr {
						t.Errorf("wrong addr: want %d, got %d", test.addr, addr)
					}
					if value != test.changed {
						t.Errorf("wrong value: want %d, got %d", test.changed, value)
					}
					return nil
				},
			}

			rc := readchanger.New(mm)
			_, err := rc.ReadAndChange(test.addr)
			if err != nil {
				t.Error(err)
			}
			if nWrites != 1 {
				t.Errorf("wrong number of write calls: want 1, got %d", nWrites)
			}
		})
	}
}
