package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetSubtract(t *testing.T) {
	t.Parallel()

	subtests := []struct {
		desc      string
		minuend   map[string]struct{}
		subtraend map[string]struct{}
		want      map[string]struct{}
	}{
		{
			desc:      "empty",
			minuend:   map[string]struct{}{},
			subtraend: map[string]struct{}{},
			want:      map[string]struct{}{},
		},
		{
			desc:      "one minus nothing",
			minuend:   map[string]struct{}{"a": struct{}{}},
			subtraend: map[string]struct{}{},
			want:      map[string]struct{}{"a": struct{}{}},
		},
		{
			desc:      "set minus itself",
			minuend:   map[string]struct{}{"a": struct{}{}},
			subtraend: map[string]struct{}{"a": struct{}{}},
			want:      map[string]struct{}{},
		},
		{
			desc:      "one set minus different set",
			minuend:   map[string]struct{}{"a": struct{}{}},
			subtraend: map[string]struct{}{"b": struct{}{}},
			want:      map[string]struct{}{"a": struct{}{}},
		},
		{
			desc:      "mixed",
			minuend:   map[string]struct{}{"a": struct{}{}, "b": struct{}{}},
			subtraend: map[string]struct{}{"b": struct{}{}, "c": struct{}{}},
			want:      map[string]struct{}{"a": struct{}{}},
		},
	}

	for _, test := range subtests {
		test := test
		t.Run(test.desc, func(t *testing.T) {
			t.Parallel()
			got := setSubtract(test.minuend, test.subtraend)
			assert.Equal(t, test.want, got)
		})
	}
}
