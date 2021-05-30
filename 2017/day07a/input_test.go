package main

import (
	"sort"
	"testing"
)

func TestParseLineError(t *testing.T) {
	t.Parallel()
	subtests := []struct {
		name string
		line string
	}{
		{
			name: "empty",
			line: "",
		}, {
			name: "only name",
			line: "llyhqfe",
		}, {
			name: "only name with space",
			line: "llyhqfe ",
		}, {
			name: "only weight",
			line: "(21)",
		}, {
			name: "only weight, with space",
			line: " (21)",
		}, {
			name: "weight has no brackets",
			line: "a 42",
		}, {
			name: "weight has no open bracket",
			line: "a 42)",
		}, {
			name: "weight has no close bracket",
			line: "a (42",
		}, {
			name: "weight is not a number",
			line: "a (b)",
		}, {
			name: "weight is a float",
			line: "a (3.5)",
		}, {
			name: "weight is negative",
			line: "a (-42)",
		}, {
			name: "only weight, with spaces",
			line: " (21) ",
		}, {
			name: "garbage after weight",
			line: "a (0) b",
		}, {
			name: "arrow but no children",
			line: "a (0) -> ",
		}, {
			name: "comma at the end",
			line: "a (0) -> b,",
		}, {
			name: "comma at the end, with space",
			line: "a (0) -> b, ",
		},
	}

	for _, test := range subtests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			_, _, _, err := ParseLine(test.line)
			if err == nil {
				t.Errorf("unexpected success")
			}

		})
	}
}

func TestParseLineOK(t *testing.T) {
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
			line: "llyhqfe (21)",
			want: want{
				name:   "llyhqfe",
				weight: 21,
			},
		},
		{
			name: "with children",
			line: "vpbdpfm (74) -> ndegtj, wnwxs",
			want: want{
				name:     "vpbdpfm",
				weight:   74,
				children: []string{"ndegtj", "wnwxs"},
			},
		},
	}

	for _, test := range subtests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			name, weight, children, err := ParseLine(test.line)
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
