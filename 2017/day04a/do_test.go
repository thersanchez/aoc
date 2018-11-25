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
			want:  0,
		}, {
			desc:  "aoc",
			input: aoc,
			want:  2,
		}, {
			desc:  "one valid line, one word",
			input: "aa",
			want:  1,
		}, {
			desc:  "one valid line, two words",
			input: "aa bb",
			want:  1,
		}, {
			desc:  "one invalid line, two words",
			input: "aa aa",
			want:  0,
		}, {
			desc:  "two valid lines, one word",
			input: "aa\nbb",
			want:  2,
		}, {
			desc:  "two valid lines, two words",
			input: "aa bb\naa bb",
			want:  2,
		}, {
			desc:  "one valid line, one invalid",
			input: "aa bb\naa aa",
			want:  1,
		}, {
			desc:  "one invalid line, one valid",
			input: "aa aa\naa bb",
			want:  1,
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
