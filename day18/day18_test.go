package day18

import (
	"testing"

	. "github.com/rafael-luigi-bekkema/advent-of-code-2016/util"
)

func TestDay18(t *testing.T) {
	TestEqual(t, 6, day18a("..^^.", 3))
	TestEqual(t, 38, day18a(".^^.^.^^^^", 10))
	input := "^.....^.^^^^^.^..^^.^.......^^..^^^..^^^^..^.^^.^.^" +
		"....^^...^^.^^.^...^^.^^^^..^^.....^.^...^.^.^^.^"
	TestEqual(t, 1974, day18a(input, 40))
	TestEqual(t, 19991126, day18a(input, 400_000))
}
