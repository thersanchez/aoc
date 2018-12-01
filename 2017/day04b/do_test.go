package main

import (
	"strings"
	"testing"
)

var aoc = `aa bb cc dd ee
aa bb cc dd aa
aa bb cc dd aaa`

func TestDo(t *testing.T) {
	t.Parallel()
	for _, test := range []struct {
		desc  string
		input string
		want  int
	}{
		{
			desc:  "empty",
			input: "",
			want:  0, // 1?
		}, {
			desc:  "aoc1",
			input: "abcde fghij",
			want:  1,
		}, {
			desc:  "aoc2",
			input: "abcde xyz ecdab",
			want:  0,
		}, {
			desc:  "aoc3",
			input: "a ab abc abd abf abj",
			want:  1,
		}, {
			desc:  "aoc4",
			input: "iiii oiii ooii oooi oooo",
			want:  1,
		}, {
			desc:  "aoc5",
			input: "oiii ioii iioi iiio",
			want:  0,
		}, {
			desc:  "aoc2",
			input: "abcde xyz ecdab",
			want:  0,
		},
	} {
		test := test
		t.Run(test.desc, func(t *testing.T) {
			t.Parallel()
			got := do(strings.NewReader(test.input))
			if got != test.want {
				t.Errorf("want %d, got %d", test.want, got)
			}
		})
	}
}

func TestHasDuplicatedWords(t *testing.T) {
	t.Parallel()
	for _, test := range []struct {
		desc  string
		input string
		want  bool
	}{
		{
			desc:  "empty",
			input: "",
			want:  false,
		}, {
			desc:  "one word",
			input: "aa",
			want:  false,
		}, {
			desc:  "two different words",
			input: "aa bb",
			want:  false,
		}, {
			desc:  "two equal words",
			input: "aa aa",
			want:  true,
		}, {
			desc:  "many non duplicated words",
			input: "aa bb cc dd ee aaa bbb aabb bbaa",
			want:  false,
		}, {
			desc:  "duplicated word among many",
			input: "aa bb cc dd ee aaa bbb cc aabb bbaa",
			want:  true,
		},
	} {
		test := test
		t.Run(test.desc, func(t *testing.T) {
			t.Parallel()
			got := hasDuplicatedWords(strings.NewReader(test.input))
			if got != test.want {
				t.Errorf("want %t, got %t", test.want, got)
			}
		})
	}
}
