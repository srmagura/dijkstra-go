// https://en.wikipedia.org/wiki/Dijkstra%27s_algorithm

package main

import (
	"fmt"
	"math"
)

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

	{0, 3, 100},
}

var source = 0
var target = 3

func findShortestPathCore() []int {
	dist := make([]float64, len(edges))

	// A value of -1 means "undefined"
	prev := make([]int, len(edges))

	// We use a map data structure as a set. The values in this map are
	// irrelevant.
	Q := make(map[int]int)

	for _, v := range vertices {
		dist[v] = math.Inf(1)
		prev[v] = -1
		Q[v] = 0
	}

	dist[source] = 0

	for len(Q) > 0 {
		u := -1

		for v := range Q {
			if u == -1 || dist[v] < dist[u] {
				u = v
			}
		}

		if u == target {
			break
		}

		delete(Q, u)

		for _, edge := range edges {
			var v int

			if edge.v0 == u {
				v = edge.v1
			} else if edge.v1 == u {
				v = edge.v0
			} else {
				continue
			}

			_, ok := Q[v]

			if !ok {
				continue
			}

			alt := dist[u] + edge.distance

			if alt < dist[v] {
				dist[v] = alt
				prev[v] = u
			}

		}
	}

	return prev
}

func findShortestPath() []int {
	prev := findShortestPathCore()

	var path []int
	u := target

	for u != -1 {
		path = append([]int{u}, path...)
		u = prev[u]
	}

	return path
}

func main() {
	path := findShortestPath()

	fmt.Println("The shortest path is:")

	for _, u := range path {
		fmt.Printf("%d\n", u)
	}
}
