package day05

import (
	"testing"

	. "github.com/rafael-luigi-bekkema/advent-of-code-2016/util"
)

func TestDay5(t *testing.T) {
	TestEqual(t, "18f47a30", day5a("abc"))
	TestEqual(t, "2414bc77", day5a("wtnhxymk"))

	TestEqual(t, "05ace8e3", day5b("abc"))
	TestEqual(t, "437e60fc", day5b("wtnhxymk"))
}
