package ostovtree

import (
	"sort"
)

type Edge struct {
	from, to int
	weight   float64
}

type DSU struct {
	parent []int
	rank   []int
}

func NewDSU(n int) *DSU {
	parent := make([]int, n)
	rank := make([]int, n)
	for i := range parent {
		parent[i] = i
	}
	return &DSU{parent, rank}
}

func (d *DSU) Find(x int) int {
	if d.parent[x] != x {
		d.parent[x] = d.Find(d.parent[x])
	}
	return d.parent[x]
}

func (d *DSU) Union(x, y int) bool {
	rootX := d.Find(x)
	rootY := d.Find(y)
	if rootX == rootY {
		return false
	}
	if d.rank[rootX] < d.rank[rootY] {
		d.parent[rootX] = rootY
	} else if d.rank[rootX] > d.rank[rootY] {
		d.parent[rootY] = rootX
	} else {
		d.parent[rootY] = rootX
		d.rank[rootX]++
	}
	return true
}

func Kruskal(n int, edges []Edge) (float64, []Edge) {
	sort.Slice(edges, func(i, j int) bool {
		return edges[i].weight < edges[j].weight
	})

	dsu := NewDSU(n)
	var mst []Edge
	totalWeight := 0.0

	for _, e := range edges {
		if dsu.Union(e.from, e.to) {
			mst = append(mst, e)
			totalWeight += e.weight
		}
	}

	return totalWeight, mst
}
