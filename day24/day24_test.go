package day24

import (
	"testing"

	. "github.com/rafael-luigi-bekkema/advent-of-code-2016/util"
)

func TestDay24(t *testing.T) {
	testMap := "###########\n" +
		"#0.1.....2#\n" +
		"#.#######.#\n" +
		"#4.......3#\n" +
		"###########\n"
	TestEqual(t, 14, day24a(testMap, false))
	TestEqual(t, 498, day24a(Input(24), false))
	TestEqual(t, 804, day24a(Input(24), true))
}
