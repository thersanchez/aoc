package main

import (
	"strings"
	"testing"
)

func TestDo(t *testing.T) {
	t.Parallel()

	subtests := []struct {
		name   string
		reader string
		want   int
		err    bool
	}{
		{
			name:   "empty",
			reader: "",
			err:    true,
		},
		{
			name:   "malformed",
			reader: "foo",
			err:    true,
		},
		{
			name:   "malformed 2",
			reader: "1\tfoo\t6",
			err:    true,
		},
		{
			name:   "one",
			reader: "1",
			want:   1,
		},
		{
			name:   "one plus zeros",
			reader: "1\t0\t0",
			want:   3,
		},
		{
			name:   "two plus zeros",
			reader: "2\t0\t0\t0",
			want:   9,
		},
		{
			name:   "two plus zeros",
			reader: "4\t10\t4\t1\t8\t4\t9\t14\t5\t1\t14\t15\t0\t15\t3\t5",
			want:   12841,
		},
	}

	for _, test := range subtests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			r := strings.NewReader(test.reader)
			got, err := do(r)

			if test.err {
				if err == nil {
					t.Errorf("unexpected success: %v", got)
				}
			} else {
				if err != nil {
					t.Errorf("got error: %v", err)
				}

				if got != test.want {
					t.Errorf("want %d, got %d", test.want, got)
				}
			}
		})
	}

}
