package day15

import "fmt"

func day15b(input []string) int {
	return day15(input, true)
}

func day15a(input []string) int {
	return day15(input, false)
}

func day15(input []string, b bool) int {
	type disc struct {
		nr        int
		positions int
		startPos  int
	}
	// Disc #6 has 19 positions; at time=0, it is at position 17.
	discs := make([]disc, len(input), len(input)+1)
	for i, line := range input {
		d := &discs[i]
		fmt.Sscanf(line, "Disc #%d has %d positions; at time=0, it is at position %d.",
			&d.nr, &d.positions, &d.startPos)
	}
	if b {
		discs = append(discs, disc{
			nr:        len(discs) + 1,
			positions: 11,
			startPos:  0,
		})
	}
	count := 0
outer:
	for {
		for i, disc := range discs {
			t := count + i + 1
			pos := (disc.startPos + t) % disc.positions
			if pos != 0 {
				count++
				continue outer
			}
		}
		return count
	}
}

