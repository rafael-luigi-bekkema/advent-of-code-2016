package day08

import (
	"fmt"
	"strings"
)

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
