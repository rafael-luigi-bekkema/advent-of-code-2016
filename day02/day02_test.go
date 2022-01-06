package day02

import (
	"testing"

	. "github.com/rafael-luigi-bekkema/advent-of-code-2016/util"
)

func TestDay2b(t *testing.T) {
	TestEqual(t, 0x5DB3, day2b([]string{"ULL", "RRDDD", "LURDL", "UUUUD"}))
	TestEqual(t, 0x3CC43, day2b(Lines(2)))
}

func TestDay2a(t *testing.T) {
	TestEqual(t, 1985, day2a([]string{"ULL", "RRDDD", "LURDL", "UUUUD"}))
	TestEqual(t, 19636, day2a(Lines(2)))
}
