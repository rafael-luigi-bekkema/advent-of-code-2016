package day10

import (
	"testing"

	. "github.com/rafael-luigi-bekkema/advent-of-code-2016/util"
)



func TestDay10(t *testing.T) {
	TestEqual(t, 2, day10([]string{
		"value 5 goes to bot 2",
		"bot 2 gives low to bot 1 and high to bot 0",
		"value 3 goes to bot 1",
		"bot 1 gives low to output 1 and high to bot 0",
		"bot 0 gives low to output 2 and high to output 0",
		"value 2 goes to bot 2",
	}, 2, 5, false))
	file := Lines(10)
	TestEqual(t, 47, day10(file, 17, 61, false))
	TestEqual(t, 2666, day10(file, 17, 61, true))
}
