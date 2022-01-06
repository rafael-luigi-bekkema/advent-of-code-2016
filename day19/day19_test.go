package day19

import (
	"testing"

	"github.com/rafael-luigi-bekkema/advent-of-code-2016/util"
)

func TestDay19(t *testing.T) {
	util.TestEqual(t, 3, day19a(5))
	util.TestEqual(t, 7, day19a(11))
	util.TestEqual(t, 37, day19a(50))
	util.TestEqual(t, 1830117, day19a(3012210))

	util.TestEqual(t, 6, day19b3(15))
	util.TestEqual(t, 19, day19b3(100))
	util.TestEqual(t, 813, day19b3(1500))
	util.TestEqual(t, 1417887, day19b3(3012210))
}
