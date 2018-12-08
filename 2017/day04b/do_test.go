package main

import (
	"strings"
	"testing"
)

// an valid passphrase
var v = "aa bb\n"

// an invalid passphrase
var i = "aa aa\n"

func TestDo(t *testing.T) {
	t.Parallel()
	for _, test := range []struct {
		desc  string
		input string
		want  int
	}{
		{
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
			desc:  "empty",
			input: "",
			want:  0, // no passphrases means zero valid
		}, {
			desc:  "one passphrase, zero valid",
			input: i,
			want:  0,
		}, {
			desc:  "two passphrases, zero valid",
			input: i + i,
			want:  0,
		}, {
			desc:  "five passphrases, zero valid",
			input: i + i + i + i + i,
			want:  0,
		}, {
			desc:  "five passphrases, one empty, zero valid",
			input: i + i + "\n" + i + i,
			want:  1, // an empty passphrase is valid
		}, {
			desc:  "five passphrases, all empty, zero valid",
			input: "\n\n\n\n",
			want:  4, // all empty passphrases are valid
		}, {
			desc:  "one passphrase, one valid",
			input: v,
			want:  1,
		}, {
			desc:  "one valid, one invalid",
			input: v + i,
			want:  1,
		}, {
			desc:  "one invalid, one valid",
			input: i + v,
			want:  1,
		}, {
			desc:  "viiii",
			input: v + i + i + i + i,
			want:  1,
		}, {
			desc:  "iivii",
			input: i + i + v + i + i,
			want:  1,
		}, {
			desc:  "iiiiv",
			input: i + i + i + i + v,
			want:  1,
		}, {
			desc:  "vv",
			input: v + v,
			want:  2,
		}, {
			desc:  "vivii",
			input: v + i + v + i + i,
			want:  2,
		}, {
			desc:  "iiivv",
			input: i + i + i + v + v,
			want:  2,
		}, {
			desc:  "vvvvv",
			input: v + v + v + v + v,
			want:  5,
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

func TestIsValidPassphrase(t *testing.T) {
	t.Parallel()
	for _, test := range []struct {
		desc  string
		input string
		want  bool
	}{
		{
			desc:  "empty", // empty passphrases are valid
			input: "",
			want:  true,
		}, {
			desc:  "one word", // a single word is always valid
			input: "aa",
			want:  true,
		}, {
			desc:  "two non-anagrams words",
			input: "aa bb",
			want:  true,
		}, {
			desc:  "two equal words", // equal words are anagrams
			input: "aa aa",
			want:  false,
		}, {
			desc:  "two anagrams",
			input: "ab ba",
			want:  false,
		}, {
			desc:  "two anagrams among many",
			input: "aa XY cc dd YX ee",
			want:  false,
		},
	} {
		test := test
		t.Run(test.desc, func(t *testing.T) {
			t.Parallel()
			got := isValidPassphrase(strings.NewReader(test.input))
			if got != test.want {
				t.Errorf("want %t, got %t", test.want, got)
			}
		})
	}
}

func TestAreAnagrams(t *testing.T) {
	t.Parallel()
	for _, test := range []struct {
		desc string
		a, b string
		want bool
	}{
		{
			desc: "same word",
			a:    "aa",
			b:    "aa",
			want: true,
		}, {
			desc: "no anagrams",
			a:    "aa",
			b:    "bb",
			want: false,
		}, {
			desc: "anagrams",
			a:    "ab",
			b:    "ba",
			want: true,
		}, {
			desc: "same letters, different letter frequency",
			a:    "ana",
			b:    "ann",
			want: false,
		},
	} {
		test := test
		t.Run(test.desc, func(t *testing.T) {
			t.Parallel()
			got := areAnagrams(test.a, test.b)
			if got != test.want {
				t.Errorf("want %t, got %t", test.want, got)
			}
		})
	}
}
