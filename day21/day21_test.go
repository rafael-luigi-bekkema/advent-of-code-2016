package day21

import (
	"testing"

	"github.com/rafael-luigi-bekkema/advent-of-code-2016/util"
)

func TestDay21(t *testing.T) {
	util.TestEqual(t, "decab", day21a("abcde", []string{
		"swap position 4 with position 0",
		"swap letter d with letter b",
		"reverse positions 0 through 4",
		"rotate left 1 step",
		"move position 1 to position 4",
		"move position 3 to position 0",
		"rotate based on position of letter b",
		"rotate based on position of letter d",
	}))
	util.TestEqual(t, "abcde", day21b("decab", []string{
		"swap position 4 with position 0",
		"swap letter d with letter b",
		"reverse positions 0 through 4",
		"rotate left 1 step",
		"move position 1 to position 4",
		"move position 3 to position 0",
		"rotate based on position of letter b",
		"rotate based on position of letter d",
	}))
	util.TestEqual(t, "acbde", day21a("abcde", []string{"swap position 2 with position 1"}))
	util.TestEqual(t, "abcde", day21b("acbde", []string{"swap position 2 with position 1"}))

	util.TestEqual(t, "acbde", day21a("abcde", []string{"swap letter c with letter b"}))
	util.TestEqual(t, "abcde", day21b("acbde", []string{"swap letter c with letter b"}))

	util.TestEqual(t, "adcbe", day21a("abcde", []string{"reverse positions 1 throught 3"}))
	util.TestEqual(t, "abcde", day21a("adcbe", []string{"reverse positions 1 throught 3"}))

	util.TestEqual(t, "cdeab", day21a("abcde", []string{"rotate left 2 steps"}))
	util.TestEqual(t, "abcde", day21b("cdeab", []string{"rotate left 2 steps"}))

	util.TestEqual(t, "cdeab", day21a("abcde", []string{"rotate right 3 steps"}))
	util.TestEqual(t, "abcde", day21b("cdeab", []string{"rotate right 3 steps"}))

	util.TestEqual(t, "cdeab", day21a("abcde", []string{"rotate based on position of letter c"}))
	util.TestEqual(t, "cdeab", day21a("deabc", []string{"rotate based on position of letter c"}))

	util.TestEqual(t, "eabcd", day21a("abcde", []string{"rotate based on position of letter e"}))
	util.TestEqual(t, "abcde", day21b("eabcd", []string{"rotate based on position of letter e"}))

	util.TestEqual(t, "abdec", day21a("abcde", []string{"move position 2 to position 4"}))
	util.TestEqual(t, "abcde", day21b("abdec", []string{"move position 2 to position 4"}))

	util.TestEqual(t, "dabce", day21a("abcde", []string{"move position 3 to position 0"}))
	util.TestEqual(t, "abcde", day21b("dabce", []string{"move position 3 to position 0"}))

	lines := util.Lines(21)
	util.TestEqual(t, "bgfacdeh", day21a("abcdefgh", lines))
	util.TestEqual(t, "abcdefgh", day21b("bgfacdeh", lines))
	util.TestEqual(t, "bdgheacf", day21b("fbgdceah", lines))
}
