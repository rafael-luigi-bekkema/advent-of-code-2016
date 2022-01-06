package day04

import (
	"testing"

	. "github.com/rafael-luigi-bekkema/advent-of-code-2016/util"
)

func TestDay4a(t *testing.T) {
	result, _ := day4a([]string{
		"aaaaa-bbb-z-y-x-123[abxyz]",
		"a-b-c-d-e-f-g-h-987[abcde]",
		"not-a-real-room-404[oarel]",
		"totally-real-room-200[decoy]"})
	TestEqual(t, 1514, result)
	result, resultID := day4a(Lines(4))
	TestEqual(t, 245102, result)
	TestEqual(t, "very encrypted name", day4b("qzmt-zixmtkozy-ivhz", 343))
	TestEqual(t, 324, resultID)
}
