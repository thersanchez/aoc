package readchanger_test

import (
	"errors"
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
	mm := mockMem{
		read: func(addr int) (int, error) {
			return 0, errors.New("some read error")
		},
	}
	rc := readchanger.NewReadChanger(mm)
	_, err := rc.ReadAndChange(10)
	if err == nil {
		t.Errorf("unexpected success")
	}
}

func TestCanHandleWriteError(t *testing.T) {
	mm := mockMem{
		read: func(addr int) (int, error) {
			return 42, nil
		},
		write: func(addr, value int) error {
			return errors.New("some write error")
		},
	}
	rc := readchanger.NewReadChanger(mm)
	_, err := rc.ReadAndChange(10)
	if err == nil {
		t.Errorf("unexpected success")
	}
}

func TestReadsCorrectly(t *testing.T) {
	want := struct {
		addr  int
		value int
	}{
		addr:  42,
		value: 10,
	}

	mm := mockMem{
		read: func(addr int) (int, error) {
			if addr != want.addr {
				t.Errorf("wrong addr passed to mem, want %d, got %d",
					want.addr, addr)
			}
			return want.value, nil
		},
		write: func(addr, value int) error { return nil },
	}

	rc := readchanger.NewReadChanger(mm)
	got, err := rc.ReadAndChange(want.addr)
	if err != nil {
		t.Error(err)
	}
	if got != want.value {
		t.Errorf("want %d, got %d", want.value, got)
	}
}
