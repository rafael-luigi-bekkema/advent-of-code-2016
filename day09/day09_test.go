package day09

import (
	"testing"

	. "github.com/rafael-luigi-bekkema/advent-of-code-2016/util"
)

func TestDay9(t *testing.T) {
	TestEqual(t, 6, day9decompressA("ADVENT"))
	TestEqual(t, 7, day9decompressA("A(1x5)BC"))
	TestEqual(t, 9, day9decompressA("(3x3)XYZ"))
	TestEqual(t, 11, day9decompressA("A(2x2)BCD(2x2)EFG"))
	TestEqual(t, 6, day9decompressA("(6x1)(1x3)A"))
	TestEqual(t, 18, day9decompressA("X(8x2)(3x3)ABCY"))
	file := Input(9)
	TestEqual(t, 115118, day9decompressA(file))

	TestEqual(t, 9, day9decompressB("(3x3)XYZ"))
	TestEqual(t, 20, day9decompressB("X(8x2)(3x3)ABCY"))
	TestEqual(t, 241920, day9decompressB("(27x12)(20x12)(13x14)(7x10)(1x12)A"))
	TestEqual(t, 445, day9decompressB("(25x3)(3x3)ABC(2x3)XY(5x2)PQRSTX(18x9)(3x2)TWO(5x7)SEVEN"))

	TestEqual(t, 11107527530, day9decompressB(file))
}
