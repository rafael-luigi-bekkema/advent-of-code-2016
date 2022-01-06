package day18

const (
	trap = '^'
	safe = '.'
)

func day18a(first string, rows int) int {
	row := []byte(first)

	isTrap := func(i int) bool {
		return i > 0 && row[i-1] == trap && row[i] == trap && (i == len(row)-1 || row[i+1] == safe) ||
			i < len(row)-1 && (i == 0 || row[i-1] == safe) && row[i] == trap && row[i+1] == trap ||
			i > 0 && row[i-1] == trap && row[i] == safe && (i == len(row)-1 || row[i+1] == safe) ||
			i < len(row)-1 && (i == 0 || row[i-1] == safe) && row[i] == safe && row[i+1] == trap
	}

	var count int
	for _, c := range row {
		if c == safe {
			count++
		}
	}
	for i := 1; i < rows; i++ {
		newrow := make([]byte, len(row))
		for i := range []byte(row) {
			if isTrap(i) {
				newrow[i] = trap
				continue
			}
			newrow[i] = safe
			count++
		}
		row = newrow
	}
	return count
}
