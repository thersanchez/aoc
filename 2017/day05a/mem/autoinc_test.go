package mem_test

import (
	"errors"
	"testing"

	"github.com/thersanchez/aoc/2017/day05a/mem"
)

type mockMem struct {
	readCallCount  int
	writeCallCount int
	addr           int
	ok             bool
	t              *testing.T
}

func newMockMem(t *testing.T) *mockMem {
	return &mockMem{
		t: t,
	}
}

func (m *mockMem) Read(addr int) (int, error) {
	m.readCallCount++
	if m.readCallCount != 1 {
		m.t.Fatal("extra read call")
	}
	if m.writeCallCount != 0 {
		m.t.Fatal("read call after calling write")
	}
	m.addr = addr
	return 42, nil
}

func (m *mockMem) Write(addr, value int) error {
	m.writeCallCount++
	if m.writeCallCount != 1 {
		m.t.Fatal("extra write call")
	}
	if m.readCallCount == 0 {
		m.t.Fatal("write call before calling read")
	}
	if m.addr != addr {
		m.t.Fatalf("wrong addr, want %d, got %d", m.addr, addr)
	}
	if value != 43 {
		m.t.Fatalf("wrong value, want %d, got %d", 43, value)
	}
	return nil
}

// Checks that AutoInc writes the read value +1
func TestAutoIncHappyPath(t *testing.T) {
	rw := newMockMem(t)
	sut := mem.NewAutoInc(rw)
	got, err := sut.ReadAndInc(0)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got != 42 {
		t.Errorf("want %d, got %d", 42, got)
	}
}

type mockMemReadError struct{}

func (m *mockMemReadError) Read(addr int) (int, error) {
	return 0, errors.New("read error")
}

func (m *mockMemReadError) Write(addr, value int) error {
	return nil
}

// Checks that AutoInc handles ReadWriter.Read errors.
func TestAutoIncReadError(t *testing.T) {
	rw := &mockMemReadError{}
	sut := mem.NewAutoInc(rw)
	got, err := sut.ReadAndInc(0)
	if err == nil {
		t.Fatalf("unexpected success, got %d", got)
	}
}

type mockMemWriteError struct{}

func (m *mockMemWriteError) Read(addr int) (int, error) {
	return 42, nil
}

func (m *mockMemWriteError) Write(addr, value int) error {
	return errors.New("write error")
}

// Checks that AutoInc handles ReadWriter.Write errors.
func TestAutoIncWriteError(t *testing.T) {
	rw := &mockMemWriteError{}
	sut := mem.NewAutoInc(rw)
	got, err := sut.ReadAndInc(0)
	if err == nil {
		t.Fatalf("unexpected success, got %d", got)
	}
}
