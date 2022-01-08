package day23

import (
	"testing"

	. "github.com/rafael-luigi-bekkema/advent-of-code-2016/util"
)

func TestDay23(t *testing.T) {
	TestEqual(t, 3, assembunny([]string{
		"cpy 2 a",
		"tgl a",
		"tgl a",
		"tgl a",
		"cpy 1 a",
		"dec a",
		"dec a",
	}, false))
	TestEqual(t, 12748, assembunny(Lines(23), true))
}
