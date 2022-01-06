package day06

import (
	"testing"

	. "github.com/rafael-luigi-bekkema/advent-of-code-2016/util"
)

func TestDay6(t *testing.T) {
	input := []string{"eedadn", "drvtee", "eandsr", "raavrd", "atevrs", "tsrnev", "sdttsa", "rasrtv",
		"nssdts", "ntnada", "svetve", "tesnvt", "vntsnd", "vrdear", "dvrsen", "enarar"}
	file := Lines(6)
	TestEqual(t, "easter", day6a(input))
	TestEqual(t, "tsreykjj", day6a(file))

	TestEqual(t, "advent", day6b(input))
	TestEqual(t, "hnfbujie", day6b(file))
}

