package day21

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/rafael-luigi-bekkema/advent-of-code-2016/util"
)

func day21b(pass string, ops []string) string {
	bpass := []byte(pass)
	for opi := len(ops) - 1; opi >= 0; opi-- {
		parts := strings.Split(ops[opi], " ")
		switch {
		case parts[0] == "swap" && parts[1] == "position":
			x, y := util.Atoi(parts[2]), util.Atoi(parts[5])
			bpass[x], bpass[y] = bpass[y], bpass[x]
		case parts[0] == "swap" && parts[1] == "letter":
			x, y := parts[2][0], parts[5][0]
			for i := len(bpass) - 1; i >= 0; i-- {
				if bpass[i] == y {
					bpass[i] = x
					continue
				}
				if bpass[i] == x {
					bpass[i] = y
				}
			}
		case parts[0] == "rotate" && parts[1] == "right":
			for i := 0; i < util.Atoi(parts[2]); i++ {
				first := bpass[0]
				copy(bpass, bpass[1:])
				bpass[len(bpass)-1] = first
			}
		case parts[0] == "rotate" && parts[1] == "left":
			for i := 0; i < util.Atoi(parts[2]); i++ {
				last := bpass[len(bpass)-1]
				copy(bpass[1:], bpass)
				bpass[0] = last
			}
		case parts[0] == "rotate" && parts[1] == "based":
			cur := string(bpass)
			var opts []string
			for i := 0; i <= len(bpass)-1; i++ {
				first := bpass[0]
				copy(bpass, bpass[1:])
				bpass[len(bpass)-1] = first

				res := day21a(string(bpass), []string{ops[opi]})
				if res == cur {
					opts = append(opts, string(bpass))
				}
			}
			if len(opts) > 1 {
				fmt.Println(opts)
			}
			copy(bpass, opts[0])
		case parts[0] == "reverse":
			for i, j := util.Atoi(parts[2]), util.Atoi(parts[4]); i < j; i, j = i+1, j-1 {
				bpass[i], bpass[j] = bpass[j], bpass[i]
			}
		case parts[0] == "move":
			y, x := util.Atoi(parts[2]), util.Atoi(parts[5])
			v := bpass[x]
			if y > x {
				copy(bpass[x:y], bpass[x+1:])
			} else {
				copy(bpass[y+1:x+1], bpass[y:])
			}
			bpass[y] = v
		default:
			panic("unknown op: " + ops[opi])
		}
	}
	return string(bpass)

}

func day21a(pass string, ops []string) string {
	bpass := []byte(pass)
	for _, op := range ops {
		parts := strings.Split(op, " ")
		switch {
		case parts[0] == "swap" && parts[1] == "position":
			x, y := util.Atoi(parts[2]), util.Atoi(parts[5])
			bpass[x], bpass[y] = bpass[y], bpass[x]
		case parts[0] == "swap" && parts[1] == "letter":
			x, y := parts[2][0], parts[5][0]
			for i, c := range bpass {
				if c == x {
					bpass[i] = y
					continue
				}
				if c == y {
					bpass[i] = x
				}
			}
		case parts[0] == "rotate" && parts[1] == "right":
			for i := 0; i < util.Atoi(parts[2]); i++ {
				last := bpass[len(bpass)-1]
				copy(bpass[1:], bpass)
				bpass[0] = last
			}
		case parts[0] == "rotate" && parts[1] == "left":
			for i := 0; i < util.Atoi(parts[2]); i++ {
				first := bpass[0]
				copy(bpass, bpass[1:])
				bpass[len(bpass)-1] = first
			}
		case parts[0] == "rotate" && parts[1] == "based":
			c := parts[6][0]
			idx := bytes.Index(bpass, []byte{c})
			n := 1 + idx
			if idx >= 4 {
				n++
			}
			for i := 0; i < n; i++ {
				last := bpass[len(bpass)-1]
				copy(bpass[1:], bpass)
				bpass[0] = last
			}
		case parts[0] == "reverse":
			for i, j := util.Atoi(parts[2]), util.Atoi(parts[4]); i < j; i, j = i+1, j-1 {
				bpass[i], bpass[j] = bpass[j], bpass[i]
			}
		case parts[0] == "move":
			x, y := util.Atoi(parts[2]), util.Atoi(parts[5])
			v := bpass[x]
			if y > x {
				copy(bpass[x:y], bpass[x+1:])
			} else {
				copy(bpass[y+1:x+1], bpass[y:])
			}
			bpass[y] = v
		default:
			panic("unknown op: " + op)
		}
	}
	return string(bpass)
}
