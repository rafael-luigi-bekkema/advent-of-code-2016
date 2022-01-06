package day01

import (
	"strconv"
	"strings"

	"github.com/rafael-luigi-bekkema/advent-of-code-2016/set"
	. "github.com/rafael-luigi-bekkema/advent-of-code-2016/util"
)

func day1a(input string) int {
	var x, y int
	dir := 0

	for _, part := range strings.Split(input, ", ") {
		switch part[0] {
		case 'R':
			dir = (dir + 1) % 4
		case 'L':
			dir = (dir + 3) % 4
		}
		n, _ := strconv.Atoi(part[1:])
		switch dir {
		case 0:
			y += n
		case 1:
			x += n
		case 2:
			y -= n
		case 3:
			x -= n
		}
	}
	blocks := Abs(x) + Abs(y)
	return blocks
}

func day1b(input string) int {
	var x, y int
	dir := 0
	hist := set.New([2]int{0, 0})

outer:
	for _, part := range strings.Split(input, ", ") {
		switch part[0] {
		case 'R':
			dir = (dir + 1) % 4
		case 'L':
			dir = (dir + 3) % 4
		}
		var dx, dy int
		n := Atoi(part[1:])
		switch dir {
		case 0:
			dy = 1
		case 1:
			dx = 1
		case 2:
			dy = -1
		case 3:
			dx = -1
		}
		for i := 1; i <= n; i++ {
			x += dx
			y += dy
			if hist.Has([2]int{x, y}) {
				break outer
			}
			hist.Add([2]int{x, y})
		}
	}
	blocks := Abs(x) + Abs(y)
	return blocks

}
