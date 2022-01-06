package day03

import (
	"testing"

	. "github.com/rafael-luigi-bekkema/advent-of-code-2016/util"
)

func TestDay3(t *testing.T) {
	TestEqual(t, 2, day3a([][3]int{{5, 6, 4}, {5, 10, 25}, {5, 5, 5}}))
	input := day3input()
	TestEqual(t, 982, day3a(input))
	TestEqual(t, 1826, day3b(input))
}
