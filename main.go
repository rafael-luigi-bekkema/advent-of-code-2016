package main

import (
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"

	"github.com/rafael-luigi-bekkema/advent-of-code-2016/set"
)

func day4b(input string, sid int) string {
	chars := []byte(strings.ReplaceAll(input, "-", " "))
	for i := 0; i < sid; i++ {
		for j, c := range chars {
			if c == ' ' {
				continue
			}
			if c == 'z' {
				chars[j] = 'a'
				continue
			}
			chars[j]++
		}
	}
	return string(chars)
}

func day4a(input []string) (int, int) {
	checksum := func(inp string) string {
		cc := map[rune]int{}
		var chars []rune
		for _, c := range inp {
			if c == '-' {
				continue
			}
			if _, ok := cc[c]; !ok {
				chars = append(chars, c)
			}
			cc[c]++
		}
		sort.Slice(chars, func(i, j int) bool {
			if cc[chars[i]] == cc[chars[j]] {
				return chars[i] < chars[j]
			}
			return cc[chars[i]] > cc[chars[j]]
		})
		return string(chars[:5])
	}

	var total int
	var partBid int
	for _, line := range input {
		lidx := strings.LastIndex(line, "-")
		csum := checksum(line[:lidx])

		var sid int
		var sum string
		fmt.Sscanf(line[lidx:], "-%d[%s]", &sid, &sum)
		decoded := day4b(line[:lidx], sid)
		if decoded == "northpole object storage" {
			partBid = sid
		}
		if sum[:len(sum)-1] == csum {
			total += sid
		}
	}
	return total, partBid
}

func day3b(input [][3]int) int {
	var count int
	f := func(s1, s2, s3 int) {
		if s1+s2 > s3 && s2+s3 > s1 && s1+s3 > s2 {
			count++
		}
	}
	for i := 0; i < len(input); i += 3 {
		f(input[i][0], input[i+1][0], input[i+2][0])
		f(input[i][1], input[i+1][1], input[i+2][1])
		f(input[i][2], input[i+1][2], input[i+2][2])
	}
	return count
}

func day3a(input [][3]int) int {
	var count int
	for _, nums := range input {
		s1, s2, s3 := nums[0], nums[1], nums[2]
		if s1+s2 > s3 && s2+s3 > s1 && s1+s3 > s2 {
			count++
		}
	}
	return count
}

func day3input() (input [][3]int) {
	s, close := InputScanner(3)
	defer close()
	for s.Scan() {
		input = append(input, [3]int{})
		t := &input[len(input)-1]
		fmt.Sscanf(s.Text(), "%d %d %d", &t[0], &t[1], &t[2])
	}
	return
}

func day2a(input []string) int {
	x := 1
	y := 1
	total := 0
	for _, line := range input {
		for _, c := range line {
			dx, dy := x, y
			switch c {
			case 'U':
				dy--
			case 'D':
				dy++
			case 'L':
				dx--
			case 'R':
				dx++
			}
			if dx < 0 || dx > 2 || dy < 0 || dy > 2 {
				continue
			}
			x, y = dx, dy
		}
		total = total*10 + (y*3 + x + 1)
	}
	return total
}

func day2b(input []string) int {
	keypad := map[[2]int]int{
		{2, 0}: 0x1,

		{1, 1}: 0x2,
		{2, 1}: 0x3,
		{3, 1}: 0x4,

		{0, 2}: 0x5,
		{1, 2}: 0x6,
		{2, 2}: 0x7,
		{3, 2}: 0x8,
		{4, 2}: 0x9,

		{1, 3}: 0xA,
		{2, 3}: 0xB,
		{3, 3}: 0xC,

		{2, 4}: 0xD,
	}
	x := 0
	y := 2
	total := 0
	for _, line := range input {
		for _, c := range line {
			dx, dy := x, y
			switch c {
			case 'U':
				dy--
			case 'D':
				dy++
			case 'L':
				dx--
			case 'R':
				dx++
			}
			if _, ok := keypad[[2]int{dx, dy}]; !ok {
				continue
			}
			x, y = dx, dy
		}
		total = total*16 + keypad[[2]int{x, y}]
	}
	return total
}

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

func main() {
	log.Println("Advent of Code 2016")

	log.Println("Not much to see here. Run the tests:\ngo test -v ./...")
}

func init() {
	log.SetFlags(0)
}
