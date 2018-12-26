package main

import (
	"strings"
	"testing"
)

func TestDo(t *testing.T) {
	t.Parallel()
	for _, test := range []struct {
		desc  string
		input string
		want  int
	}{
		{
			desc:  "TODO",
			input: "1\n0",
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
