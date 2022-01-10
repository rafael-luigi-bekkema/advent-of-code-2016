package day11

import (
	"testing"

	. "github.com/rafael-luigi-bekkema/advent-of-code-2016/util"
)

func TestDay11(t *testing.T) {
	TestEqual(t, 37, day11b(false))
	TestEqual(t, 61, day11b(true))
}
