package day22

import "fmt"

func day22a(input []string) int {
	type Node struct {
		x, y                   int
		size, used, avail, use int
	}
	nodes := make([]Node, len(input)-2)
	for i, line := range input[2:] {
		n := &nodes[i]
		fmt.Sscanf(line, "/dev/grid/node-x%d-y%d %dT %dT %dT %d%%",
			&n.x, &n.y, &n.size, &n.used, &n.avail, &n.use)
		fmt.Println(*n)
	}
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
