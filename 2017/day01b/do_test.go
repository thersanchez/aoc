package main

import "testing"

func TestDo(t *testing.T) {
	for _, tt := range []struct {
		input string
		want  int
	}{
		{input: "1212", want: 6},
		{input: "1221", want: 0},
		{input: "123425", want: 4},
		{input: "123123", want: 12},
		{input: "12131415", want: 4},
		{input: "11", want: 1 + 1},
		{input: "12", want: 0},
		{input: "1111", want: 4},
		{input: "1234", want: 0},
		{input: "1214", want: 2},
		{input: "1232", want: 4},
		{input: "111111", want: 6},
		{input: "123456", want: 0},
		{input: "123153", want: 8},
	} {
		t.Run(tt.input, func(t *testing.T) {
			got := do(tt.input)
			if tt.want != got {
				t.Errorf("want=%d, got=%d", tt.want, got)
			}
		})
	}
}
