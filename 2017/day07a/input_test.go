package main

import (
	"sort"
	"testing"
)

func TestParseNodeOK(t *testing.T) {
	t.Parallel()

	type want struct {
		name     string
		weight   int
		children []string
	}

	subtests := []struct {
		name string
		line string
		want want
	}{
		{
			name: "no children",
			line: "a (42)",
			want: want{
				name:   "a",
				weight: 42,
			},
		},
	}

	for _, test := range subtests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			name, weight, children, err := parseNode(test.line)
			if err != nil {
				t.Fatalf("error parsing line: %v", err)
			}

			if test.want.name != name {
				t.Errorf("wrong name: want %q, got %q",
					test.want.name, name)
			}

			if test.want.weight != weight {
				t.Errorf("wrong weight: want %d, got %d",
					test.want.weight, weight)
			}

			if !equalSlices(test.want.children, children) {
				t.Errorf("wrong children: want %q, got %q",
					test.want.children, children)
			}
		})
	}

}

// equalSlices returns true if two slices have the same elements, in any order.
// Calling this function might alter the order of the elements in the slices.
func equalSlices(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	sort.Strings(a)
	sort.Strings(b)

	for i, ea := range a {
		if ea != b[i] {
			return false
		}
	}

	return true
}
