package day11

import (
	"bytes"
	"fmt"
	"sort"

	"github.com/rafael-luigi-bekkema/advent-of-code-2016/util"
)

func day11a() int {
	lift := 0
	levels := [4][]string{
		{"SG", "SM", "PG", "PM"},
		{"TG", "RG", "RM", "CG", "CM"},
		{"TM"},
		{},
	}

	validate := func() {
		for _, level := range levels {
			var hasGen bool
			match := map[byte]int{}
			for _, item := range level {
				if item[1] == 'G' {
					hasGen = true
					match[item[0]]--
					continue
				}
				match[item[0]]++
			}
			if hasGen {
				for c, m := range match {
					if m > 0 {
						panic(fmt.Sprintf("unprotected M: %cM", c))
					}
				}
			}
		}
	}

	move := func(from, to int, items ...string) {
		if len(items) == 0 || len(items) > 2 {
			panic("not enough items")
		}
		if from != lift {
			panic("wrong level")
		}
		for _, item := range items {
			var found bool
			for i := len(levels[from]) - 1; i >= 0; i-- {
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
		lift = to
	}

	render := func() {
		var all []string
		for _, l := range levels {
			all = append(all, l...)
		}
		sort.Strings(all)

		fmt.Println()
		for i := 3; i >= 0; i-- {
			fmt.Printf("  %d |", i+1)
			if i == lift {
				fmt.Printf(" E |")
			} else {
				fmt.Printf("   |")
			}
		all:
			for _, a := range all {
				for _, sa := range levels[i] {
					if a == sa {
						fmt.Printf(" %s |", a)
						continue all
					}
				}
				fmt.Printf("    |")
			}
			fmt.Println()
		}
	}

	type Move struct {
		from, to int
		items    []string
	}
	moves := []Move{
		{0, 1, []string{"SG", "PG"}},
		{1, 2, []string{"SG", "TG"}},
		{2, 1, []string{"SG"}},
		{1, 2, []string{"SG", "PG"}},
		{2, 1, []string{"SG"}},
		{1, 2, []string{"RG", "RM"}},
		{2, 1, []string{"PG"}},
		{1, 2, []string{"PG", "SG"}},
		{2, 1, []string{"PG"}},
		{1, 2, []string{"CG", "CM"}},
		{2, 1, []string{"SG"}},
		{1, 2, []string{"SG", "PG"}},
		{2, 3, []string{"CG", "CM"}},
		{3, 2, []string{"CG"}},
		{2, 3, []string{"CG", "PG"}},
		{3, 2, []string{"PG"}},
		{2, 3, []string{"PG", "SG"}},
		{3, 2, []string{"SG"}},
		{2, 3, []string{"TG", "TM"}},
		{3, 2, []string{"PG"}},
		{2, 3, []string{"PG", "SG"}},
		{3, 2, []string{"SG"}},
		{2, 3, []string{"RG", "RM"}},
		{3, 2, []string{"PG"}},
		{2, 3, []string{"PG", "SG"}},
		{3, 2, []string{"RM"}},
		{2, 1, []string{"RM"}},
		{1, 0, []string{"RM"}},
		{0, 1, []string{"RM", "PM"}},
		{1, 2, []string{"RM", "PM"}},
		{2, 3, []string{"RM", "PM"}},
		{3, 2, []string{"RM"}},
		{2, 1, []string{"RM"}},
		{1, 0, []string{"RM"}},
		{0, 1, []string{"RM", "SM"}},
		{1, 2, []string{"RM", "SM"}},
		{2, 3, []string{"RM", "SM"}},
	}

	render()
	for i, m := range moves {
		fmt.Printf("\nmove %d: %v to %d\n", i+1, m.items, m.to)
		move(m.from, m.to, m.items...)
		render()
		validate()
	}
	return len(moves)
}

func day11b(b bool) int {
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
		if move.item1 == "" {
			panic("not enough items")
		}
		for _, item := range []string{move.item1, move.item2} {
			if item == "" {
				continue
			}
			from, to := move.from, move.to
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
		}
		if up < 0 {
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

	// var minfound int
	// moves := []Move{
	// 	{0, 1, "SG", "PG"},
	// 	{1, 2, "SG", "TG"},
	// 	{2, 1, "SG", ""},
	// 	{1, 2, "SG", "PG"},
	// 	{2, 1, "SG", ""},
	// 	{1, 2, "RG", "RM"},
	// 	{2, 1, "PG", ""},
	// 	{1, 2, "PG", "SG"},
	// 	{2, 1, "PG", ""},
	// 	{1, 2, "CG", "CM"},
	// 	{2, 1, "SG", ""},
	// 	{1, 2, "SG", "PG"},
	// 	{2, 3, "CG", "CM"},
	// 	{3, 2, "CG", ""},
	// 	{2, 3, "CG", "PG"},
	// 	{3, 2, "PG", ""},
	// 	{2, 3, "PG", "SG"},
	// 	{3, 2, "SG", ""},
	// 	{2, 3, "TG", "TM"},
	// 	{3, 2, "PG", ""},
	// 	{2, 3, "PG", "SG"},
	// 	{3, 2, "SG", ""},
	// 	{2, 3, "RG", "RM"},
	// 	{3, 2, "PG", ""},
	// 	{2, 3, "PG", "SG"},
	// 	{3, 2, "RM", ""},
	// 	{2, 1, "RM", ""},
	// 	{1, 0, "RM", ""},
	// 	{0, 1, "RM", "PM"},
	// 	{1, 2, "RM", "PM"},
	// 	{2, 3, "RM", "PM"},
	// 	{3, 2, "RM", ""},
	// 	{2, 1, "RM", ""},
	// 	{1, 0, "RM", ""},
	// 	{0, 1, "RM", "SM"},
	// 	{1, 2, "RM", "SM"},
	// 	{2, 3, "RM", "SM"},
	// }
	// var count int
	// for _, mv := range moves {
	// 	items := []string{mv.item1}
	// 	if mv.item2 != "" {
	// 		items = append(items, mv.item2)
	// 	}
	// 	levels = move(levels, mv.from, mv.to, items...)
	// 	if !validate(levels) {
	// 		panic("invalid")
	// 	}
	// 	count++
	// }
	// return count

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
	res := day11b(true)
	fmt.Println("day 11a:", res)
}
