package main

import (
	"crypto/md5"
	"fmt"
	"log"
	"math"
	"sort"
	"strconv"
	"strings"

	"github.com/rafael-luigi-bekkema/advent-of-code-2016/set"
)

func day14b(salt string, nth int) int {
	return day14(salt, nth, true)
}

func day14a(salt string, nth int) int {
	return day14(salt, nth, false)
}

func day14(salt string, nth int, b bool) int {
	stretch := func(hash []byte) []byte {
		for i := 0; i < 2016; i++ {
			hash = []byte(fmt.Sprintf("%x", md5.Sum(hash)))
		}
		return hash
	}

	idx := 0
	type match struct {
		char byte
		hash []byte
		idx  int
	}
	var matches []match
	var keys []match
	for {
		hash := []byte(fmt.Sprintf("%x", md5.Sum([]byte(fmt.Sprintf("%s%d", salt, idx)))))
		if b {
			hash = stretch(hash)
		}

		for i, c := range hash[:len(hash)-4] {
			if !(c == hash[i+1] && c == hash[i+2] && c == hash[i+3] && c == hash[i+4]) {
				continue
			}
		matcher:
			for mi := len(matches) - 1; mi >= 0; mi-- {
				match := &matches[mi]
				if match.idx+1000 < idx {
					matches[mi], matches[len(matches)-1] = matches[len(matches)-1], matches[mi]
					matches = matches[:len(matches)-1]
					continue
				}
				if match.char == c {
					keys = append(keys, *match)
					matches[mi], matches[len(matches)-1] = matches[len(matches)-1], matches[mi]
					matches = matches[:len(matches)-1]
					sort.Slice(keys, func(i, j int) bool {
						return keys[i].idx < keys[j].idx
					})
					if len(keys) >= nth {
						maxIdx := keys[nth-1].idx
						for _, match := range matches {
							if match.idx < maxIdx {
								continue matcher
							}
						}
						return keys[nth-1].idx
					}
				}
			}
			break
		}

		for i, c := range hash[:len(hash)-2] {
			if c == hash[i+1] && c == hash[i+2] {
				matches = append(matches, match{c, hash, idx})
				break
			}
		}

		idx++
		if idx > 2372900 {
			break
		}
	}
	return 0
}

type Coord struct {
	x, y int
}

var inf = math.Inf(0)

func dijkstra(nodes map[Coord]struct{}, edges map[Coord]map[Coord]float64, source Coord) (dist map[Coord]float64, prev map[Coord]Coord) {
	Q := make([]Coord, 0, len(nodes))
	dist = map[Coord]float64{}
	prev = map[Coord]Coord{}
	for node := range nodes {
		dist[node] = inf
		Q = append(Q, node)
	}
	dist[source] = 0

	for len(Q) > 0 {
		var minI int
		for i, u := range Q {
			if i == 0 || dist[u] < dist[Q[minI]] {
				minI = i
			}
		}
		u := Q[minI]
		Q[minI], Q[len(Q)-1] = Q[len(Q)-1], Q[minI]
		Q = Q[:len(Q)-1]

		for v, l := range edges[u] {
			if !In(v, Q...) {
				continue
			}
			alt := dist[u] + l
			if alt < dist[v] {
				dist[v] = alt
				prev[v] = u
			}
		}
	}
	return dist, prev
}

func day13a(favNum int, target Coord) (int, int) {
	isOpen := func(c Coord) bool {
		x, y := c.x, c.y
		val := x*x + 3*x + 2*x*y + y + y*y + favNum
		count := 0
		for _, r := range fmt.Sprintf("%b", val) {
			if r == '1' {
				count++
			}
		}
		return count%2 == 0
	}

	nodes := map[Coord]struct{}{}
	edges := map[Coord]map[Coord]float64{}
	for y := 0; y <= target.y+10; y++ {
		for x := 0; x <= target.x+10; x++ {
			coord := Coord{x, y}
			nodes[coord] = struct{}{}
			for _, offset := range [][2]int{{0, -1}, {-1, 0}, {1, 0}, {0, 1}} {
				rx := x + offset[0]
				ry := y + offset[1]
				if rx < 0 || ry < 0 {
					continue
				}
				if rc := (Coord{rx, ry}); isOpen(rc) {
					if edges[coord] == nil {
						edges[coord] = map[Coord]float64{}
					}
					edges[coord][rc] = 1
				}
			}
		}
	}

	dist, _ := dijkstra(nodes, edges, Coord{1, 1})
	var count int
	for _, distance := range dist {
		if distance <= 50 {
			count++
		}
	}
	return int(dist[target]), count
}

func day12b(input []string) int {
	return day12(input, true)
}

func day12a(input []string) int {
	return day12(input, false)
}

func day12(input []string, b bool) int {
	type CPU struct {
		regs    [4]int
		counter int
	}

	var cpu CPU
	if b {
		cpu.regs[2] = 1
	}

	set := func(s string, v int) {
		cpu.regs[s[0]-'a'] = v
	}

	get := func(s string) int {
		if '0' <= s[0] && s[0] <= '9' || s[0] == '-' {
			return Atoi(s)
		}
		return cpu.regs[s[0]-'a']
	}

	run := func(ins string) {
		cmd, rest, _ := strings.Cut(ins, " ")
		switch cmd {
		case "cpy":
			from, to, _ := strings.Cut(rest, " ")
			set(to, get(from))
		case "inc":
			set(rest, get(rest)+1)
		case "dec":
			set(rest, get(rest)-1)
		case "jnz":
			x, jmp, _ := strings.Cut(rest, " ")
			if get(x) != 0 {
				cpu.counter += get(jmp) - 1
			}
		default:
			panic("unknown cmd")
		}
	}

	for cpu.counter = 0; cpu.counter < len(input); cpu.counter++ {
		run(input[cpu.counter])
	}

	return cpu.regs[0] // register 'a'
}

/*
F4 .  .  .  .  .
F3 .  .  .  LG .
F2 .  HG .  .  .
F1 E  .  HM .  LM
*/

func day11a(floors [4][]string) int {
	safe := func(items []string) bool {
		hasgen := false
		mols := map[string]int{}
		for _, item := range items {
			if item[1] == 'M' {
				mols[item]++
			} else {
				mols[string([]byte{item[0], 'M'})]--
				hasgen = true
			}
		}
		if !hasgen {
			return true
		}
		for _, c := range mols {
			if c > 0 {
				return false
			}
		}
		return true
	}

	/*
		F4 .  .  .  .  .  .  .  .  .  .  .  .
		F3 .  .  .  .  .  .  .  TM .  .  .  .
		F2 .  .  .  .  .  .  TG . RG RM CG CM
		F1 E  SG  SM PG  PM  .  .  .  .  .  .
	*/
	// All Ms op top floor: cannot bring generators into top floor without destroying Ms

	copyFloors := func(floors [4][]string) [4][]string {
		var f [4][]string
		for i, floor := range floors {
			f[i] = make([]string, len(floor))
			copy(f[i], floor)
		}
		return f
	}

	var minMoves int
	var step func(lift int, floors [4][]string, moves int)
	step = func(lift int, floors [4][]string, moves int) {
		if minMoves != 0 && moves >= minMoves {
			return
		}
		for _, up := range []int{1, -1} {
			if up == -1 && lift == 0 || up == 1 && lift == 3 {
				continue
			}
			for i, item := range floors[lift] {
				for _, item2 := range append(floors[lift][i+1:], "") {
					items := []string{item}
					if item2 != "" { // use this to try single item in lift
						items = append(items, item2)
					}
					if lift == 3 && item2 != "" {
						continue
					}
					if up < 0 && item2 != "" {
						continue
					}
					for len(floors[lift+up]) == 0 {
						if up > 0 && lift+up < 3 {
							up++
							continue
						}
						if up < 0 && lift+up > 0 {
							up--
							continue
						}
						break
					}
					if lift+up == 0 && len(floors[0]) == 0 {
						continue
					}
					if lift+up == 3 && len(items) == 1 {
						continue
					}
					if !safe(append(items, floors[lift+up]...)) || !safe(Remove(floors[lift], items...)) {
						continue
					}

					newFloors := copyFloors(floors)
					newFloors[lift] = Remove(newFloors[lift], items...)
					newFloors[lift+up] = append(newFloors[lift+up], items...)
					fmt.Println("try", items, "to", lift+up)
					fmt.Println(newFloors)

					if len(newFloors[0])+len(newFloors[1])+len(newFloors[2]) == 0 {
						nmoves := moves + Abs(up)
						if minMoves == 0 || nmoves < minMoves {
							minMoves = nmoves
							fmt.Println("have minmoves", minMoves)
						}
					} else {
						step(lift+up, newFloors, moves+Abs(up))
					}
				}
			}
		}
	}
	step(0, floors, 0)

	return minMoves
}

func day10(input []string, checklow, checkhigh int, b bool) int {
	type Ins struct {
		frombot, tolow, tohigh int
		towhatlow, towhathigh  string
	}
	var instructions []Ins

	bots := map[int][]int{}
	for i := len(input) - 1; i >= 0; i-- {
		if input[i][:5] != "value" {
			instructions = append(instructions, Ins{})
			ins := &instructions[len(instructions)-1]
			Must(fmt.Sscanf(input[i], "bot %d gives low to %s %d and high to %s %d",
				&ins.frombot, &ins.towhatlow, &ins.tolow, &ins.towhathigh, &ins.tohigh))
			continue
		}
		var val, botnum int
		fmt.Sscanf(input[i], "value %d goes to bot %d", &val, &botnum)
		bots[botnum] = append(bots[botnum], val)

		input[i], input[len(input)-1] = input[len(input)-1], input[i]
		input = input[:len(input)-1]
	}

	outputs := map[int]int{}
	for i := len(instructions) - 1; i >= 0; i-- {
		ins := &instructions[i]

		if len(bots[ins.frombot]) < 2 {
			continue
		}

		low, high := MinMax(bots[ins.frombot])
		if !b && low == checklow && high == checkhigh {
			return ins.frombot
		}

		bots[ins.frombot] = nil
		if ins.towhatlow == "bot" {
			bots[ins.tolow] = append(bots[ins.tolow], low)
		} else {
			outputs[ins.tolow] = low
		}
		if ins.towhathigh == "bot" {
			bots[ins.tohigh] = append(bots[ins.tohigh], high)
		} else {
			outputs[ins.tohigh] = high
		}

		instructions[i], instructions[len(instructions)-1] = instructions[len(instructions)-1], instructions[i]
		instructions = instructions[:len(instructions)-1]

		i = len(instructions)
	}
	return outputs[0] * outputs[1] * outputs[2]
}

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
