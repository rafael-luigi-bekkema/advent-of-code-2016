package day06


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
