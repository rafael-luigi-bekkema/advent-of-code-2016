package day01

import (
	"testing"

	. "github.com/rafael-luigi-bekkema/advent-of-code-2016/util"
)

func TestDay1a(t *testing.T) {
	TestEqual(t, 5, day1a("R2, L3"))
	TestEqual(t, 2, day1a("R2, R2, R2"))
	TestEqual(t, 12, day1a("R5, L5, R5, R3"))
	TestEqual(t, 288, day1a(Input(1)))
}
func TestDay1b(t *testing.T) {
	TestEqual(t, 4, day1b("R8, R4, R4, R8"))
	TestEqual(t, 111, day1b(Input(1)))
}
