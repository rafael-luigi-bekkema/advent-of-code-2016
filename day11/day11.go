package day11

import (
	"bytes"
	"fmt"
	"sort"

	"github.com/rafael-luigi-bekkema/advent-of-code-2016/util"
)

func day11(b bool) int {
	var levels [4][]string
	if !b {
		levels = [4][]string{
			{"SG", "SM", "PG", "PM"},
			{"TG", "RG", "RM", "CG", "CM"},
			{"TM"},
			{},
		}
	} else {
		levels = [4][]string{
			{"SG", "SM", "PG", "PM", "EG", "EM", "DG", "DM"},
			{"TG", "RG", "RM", "CG", "CM"},
			{"TM"},
			{},
		}
	}

	var all []string
	for _, l := range levels {
		all = append(all, l...)
	}
	sort.Strings(all)

	copyLevels := func(levels [4][]string) (result [4][]string) {
		for i, l := range levels {
			result[i] = make([]string, len(l))
			copy(result[i], l)
		}
		return
	}

	validate := func(levels [4][]string) bool {
		if len(levels[3]) == len(all) {
			return true
		}
		for n, level := range levels {
			match := map[byte]int{}
			var mcount, gcount int
			for _, item := range level {
				if item[1] == 'G' {
					match[item[0]]--
					gcount++
					continue
				}
				mcount++
				match[item[0]]++
			}
			if n == 3 && (mcount-gcount) > 2 {
				return false
			}
			// if n == 3 && mcount > 0 && mcount == gcount {
			// 	return false
			// }
			if gcount > 0 {
				for _, m := range match {
					if m > 0 {
						return false
						// panic(fmt.Sprintf("unprotected M: %cM", c))
					}
				}
			}
		}
		return true
	}

	render := func(levels [4][]string, lift, up int) string {
		bb := &bytes.Buffer{}
		fmt.Fprintln(bb)
		for i := 3; i >= 0; i-- {
			fmt.Fprintf(bb, "  %d |", i+1)
			if i == lift {
				ups := " "
				if up < 0 {
					ups = "v"
				} else if up > 0 {
					ups = "^"
				}
				fmt.Fprintf(bb, "%sE |", ups)
			} else {
				fmt.Fprintf(bb, "   |")
			}
		all:
			for _, a := range all {
				for _, sa := range levels[i] {
					if a == sa {
						fmt.Fprintf(bb, " %s |", a)
						continue all
					}
				}
				fmt.Fprintf(bb, "    |")
			}
			fmt.Fprintln(bb)
		}
		return bb.String()
	}

	type Move struct {
		from, to     int
		item1, item2 string
	}

	move := func(levels [4][]string, move Move) [4][]string {
		from, to := move.from, move.to
		for _, item := range []string{move.item1, move.item2} {
			if item == "" {
				continue
			}
			var found bool
			for i := len(levels[move.from]) - 1; i >= 0; i-- {
				if levels[from][i] == item {
					levels[from][i], levels[from][len(levels[from])-1] = levels[from][len(levels[from])-1], levels[from][i]
					levels[from] = levels[from][:len(levels[from])-1]
					found = true
					break
				}
			}
			if !found {
				panic("moved item from wrong level")
			}
			levels[to] = append(levels[to], item)
		}
		return levels
	}

	validMoves := func(levels [4][]string, lift, up int) (result []Move) {
		if up > 0 {
			for i, s := range levels[lift] {
				for _, s2 := range levels[lift][i+1:] {
					if s[1] != s2[1] && s[0] != s2[0] {
						continue
					}
					mv := Move{lift, lift + up, s, s2}
					levels := move(copyLevels(levels), mv)
					if validate(levels) {
						result = append(result, mv)
					}
				}
			}
		}
		if up < 0 {
			if lift == 1 && len(levels[0]) == 0 {
				return
			}
			if lift == 2 && len(levels[0]) == 0 && len(levels[1]) == 0 {
				return
			}
			for _, s := range levels[lift] {
				mv := Move{lift, lift + up, s, ""}
				levels := move(copyLevels(levels), mv)
				if validate(levels) {
					result = append(result, mv)
				}
			}
		}
		return
	}

	fmt.Println(render(levels, 0, 0))
	var minfound int
	var doit func(levels [4][]string, lift int, hist []Move, depth int) int
	doit = func(levels [4][]string, lift int, hist []Move, depth int) int {
		// render(levels, lift)
		// check end condition
		if len(levels[3]) == len(all) {
			if depth < minfound || minfound == 0 {
				minfound = depth
				fmt.Println(hist)
				fmt.Println("found it!", depth)
				return depth
			}
			return 0
		}
		if minfound != 0 && depth >= minfound {
			return 0
		}

		var minres int
		for _, up := range []int{1, -1} {
			if up > 0 && lift == 3 || up < 0 && lift == 0 {
				continue
			}
			moves := validMoves(levels, lift, up)
			for _, m := range moves {
				nlevels := move(copyLevels(levels), m)
				if util.In(m, hist...) {
					continue
				}
				res := doit(nlevels, lift+up, append(hist, m), depth+1)
				if res != 0 {
					return res
				}
			}
		}
		return minres
	}

	return doit(levels, 0, nil, 0)
}

func Day11b() {
	res := day11(true)
	fmt.Println("day 11a:", res)
}
