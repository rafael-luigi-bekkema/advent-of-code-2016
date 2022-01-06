package day15

import (
	"testing"

	. "github.com/rafael-luigi-bekkema/advent-of-code-2016/util"
)

func TestDay15(t *testing.T) {
	TestEqual(t, 5, day15a([]string{
		"Disc #1 has 5 positions; at time=0, it is at position 4.",
		"Disc #2 has 2 positions; at time=0, it is at position 1.",
	}))
	TestEqual(t, 122318, day15a(Lines(15)))
	TestEqual(t, 3208583, day15b(Lines(15)))
}
