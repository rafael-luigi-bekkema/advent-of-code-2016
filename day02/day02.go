package day02

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
