package main

import "testing"

func TestDo(t *testing.T) {
	for _, tt := range []struct {
		name  string
		input string
		want  int
	}{
		{
			name:  "empty",
			input: "",
			want:  0,
		}, {
			name:  "single number",
			input: "1",
			want:  1,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			got := do(tt.input)
			if tt.want != got {
				t.Errorf("want=%d, got=%d", tt.want, got)
			}
		})
	}
}
