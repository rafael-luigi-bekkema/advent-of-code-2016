package day17

import (
	"fmt"
	"log"

	"github.com/rafael-luigi-bekkema/advent-of-code-2016/util"
)

func isOpen(hash string, dir byte) bool {
	var idx int
	switch dir {
	case 'U':
		idx = 0
	case 'D':
		idx = 1
	case 'L':
		idx = 2
	case 'R':
		idx = 3
	default:
		log.Panicf("unknown direction: %c", dir)
	}
	return util.In(hash[idx], 'b', 'c', 'd', 'e', 'f')
}

func day17aRecur(passcode string, path []byte, x, y uint8, minpath string) string {
	if x == 3 && y == 3 {
		// reached destination
		return string(path)
	}
	if minpath != "" && len(path) >= len(minpath) {
		return ""
	}

	hash := util.MD5(fmt.Sprintf("%s%s", passcode, path))
	for _, dir := range []byte{'U', 'D', 'L', 'R'} {
		if dir == 'L' && x == 0 || dir == 'R' && x == 3 || dir == 'U' && y == 0 || dir == 'D' && y == 3 {
			continue
		}
		dx, dy := x, y
		switch dir {
		case 'U':
			dy--
		case 'D':
			dy++
		case 'L':
			dx--
		case 'R':
			dx++
		}
		open := isOpen(hash, dir)
		if open {
			path := day17aRecur(passcode, append(path, dir), dx, dy, minpath)
			if minpath == "" || (path != "" && len(path) < len(minpath)) {
				minpath = path
			}
		}
	}
	return minpath
}

func day17bRecur(passcode string, path []byte, x, y uint8, minpath string) string {
	if x == 3 && y == 3 {
		// reached destination
		return string(path)
	}

	hash := util.MD5(fmt.Sprintf("%s%s", passcode, path))
	for _, dir := range []byte{'U', 'D', 'L', 'R'} {
		if dir == 'L' && x == 0 || dir == 'R' && x == 3 || dir == 'U' && y == 0 || dir == 'D' && y == 3 {
			continue
		}
		dx, dy := x, y
		switch dir {
		case 'U':
			dy--
		case 'D':
			dy++
		case 'L':
			dx--
		case 'R':
			dx++
		}
		open := isOpen(hash, dir)
		if open {
			path := day17bRecur(passcode, append(path, dir), dx, dy, minpath)
			if minpath == "" || (path != "" && len(path) > len(minpath)) {
				minpath = path
			}
		}
	}
	return minpath
}

func day17a(passcode string) string {
	return day17aRecur(passcode, nil, 0, 0, "")
}

func day17b(passcode string) int {
	return len(day17bRecur(passcode, nil, 0, 0, ""))
}
