package day13

import (
	"testing"

	. "github.com/rafael-luigi-bekkema/advent-of-code-2016/util"
)

func TestDay13(t *testing.T) {
	testa, _ := day13a(10, Coord{7, 4})
	TestEqual(t, 11, testa)
	a, b := day13a(1358, Coord{31, 39})
	TestEqual(t, 96, a)
	TestEqual(t, 141, b)
}
