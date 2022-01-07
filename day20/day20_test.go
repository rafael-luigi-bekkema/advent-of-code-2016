package day20

import (
	"testing"

	"github.com/rafael-luigi-bekkema/advent-of-code-2016/util"
)

func TestDay20(t *testing.T) {
	a, b := day20(9, []string{"5-8", "0-2", "4-7"})
	util.TestEqual(t, 3, a, "test a")
	util.TestEqual(t, 2, b, "test b")
	a, b = day20(4294967295, util.Lines(20))
	util.TestEqual(t, 22887907, a, "final a")
	util.TestEqual(t, 109, b, "final b")
}
