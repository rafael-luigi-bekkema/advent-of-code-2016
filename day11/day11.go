package day11

import (
	"fmt"
	"sort"
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

func day11b() int {
	step := 0
	_ = step
	levels := [4][]string{
		{"SG", "SM", "PG", "PM", "EG", "EM", "DG", "DM"},
		{"TG", "RG", "RM", "CG", "CM"},
		{"TM"},
		{},
	}

	copyLevels := func(levels [4][]string) (result [4][]string) {
		for i, l := range levels {
			result[i] = make([]string, len(l))
			copy(result[i], l)
		}
		return
	}

	validate := func(levels [4][]string) bool {
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

	var all []string
	for _, l := range levels {
		all = append(all, l...)
	}
	sort.Strings(all)

	render := func(levels [4][]string, lift int) {
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

	move := func(levels [4][]string, from, to int, items ...string) [4][]string {
		// fmt.Printf("\nmove %d: %v to %d\n", step+1, items, to)
		if len(items) == 0 || len(items) > 2 {
			panic("not enough items")
		}
		// if from != lift {
		// 	panic("wrong level")
		// }
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
		// lift = to
		// step++
		// render()
		// validate(levels)
		return levels
	}

	type Move struct {
		from, to     int
		item1, item2 string
	}

	// up = always 2
	// down = always 1

	// h
	// 	{1, 2, []string{"TG"}},
	// 	{2, 3, []string{"TG", "TM"}},
	// 	{3, 2, []string{"TG"}},
	// 	{2, 1, []string{"TG"}},
	// 	{1, 2, []string{"RG", "RM"}},
	// 	{2, 1, []string{"RG"}},
	// 	{1, 0, []string{"RG"}},
	// 	{0, 1, []string{"SG", "SM"}},
	// 	{1, 0, []string{"TG"}},
	// 	{0, 1, []string{"TG", "RG"}},
	// 	{1, 0, []string{"RG"}},
	// 	{0, 1, []string{"EG", "EM"}},
	// 	{1, 0, []string{"TG"}},
	// 	{0, 1, []string{"TG", "RG"}},
	// 	{1, 0, []string{"TG"}},
	// 	{0, 1, []string{"TG", "PG"}},
	// 	{1, 2, []string{"RG", "TG"}},
	// 	{2, 3, []string{"RG", "TG"}},
	// 	{3, 2, []string{"RG"}},
	// 	{2, 1, []string{"RG"}},
	// 	{1, 2, []string{"PG", "RG"}},
	// 	{2, 3, []string{"PG", "RG"}},
	// 	{3, 2, []string{"RG"}},
	// 	{2, 3, []string{"RM", "RG"}},
	// 	{3, 2, []string{"PG"}},
	// 	{2, 1, []string{"PG"}},
	// 	{1, 2, []string{"EG", "EM"}},
	// 	{2, 1, []string{"EG"}},
	// 	{1, 2, []string{"EG", "PG"}},
	// 	{2, 3, []string{"EG", "PG"}},
	// 	{3, 2, []string{"EG"}},
	// 	{2, 3, []string{"EG", "EM"}},
	// 	{3, 2, []string{"PG"}},
	// 	{2, 1, []string{"PG"}},
	// 	{1, 2, []string{"SG", "SM"}},
	// 	{2, 1, []string{"SG"}},
	// 	{1, 2, []string{"SG", "PG"}},
	// 	{2, 3, []string{"SG", "PG"}},
	// 	{3, 2, []string{"SG"}},
	// 	{2, 3, []string{"SG", "SM"}},
	// 	{3, 2, []string{"PG"}},
	// 	{2, 1, []string{"PG"}},
	// 	{1, 2, []string{"CG", "CM"}},
	// 	{2, 1, []string{"CG"}},
	// 	{1, 2, []string{"CG", "PG"}},
	// 	{2, 3, []string{"CG", "PG"}},
	// 	{3, 2, []string{"CG"}},
	// 	{2, 3, []string{"CG", "CM"}},

	// 	{3, 2, []string{"PG"}},
	// 	{2, 1, []string{"PG"}},
	// 	{1, 0, []string{"PG"}},
	// 	{0, 1, []string{"PG", "PM"}},
	// 	{1, 2, []string{"PG", "PM"}},
	// 	{2, 1, []string{"PG"}},
	// 	{1, 2, []string{"PG", "DG"}},
	// 	{2, 3, []string{"PG", "DG"}},
	// 	{3, 2, []string{"PG"}},
	// 	{2, 3, []string{"PG", "PM"}},
	// 	{3, 2, []string{"DG"}},
	// 	{2, 1, []string{"DG"}},
	// 	{1, 2, []string{"DG", "DM"}},
	// 	{2, 3, []string{"DG", "DM"}},
	// }

	validMoves := func(levels [4][]string, lift, up int) (result [][]string) {
		if up > 0 {
			for i, s := range levels[lift] {
				for _, s2 := range levels[lift][i+1:] {
					if s[1] != s2[1] && s[0] != s2[0] {
						continue
					}
					levels := move(copyLevels(levels), lift, lift+up, s, s2)
					if validate(levels) {
						result = append(result, []string{s, s2})
					}
				}
			}
		}
		for _, s := range levels[lift] {
			levels := move(copyLevels(levels), lift, lift+up, s)
			if validate(levels) {
				result = append(result, []string{s})
			}
		}
		return
	}

	var doit func(levels [4][]string, lift int, hist []Move, depth int)
	doit = func(levels [4][]string, lift int, hist []Move, depth int) {
		render(levels, lift)
		// check end condition
		if len(levels[3]) == len(all) {
			fmt.Println("found it!")
			return
		}

		for _, up := range []int{1, -1} {
			if up > 0 && lift == 3 || up < 0 && lift == 0 {
				continue
			}
			moves := validMoves(levels, lift, up)
		moves:
			for _, m := range moves {
				h := Move{lift, lift + up, m[0], ""}
				if len(m) >= 2 {
					h.item2 = m[1]
				}

				for _, hh := range hist {
					if h == hh {
						continue moves
					}
				}
				// fmt.Println(h, depth)
				nlevels := copyLevels(levels)
				nlevels = move(nlevels, lift, lift+up, m...)
				doit(nlevels, lift+up, append(hist, h), depth+1)
			}
		}
	}

	_ = render
	doit(levels, 0, nil, 0)

	return 0
}
