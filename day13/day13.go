package day13

import (
	"fmt"
	"math"

	. "github.com/rafael-luigi-bekkema/advent-of-code-2016/util"
)

type Coord struct {
	x, y int
}

var inf = math.Inf(0)

func dijkstra(nodes map[Coord]struct{}, edges map[Coord]map[Coord]float64, source Coord) (dist map[Coord]float64, prev map[Coord]Coord) {
	Q := make([]Coord, 0, len(nodes))
	dist = map[Coord]float64{}
	prev = map[Coord]Coord{}
	for node := range nodes {
		dist[node] = inf
		Q = append(Q, node)
	}
	dist[source] = 0

	for len(Q) > 0 {
		var minI int
		for i, u := range Q {
			if i == 0 || dist[u] < dist[Q[minI]] {
				minI = i
			}
		}
		u := Q[minI]
		Q[minI], Q[len(Q)-1] = Q[len(Q)-1], Q[minI]
		Q = Q[:len(Q)-1]

		for v, l := range edges[u] {
			if !In(v, Q...) {
				continue
			}
			alt := dist[u] + l
			if alt < dist[v] {
				dist[v] = alt
				prev[v] = u
			}
		}
	}
	return dist, prev
}

func day13a(favNum int, target Coord) (int, int) {
	isOpen := func(c Coord) bool {
		x, y := c.x, c.y
		val := x*x + 3*x + 2*x*y + y + y*y + favNum
		count := 0
		for _, r := range fmt.Sprintf("%b", val) {
			if r == '1' {
				count++
			}
		}
		return count%2 == 0
	}

	nodes := map[Coord]struct{}{}
	edges := map[Coord]map[Coord]float64{}
	for y := 0; y <= target.y+10; y++ {
		for x := 0; x <= target.x+10; x++ {
			coord := Coord{x, y}
			nodes[coord] = struct{}{}
			for _, offset := range [][2]int{{0, -1}, {-1, 0}, {1, 0}, {0, 1}} {
				rx := x + offset[0]
				ry := y + offset[1]
				if rx < 0 || ry < 0 {
					continue
				}
				if rc := (Coord{rx, ry}); isOpen(rc) {
					if edges[coord] == nil {
						edges[coord] = map[Coord]float64{}
					}
					edges[coord][rc] = 1
				}
			}
		}
	}

	dist, _ := dijkstra(nodes, edges, Coord{1, 1})
	var count int
	for _, distance := range dist {
		if distance <= 50 {
			count++
		}
	}
	return int(dist[target]), count
}
