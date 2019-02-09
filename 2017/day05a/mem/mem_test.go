package mem_test

import (
	"errors"
	"fmt"
	"io"
	"strings"
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
				t.Fatalf("cannot create mem: %v", err)
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
		t.Fatalf("cannot create mem: %v", err)
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
				t.Fatalf("cannot create mem: %v", err)
			}
			if err = m.Write(addr, want); err != nil {
				t.Fatalf("cannot write mem: %v", err)
			}
			got, err := m.Read(addr)
			if err != nil {
				t.Fatalf("cannot read mem: %v", err)
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
				t.Fatalf("cannot create mem: %v", err)
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
				t.Fatalf("cannot create mem: %v", err)
			}
			if err := m.Write(addr, 42); err == nil {
				t.Errorf("unexpected success")
			}
		})
	}
}

func TestSize(t *testing.T) {
	t.Parallel()

	for _, size := range []int{
		0, 1, 2, 1000,
	} {
		size := size
		t.Run(fmt.Sprintf("%d", size), func(t *testing.T) {
			t.Parallel()
			m, err := mem.New(size)
			if err != nil {
				t.Fatalf("cannot create mem: %v", err)
			}
			if got := m.Size(); got != size {
				t.Errorf("wrong size, want %d, got %d", size, got)
			}
		})
	}
}

func TestNewFromReaderCorrectSize(t *testing.T) {
	t.Parallel()

	for _, test := range []struct {
		desc        string
		size        int
		readerAsStr string
	}{
		{
			desc:        "0 no eol",
			size:        0,
			readerAsStr: "",
		},
		// this reader make the mem ctor fail because
		// the first line is empty and it expects a number.
		/*
			{
				desc:        "0 eol",
				size:        0,
				readerAsStr: "\n",
			},
		*/
		{
			desc:        "1 no eol",
			size:        1,
			readerAsStr: "42",
		},
		{
			desc:        "1 eol",
			size:        1,
			readerAsStr: "42\n",
		},
		{
			desc:        "2 no eol",
			size:        2,
			readerAsStr: "42\n13",
		},
		{
			desc:        "2 eol",
			size:        2,
			readerAsStr: "42\n13\n",
		},
		{
			desc:        "7 no eol",
			size:        7,
			readerAsStr: "1\n2\n3\n4\n5\n6\n7",
		},
		{
			desc:        "7 eol",
			size:        7,
			readerAsStr: "1\n2\n3\n4\n5\n6\n7\n",
		},
	} {
		test := test
		t.Run(test.desc, func(t *testing.T) {
			t.Parallel()
			m, err := mem.NewFromReader(
				strings.NewReader(test.readerAsStr))
			if err != nil {
				t.Fatalf("cannot create mem: %v", err)
			}
			if got := m.Size(); got != test.size {
				t.Errorf("wrong size, want %d, got %d", test.size, got)
			}
		})
	}
}

func TestNewFromReaderCorrectData(t *testing.T) {
	t.Parallel()

	reader := strings.NewReader("42\n0\n1")
	m, err := mem.NewFromReader(reader)
	if err != nil {
		t.Fatalf("cannot create mem: %v", err)
	}

	want := 42
	got, err := m.Read(0)
	if err != nil {
		t.Fatalf("cannot read from mem: %v", err)
	}
	if got != want {
		t.Errorf("wrong data, want %d, got %d", want, got)
	}

	want = 0
	got, err = m.Read(1)
	if err != nil {
		t.Fatalf("cannot read from mem: %v", err)
	}
	if got != want {
		t.Errorf("wrong data, want %d, got %d", want, got)
	}

	want = 1
	got, err = m.Read(2)
	if err != nil {
		t.Fatalf("cannot read from mem: %v", err)
	}
	if got != want {
		t.Errorf("wrong data, want %d, got %d", want, got)
	}
}

func TestNewFromReaderError(t *testing.T) {
	t.Parallel()

	for _, test := range []struct {
		desc   string
		reader io.Reader
	}{
		{
			desc: "reader returns a read error",
			reader: errorReader{
				error: errors.New("read error"),
			},
		},
		{
			desc:   "invalid data",
			reader: strings.NewReader("42\nnot a number\n13"),
		},
		{
			desc:   "empty lines",
			reader: strings.NewReader("42\n\n13"),
		},
		{
			desc:   "first line is empty",
			reader: strings.NewReader("\n42\n13"),
		},
		{
			desc:   "just an empty line",
			reader: strings.NewReader("\n"),
		},
	} {
		test := test
		t.Run(test.desc, func(t *testing.T) {
			t.Parallel()
			_, err := mem.NewFromReader(test.reader)
			if err == nil {
				t.Fatalf("unexpected success")
			}
		})
	}
}

type errorReader struct {
	error
}

func (e errorReader) Read(p []byte) (n int, err error) {
	return 0, e.error
}
