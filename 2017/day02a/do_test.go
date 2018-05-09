package main

import (
	"fmt"
	"testing"
)

func TestDoOK(t *testing.T) {
	for _, tt := range []struct {
		input string
		want  int
	}{
		{input: "5 1 9 5\n7 5 3\n2 4 6 8", want: 18},
		{input: "0", want: 0},
		{input: "1", want: 0},
		{input: "1234", want: 0},
		{input: "1 1 1 1", want: 0},
		{input: "1 2 3 4", want: 3},
		{input: "4 3 2 1", want: 3},
		{input: "1 2 3 4\n12 12 12\n10 9 8 7 1 45 6 8", want: 47},
		{input: "12\t34", want: 22},
		{input: "1 2\t3 4", want: 3},
	} {
		t.Run(fmt.Sprintf("%q", tt.input), func(t *testing.T) {
			got, err := do(tt.input)
			if err != nil {
				t.Errorf("want nil error, got=%q", err)
			}
			if tt.want != got {
				t.Errorf("want=%d, got=%d", tt.want, got)
			}
		})
	}
}

func TestDoError(t *testing.T) {
	for _, tt := range []string{
		"error",
		"",
		"....",
		"a",
		"1 2 a",
		"1 2\n3 4\nerror\n5 6",
		"1 3\n",
		"1 3\n\n",
		"\n1 3",
		"\n\n1 3",
		"1\n\n3",
	} {
		t.Run(fmt.Sprintf("%q", tt), func(t *testing.T) {
			got, err := do(tt)
			if err == nil {
				t.Errorf("expected an error, got nil")
			}
			if got != 0 {
				t.Errorf("want=0, got=%d", got)
			}
		})
	}
}
