package day22

import (
	"fmt"

	"github.com/rafael-luigi-bekkema/advent-of-code-2016/util"
)

type Node struct {
	x, y                   int
	size, used, avail, use int
}

func parseNodes(input []string) []Node {
	nodes := make([]Node, len(input)-2)
	for i, line := range input[2:] {
		n := &nodes[i]
		fmt.Sscanf(line, "/dev/grid/node-x%d-y%d %dT %dT %dT %d%%",
			&n.x, &n.y, &n.size, &n.used, &n.avail, &n.use)
	}
	return nodes
}

func day22a(input []string) int {
	nodes := parseNodes(input)
	var total int
	for i, n1 := range nodes {
		for _, n2 := range nodes[i+1:] {
			if n1.used > 0 && n2.avail >= n1.used || n2.used > 0 && n1.avail >= n2.used {
				total++
			}
		}
	}
	return total
}

func day22b(input []string) int {
	nodes := parseNodes(input)
	last := &nodes[len(nodes)-1]
	height := last.y + 1
	width := last.x + 1
	target := &nodes[last.x*height]
	tsize := target.used

	var empty func(node *Node, hist [][2]int, moves int) int
	empty = func(node *Node, hist [][2]int, moves int) int {
		for _, offs := range [][2]int{{0, -1}, {1, 0}, {0, 1}, {-1, 0}} {
			dx, dy := node.x+offs[0], node.y+offs[1]
			if dx < 0 || dx > width-1 || dy < 0 || dy > height-1 {
				continue
			}
			if util.In([2]int{dx, dy}, hist...) {
				continue
			}
			to := &nodes[dx*height+dy]
			if to.avail > node.used {

			} else {
				empty(to, append(hist, [2]int{dx, dy}), moves+1)
			}
		}
		node.used = 0
		return moves
	}

	_ = width
	_ = tsize
	var total int
	for x := target.x; x > 0; x-- {
		to := &nodes[(x-1)*height]
		if to.avail < tsize {
			total += empty(to, [][2]int{{target.x, target.y}}, 0)
		}
		total++
		target.used = 0
		to.used += tsize
		target = to
	}

	return total
}
