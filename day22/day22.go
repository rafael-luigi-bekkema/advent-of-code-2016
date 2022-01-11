package day22

import (
	"fmt"

	"github.com/rafael-luigi-bekkema/advent-of-code-2016/dijkstra"
	"github.com/rafael-luigi-bekkema/advent-of-code-2016/util"
)

type Node struct {
	x, y       int
	size, used int
}

func (n Node) avail() int {
	return n.size - n.used
}

func (n *Node) String() string {
	return fmt.Sprintf("node-x%d-y%d", n.x, n.y)
}

func parseNodes(input []string) []Node {
	nodes := make([]Node, len(input)-2)
	var use, avail int
	for i, line := range input[2:] {
		n := &nodes[i]
		fmt.Sscanf(line, "/dev/grid/node-x%d-y%d %dT %dT %dT %d%%",
			&n.x, &n.y, &n.size, &n.used, &avail, &use)
	}
	return nodes
}

func day22a(input []string) int {
	nodes := parseNodes(input)
	var total int
	for i, n1 := range nodes {
		for _, n2 := range nodes[i+1:] {
			if n1.used > 0 && n2.avail() >= n1.used || n2.used > 0 && n1.avail() >= n2.used {
				total++
			}
		}
	}
	return total
}

func Day22b() {
	res := day22b(util.Lines(22))
	fmt.Println("day 22b:", res)
}

func buildGraph(nodes [][]*Node, skip *Node) (graph dijkstra.Graph[*Node], rdy map[[2]*Node]struct{}) {
	rdy = make(map[[2]*Node]struct{})
	for y, row := range nodes {
		for x, node := range row {
			if node == skip {
				continue
			}
			graph.AddNode(node)
			for _, offs := range [][2]int{{0, -1}, {1, 0}, {0, 1}, {-1, 0}} {
				dx, dy := x+offs[0], y+offs[1]
				if dx < 0 || dx > len(row)-1 || dy < 0 || dy > len(nodes)-1 {
					continue
				}
				from := nodes[dy][dx]
				if from.used <= node.size {
					if from.used <= node.avail() {
						// node can be offloaded here
						// don't add this path
						rdy[[2]*Node{node, from}] = struct{}{}
					}
					graph.AddEdge(from, node, 1)
				}
			}
		}
	}
	return
}

func snakeMove(prevs map[*Node]*Node, cur *Node) (total int) {
	for {
		prev := prevs[cur]
		if prev == nil {
			break
		}
		// fmt.Printf("%v (used %d) -> %v (avail %d)\n", prev, prev.used, cur, cur.avail())
		if prev.used > cur.avail() {
			panic("disk overflow")
		}
		cur.used += prev.used
		prev.used = 0
		cur = prev
		total++
	}
	return total
}

func day22b(input []string) int {
	rnodes := parseNodes(input)
	last := &rnodes[len(rnodes)-1]

	// Move into multi-dimensional array to make it easier to work with
	nodes := make([][]*Node, last.y+1)
	for i := range rnodes {
		node := &rnodes[i]
		if nodes[node.y] == nil {
			nodes[node.y] = make([]*Node, last.x+1)
		}
		nodes[node.y][node.x] = node
	}

	target := nodes[0][len(nodes[0])-1]

	var total int
	for x := target.x - 1; x >= 0; x-- {
		source := nodes[0][x]
		// fmt.Println("source", source)
		graph, rdy := buildGraph(nodes, target)
		dest, prev := dijkstra.Run(graph, source)
		var mindest float64
		var minn *Node
		for rn := range rdy {
			if rn[0] == source {
				continue
			}
			if prev[rn[0]] != rn[1] {
				continue
			}
			if mindest == 0 || dest[rn[0]] < mindest {
				mindest = dest[rn[0]]
				minn = rn[0]
			}
		}
		if minn == nil {
			panic("no result")
		}
		total += snakeMove(prev, minn) + 1

		source.used += target.used
		target.used = 0
		target = source
	}

	return int(total)
}
