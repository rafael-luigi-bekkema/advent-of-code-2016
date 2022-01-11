package day25

import (
	"testing"

	. "github.com/rafael-luigi-bekkema/advent-of-code-2016/util"
)

func TestDay25(t *testing.T) {
	TestEqual(t, 180, day25a(Lines(25)))
}
