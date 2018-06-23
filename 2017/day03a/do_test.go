package main_test

import (
	"fmt"
	"testing"

	day03a "github.com/thersanchez/aoc/2017/day03a"
)

func TestDoOK(t *testing.T) {
	for _, tt := range []struct {
		input int
		want  int
	}{
		{input: 1, want: 0},
		{input: 2, want: 1},
		{input: 3, want: 2},
		{input: 4, want: 1},
		{input: 5, want: 2},
		{input: 6, want: 1},
		{input: 7, want: 2},
		{input: 8, want: 1},
		{input: 9, want: 2},
		{input: 10, want: 3},
		{input: 11, want: 2},
		{input: 12, want: 3},
		{input: 13, want: 4},
		{input: 14, want: 3},
		{input: 15, want: 2},
		{input: 16, want: 3},
		{input: 17, want: 4},
		{input: 18, want: 3},
		{input: 19, want: 2},
		{input: 20, want: 3},
		{input: 21, want: 4},
		{input: 22, want: 3},
		{input: 23, want: 2},
		{input: 24, want: 3},
		{input: 25, want: 4},
		{input: 26, want: 5},
		{input: 27, want: 4},
		{input: 28, want: 3},
		{input: 29, want: 4},
		{input: 30, want: 5},
		{input: 31, want: 6},
		{input: 32, want: 5},
		{input: 33, want: 4},
		{input: 34, want: 3},
		{input: 35, want: 4},
		{input: 36, want: 5},
		{input: 37, want: 6},
		{input: 38, want: 5},
		{input: 39, want: 4},
		{input: 40, want: 3},
		{input: 41, want: 4},
		{input: 42, want: 5},
		{input: 43, want: 6},
		{input: 44, want: 5},
		{input: 45, want: 4},
		{input: 46, want: 3},
		{input: 47, want: 4},
		{input: 48, want: 5},
		{input: 49, want: 6},
		{input: 50, want: 7},
	} {
		description := fmt.Sprintf("%d", tt.input)
		t.Run(description, func(t *testing.T) {
			got, err := day03a.Do(tt.input)
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
	for _, input := range []int{
		0, -1, -2, -1000,
	} {
		description := fmt.Sprintf("%d", input)
		t.Run(description, func(t *testing.T) {
			got, err := day03a.Do(input)
			if err == nil {
				t.Errorf("unexpected success, wanted an error, got %d", got)
			}
			if 0 != got {
				t.Errorf("want=0, got=%d", got)
			}
		})
	}
}
