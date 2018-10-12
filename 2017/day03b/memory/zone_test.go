package memory_test

import (
	"testing"

	"github.com/thersanchez/aoc/2017/day03b/memory"
)

func TestZoneString(t *testing.T) {
	t.Parallel()
	for _, tt := range []struct {
		zone memory.Zone
		want string
	}{
		{memory.Top, "Top"},
		{memory.Bottom, "Bottom"},
		{memory.Left, "Left"},
		{memory.Right, "Right"},
	} {
		tt := tt
		t.Run(tt.want, func(t *testing.T) {
			t.Parallel()
			got := tt.zone.String()
			if got != tt.want {
				t.Errorf("want=%s, got=%s", tt.want, got)
			}
		})
	}
}

func TestZoneStringPanic(t *testing.T) {
	t.Parallel()
	defer func() {
		if r := recover(); r == nil {
			t.Fatal("no panic detected")
		}
	}()
	zone := memory.Zone(-12)
	_ = zone.String()
}
