package day07

import (
	"testing"

	. "github.com/rafael-luigi-bekkema/advent-of-code-2016/util"
)



func TestDay7(t *testing.T) {
	TestEqual(t, true, day7abba("abba[mnop]qrst"))
	TestEqual(t, false, day7abba("abcd[bddb]xyyx"))
	testA, _ := day7([]string{"abba[mnop]qrst", "abcd[bddb]xyyx", "aaaa[qwer]tyui", "ioxxoj[asdfgh]zxcvbn"})
	TestEqual(t, 2, testA)
	resultA, resultB := day7(Lines(7))
	TestEqual(t, 118, resultA)

	TestEqual(t, true, day7ssl("aba[bab]xyz"))
	TestEqual(t, false, day7ssl("xyx[xyx]xyx"))
	TestEqual(t, true, day7ssl("aaa[kek]eke"))
	TestEqual(t, true, day7ssl("zazbz[bzb]cdb"))
	TestEqual(t, 260, resultB)
}
