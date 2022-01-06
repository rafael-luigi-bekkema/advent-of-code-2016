package day12

import (
	"testing"

	. "github.com/rafael-luigi-bekkema/advent-of-code-2016/util"
)

func TestDay12(t *testing.T) {
	TestEqual(t, 42, day12a([]string{"cpy 41 a",
		"inc a",
		"inc a",
		"dec a",
		"jnz a 2",
		"dec a",
	}))
	TestEqual(t, 6, day12a([]string{"cpy 5 a",
		"jnz a 2",
		"inc a",
		"inc a",
	}))
	TestEqual(t, 317993, day12a(Lines(12)))
	TestEqual(t, 9227647, day12b(Lines(12)))
}
