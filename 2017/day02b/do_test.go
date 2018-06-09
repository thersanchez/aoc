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
		{input: "5 9 2 8\n9 4 7 3\n3 8 6 5", want: 9},
		{input: "4 2", want: 2},
		{input: "4 2 3", want: 2},
		{input: "4 3 2 ", want: 2},
		{input: "3 4 2 ", want: 2},
		{input: "2 4", want: 2},
		{input: "2 4 3", want: 2},
		{input: "2 3 4 ", want: 2},
		{input: "3 2 4", want: 2},
		{input: "4 2\n6 2", want: 5},
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
