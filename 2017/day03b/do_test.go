package main_test

import (
	"fmt"
	"testing"

	day03b "github.com/thersanchez/aoc/2017/day03b"
)

func TestDoOK(t *testing.T) {
	for _, tt := range []struct {
		input int
		want  int
	}{
		{input: -10, want: 1},
		{input: -1, want: 1},
		{input: 0, want: 1},
		{input: 1, want: 2},
		{input: 2, want: 4},
		{input: 3, want: 4},
		{input: 4, want: 5},
		{input: 5, want: 10},
		{input: 6, want: 10},
		{input: 7, want: 10},
		{input: 8, want: 10},
		{input: 9, want: 10},
		{input: 10, want: 11},
		{input: 11, want: 23},
		{input: 12, want: 23},
		{input: 13, want: 23},
		{input: 23, want: 25},
		{input: 24, want: 25},
		{input: 25, want: 26},
		{input: 26, want: 54},
		{input: 53, want: 54},
		{input: 54, want: 57},
		{input: 56, want: 57},
		{input: 57, want: 59},
		{input: 58, want: 59},
		{input: 59, want: 122},
		{input: 121, want: 122},
		{input: 122, want: 133},
		{input: 132, want: 133},
		{input: 133, want: 142},
		{input: 141, want: 142},
		{input: 142, want: 147},
		{input: 146, want: 147},
		{input: 147, want: 304},
		{input: 303, want: 304},
		{input: 304, want: 330},
		{input: 330, want: 351},
		{input: 351, want: 362},
		{input: 362, want: 747},
		{input: 747, want: 806},
		{input: 806, want: 880},
		{input: 880, want: 931},
		{input: 931, want: 957},
		{input: 37408, want: 39835},
		{input: 368078, want: 369601},
	} {
		description := fmt.Sprintf("%d", tt.input)
		t.Run(description, func(t *testing.T) {
			got := day03b.Do(tt.input)
			if tt.want != got {
				t.Errorf("want=%d, got=%d", tt.want, got)
			}
		})
	}
}
