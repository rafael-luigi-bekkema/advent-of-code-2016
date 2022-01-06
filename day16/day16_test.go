package day16

import (
	"testing"

	. "github.com/rafael-luigi-bekkema/advent-of-code-2016/util"
)

func TestDay16(t *testing.T) {
	TestEqual(t, "01100", day16a(20, "10000"))
	TestEqual(t, "01110011101111011", day16a(272, "11110010111001001"))
	TestEqual(t, "11001111011000111", day16a(35651584, "11110010111001001"))
}
