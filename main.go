// https://en.wikipedia.org/wiki/Dijkstra%27s_algorithm

package dijkstra_go

import "math"

var vertices = []int{0, 1, 2, 3, 4, 5}

type Edge struct {
	v0       int
	v1       int
	distance float64
}

var edges = []Edge{
	{0, 1, 1},
	{1, 2, 2},
	{2, 3, 3},
	{0, 4, 7},
	{1, 4, 9},
	{4, 5, 10},
	{2, 5, 13},
	{3, 5, 14},
}

var source = 0
var target = 3

func main() {
	dist := make([]float64, len(edges))

	// A value of -1 means "undefined"
	prev := make([]float64, len(edges))

	// We use a map data structure as a set. The values in this map are
	// irrelevant.
	Q := make(map[int]int)

	for _, v := range vertices {
		dist[v] = math.Inf(1)
		prev[v] = -1
		Q[v] = 0
	}

	dist[source] = 0
}
