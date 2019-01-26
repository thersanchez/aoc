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
			desc:  "empty memory",
			input: "",
			want:  0,
		},
		{
			desc:  "1 jump",
			input: "1",
			want:  1,
		},
		{
			desc:  "jump forth and back",
			input: "1\n-1",
			want:  3,
		},
	} {
		test := test
		t.Run(test.desc, func(t *testing.T) {
			t.Parallel()
			got, err := do(strings.NewReader(test.input))
			if err != nil {
				t.Fatal(err)
			}
			if got != test.want {
				t.Errorf("want %d, got %d", test.want, got)
			}
		})
	}
}

func TestDoError(t *testing.T) {
	t.Parallel()
	input := "1\n3\nerror\n5"
	_, err := do(strings.NewReader(input))
	if err == nil {
		t.Errorf("unexpected success")
	}
}
