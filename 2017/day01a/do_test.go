package main

import "testing"

func TestDo(t *testing.T) {
	for _, tt := range []struct {
		input string
		want  int
	}{
		{input: "1122", want: 3},
		{input: "1111", want: 4},
		{input: "1234", want: 0},
		{input: "91212129", want: 9},
		{input: "", want: 0},
		{input: "1", want: 0},
		{input: "a", want: 0},
		{input: "11", want: 1 + 1},
		{input: "12", want: 0},
		{input: "1a", want: 0},
		{input: "111", want: 1 + 1 + 1},
		{input: "123", want: 0},
		{input: "112", want: 1},
		{input: "121", want: 1},
		{input: "211", want: 1},
		{input: "a111", want: 1 + 1 + 1},
		{input: "a123", want: 0},
		{input: "a112", want: 1},
		{input: "a121", want: 1},
		{input: "a211", want: 1},
		{input: "1a11", want: 1 + 1 + 1},
		{input: "1a23", want: 0},
		{input: "1a12", want: 1},
		{input: "1a21", want: 1},
		{input: "2a11", want: 1},
		{input: "11a1", want: 1 + 1 + 1},
		{input: "12a3", want: 0},
		{input: "11a2", want: 1},
		{input: "12a1", want: 1},
		{input: "21a1", want: 1},
		{input: "111a", want: 1 + 1 + 1},
		{input: "123a", want: 0},
		{input: "112a", want: 1},
		{input: "121a", want: 1},
		{input: "211a", want: 1},
		{input: "911422279", want: 1 + 2 + 2 + 9},
		{input: "91asds14bb2ccsd2as27aa9fdwwtyg", want: 1 + 2 + 2 + 9},
		{input: "┴4", want: 0}, // `┴` is U+2534 and `4` is U+34.
	} {
		t.Run(tt.input, func(t *testing.T) {
			got := do(tt.input)
			if tt.want != got {
				t.Errorf("want=%d, got=%d", tt.want, got)
			}
		})
	}
}
