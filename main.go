package main

import (
	"crypto/md5"
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"

	"github.com/rafael-luigi-bekkema/advent-of-code-2016/set"
)

func day9decompressB(input string) int {
	return day9decompress(input, true, 0)
}

func day9decompressA(input string) int {
	return day9decompress(input, false, 0)
}

func day9decompress(input string, recursive bool, depth int) int {
	var chunk string
	var ok bool
	var length, n int
	var count int
	for {
		chunk, input, ok = strings.Cut(input, "(")
		count += len(chunk)
		if !ok {
			break
		}
		chunk, input, ok = strings.Cut(input, ")")
		fmt.Sscanf(chunk, "%dx%d", &length, &n)

		chunk = input[:length]
		chunkLen := len(chunk)
		if recursive {
			chunkLen = day9decompress(chunk, true, depth+1)
		}
		count += n * chunkLen
		input = input[length:]
	}
	return count
}

func day8(input []string, w, h int) (int, string) {
	display := make([][]bool, w)
	for y := 0; y < h; y++ {
		display[y] = make([]bool, w)
	}

	render := func() string {
		s := &strings.Builder{}
		for y := 0; y < h; y++ {
			for x := 0; x < w; x++ {
				if display[y][x] {
					fmt.Fprint(s, "#")
				} else {
					fmt.Fprint(s, " ")
				}
			}
			fmt.Fprintln(s)
		}
		return s.String()
	}

	var cmd string
	for _, line := range input {
		cmd, line, _ = strings.Cut(line, " ")
		switch cmd {
		case "rect":
			var wide, tall int
			fmt.Sscanf(line, "%dx%d", &wide, &tall)
			for x := 0; x < wide; x++ {
				for y := 0; y < tall; y++ {
					display[y][x] = true
				}
			}
		case "rotate":
			what, rest, _ := strings.Cut(line, "=")
			var axis, amount int
			fmt.Sscanf(rest, "%d by %d", &axis, &amount)
			for n := 1; n <= amount; n++ {
				if what[:3] == "row" {
					y := axis
					last := display[y][w-1]
					for x := w - 1; x > 0; x-- {
						display[y][x] = display[y][x-1]
					}
					display[y][0] = last
					continue
				}
				x := axis
				last := display[h-1][x]
				for y := h - 1; y > 0; y-- {
					display[y][x] = display[y-1][x]
				}
				display[0][x] = last
			}
		}
	}

	var count int
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			if display[y][x] {
				count++
			}
		}
	}
	return count, render()
}

func day7ssl(input string) bool {
	aba := func(s string) (a []string) {
		for i := 0; i < len(s)-2; i++ {
			if s[i] == s[i+2] && s[i] != s[i+1] {
				a = append(a, s[i:i+3])
			}
		}
		return
	}
	bab := func(hyp, aba string) bool {
		for i := 0; i < len(hyp)-2; i++ {
			if hyp[i] == aba[1] && hyp[i+1] == aba[0] && hyp[i+2] == aba[1] {
				return true
			}
		}
		return false
	}

	var addr, hyp string
	var hypers []string
	var abas []string
	for {
		var ok bool
		addr, input, ok = strings.Cut(input, "[")
		if a := aba(addr); a != nil {
			abas = append(abas, a...)
		}
		if !ok {
			break
		}
		hyp, input, _ = strings.Cut(input, "]")
		hypers = append(hypers, hyp)
	}
	for _, aba := range abas {
		for _, hyp := range hypers {
			if bab(hyp, aba) {
				return true
			}
		}
	}
	return false
}

func day7abba(input string) bool {
	abba := func(s string) bool {
		for i := 0; i < len(s)-3; i++ {
			if s[i] != s[i+1] && s[i] == s[i+3] && s[i+1] == s[i+2] {
				return true
			}
		}
		return false
	}

	var addr, hyp string
	isAbba := false
	for {
		var ok bool
		addr, input, ok = strings.Cut(input, "[")
		if abba(addr) {
			isAbba = true
		}
		if !ok {
			break
		}
		hyp, input, ok = strings.Cut(input, "]")
		if abba(hyp) {
			// fmt.Println("hyper abba", hyp)
			return false
		}
	}
	return isAbba
}

func day7(input []string) (a int, b int) {
	for _, line := range input {
		if day7abba(line) {
			a++
		}
		if day7ssl(line) {
			b++
		}
	}
	return
}

func day6b(input []string) string {
	var msg []byte
	cmap := map[int]map[byte]int{}
	for i, line := range input {
		if i == 0 {
			msg = make([]byte, len(line))
			for j := 0; j < len(line); j++ {
				cmap[j] = make(map[byte]int)
			}
		}
		for j, c := range []byte(line) {
			cmap[j][c]++
		}
	}
	for pos, chars := range cmap {
		var mincount int
		var minchar byte
		for char, count := range chars {
			if mincount == 0 || count < mincount {
				mincount = count
				minchar = char
			}
		}
		msg[pos] = minchar
	}
	return string(msg)
}

func day6a(input []string) string {
	var msg []byte
	cmap := map[int]map[byte]int{}
	var ccount []int
	for i, line := range input {
		if i == 0 {
			msg = make([]byte, len(line))
			ccount = make([]int, len(line))
			for j := 0; j < len(line); j++ {
				cmap[j] = make(map[byte]int)
			}
		}
		for j, c := range []byte(line) {
			cmap[j][c]++
			if ccount[j] == 0 || cmap[j][c] > ccount[j] {
				msg[j] = c
				ccount[j] = cmap[j][c]
			}
		}
	}
	return string(msg)
}

func day5b(doorID string) string {
	i := -1
	var found int
	pass := [8]byte{'_', '_', '_', '_', '_', '_', '_', '_'}
	for {
		i++
		hash := fmt.Sprintf("%x", md5.Sum([]byte(fmt.Sprint(doorID, i))))
		if hash[:5] != "00000" {
			continue
		}
		pos := int(hash[5] - '0')
		if !(0 <= pos && pos <= 7) || pass[pos] != '_' {
			continue
		}

		pass[pos] = hash[6]
		found++

		fmt.Printf("\rfound %s %8d %s", string(pass[:]), i, hash)
		if found == 8 {
			fmt.Print("\n\n")
			break
		}
	}
	return string(pass[:])
}

func day5a(doorID string) string {
	i := -1
	pass := make([]byte, 0, 8)
	for {
		i++
		hash := fmt.Sprintf("%x", md5.Sum([]byte(fmt.Sprint(doorID, i))))
		if hash[:5] != "00000" {
			continue
		}
		pass = append(pass, hash[5])
		if len(pass) == 8 {
			break
		}
	}
	return string(pass)
}

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
