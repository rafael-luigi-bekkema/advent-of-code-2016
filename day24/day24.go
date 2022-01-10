package day24

import (
	"bytes"
	"fmt"
	"math"

	"github.com/rafael-luigi-bekkema/advent-of-code-2016/util"
)

const (
	wall byte = '#'
)

type Coord struct {
	x, y int
}

type Graph struct {
	nodes map[Coord]struct{}
	edges map[Coord]map[Coord]float64
}

func (g *Graph) AddNode(n Coord) {
	if g.nodes == nil {
		g.nodes = make(map[Coord]struct{})
	}
	g.nodes[n] = struct{}{}
}

func (g *Graph) AddEdge(n, e Coord, v float64) {
	if g.edges == nil {
		g.edges = make(map[Coord]map[Coord]float64)
	}
	if g.edges[n] == nil {
		g.edges[n] = make(map[Coord]float64)
	}
	g.edges[n][e] = v
}

var inf = math.Inf(0)

func dijkstra(graph Graph, source Coord) (dist map[Coord]float64, prev map[Coord]Coord) {
	Q := make([]Coord, 0, len(graph.nodes))
	dist = map[Coord]float64{}
	prev = map[Coord]Coord{}
	for node := range graph.nodes {
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

		for v, l := range graph.edges[u] {
			if !util.In(v, Q...) {
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

func path(dijkstras map[int]map[Coord]float64, keys map[int]Coord, source int, total, mindist float64,
	cpath []int, b bool) float64 {

	if len(keys) == len(cpath) {
		if b {
			dest := dijkstras[source]
			return total + dest[keys[0]]
		}
		return total
	}
	if mindist != 0 && total >= mindist {
		return 0
	}

	dist := dijkstras[source]
	var minresult float64
	for num, dest := range keys {
		if util.In(num, cpath...) {
			continue
		}
		r := path(dijkstras, keys, num, total+dist[dest], mindist, append(cpath, num), b)
		if minresult == 0 || (r != 0 && r < minresult) {
			minresult = r
		}
		if mindist == 0 || (r != 0 && r < mindist) {
			mindist = r
		}

	}
	return minresult
}

func day24a(tmap string, b bool) float64 {
	bmap := []byte(tmap)
	width := bytes.Index(bmap, []byte{'\n'})
	bmap = bytes.ReplaceAll(bmap, []byte{'\n'}, nil)
	height := len(bmap) / width

	var graph Graph
	keys := map[int]Coord{}
	for i, c := range bmap {
		x, y := i%width, i/width
		node := Coord{x, y}
		if '0' <= c && c <= '9' {
			keys[int(c-'0')] = node
		}
		graph.AddNode(node)
		for _, offs := range [][2]int{{0, -1}, {1, 0}, {0, 1}, {-1, 0}} {
			dx, dy := x+offs[0], y+offs[1]
			if dx < 0 || dx > width-1 || dy < 0 || dy > height-1 {
				continue
			}
			newi := dy*width + dx
			if bmap[newi] != wall {
				graph.AddEdge(node, Coord{dx, dy}, 1)
			}
		}
	}

	dijkstras := map[int]map[Coord]float64{}
	for num, node := range keys {
		dijkstras[num], _ = dijkstra(graph, node)
	}

	source := 0
	total := path(dijkstras, keys, source, 0, 0, []int{source}, b)
	return total
}

func Day24a() {
	res := day24a(util.Input(24), false)
	fmt.Println("day24 a:", res)
}

func Day24b() {
	res := day24a(util.Input(24), true)
	fmt.Println("day24 b:", res)
}
