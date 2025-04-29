package ostovtree

import (
	"gogeom/mathobjects/voronoi"
	"math"
)

func FindOstovTree(points [][2]float64) [][2][2]float64 {
	n := len(points)
	width := 1000.0
	height := 1000.0

	sites := make([]voronoi.Vertex, n)

	for i := range n {
		sites = append(sites, voronoi.Vertex{points[i][0], points[i][1]})
	}

	bbox := voronoi.BBox{-width / 2, width / 2, -height / 2, height / 2}

	diagram := voronoi.ComputeDiagram(sites, bbox, true)
	edges := make([]Edge, 0)

	for _, edge := range diagram.Edges {
		if edge.LeftCell != nil && edge.RightCell != nil {
			x1 := edge.LeftCell.Site.X
			x2 := edge.RightCell.Site.X
			y1 := edge.LeftCell.Site.Y
			y2 := edge.RightCell.Site.Y

			edges = append(edges, Edge{[2]float64(x1, y1), [2]float64(x2, y2), math.Sqrt((x1-x2)*(x1-x2) + (y1-y2)*(y1-y2))})
		}
	}

	for i := 0; i < len(points); i++ {
		for j := 0; j < i; j++ {
			edges = append(edges, Edge{i, j, math.Sqrt((points[i][0]-points[j][0])*(points[i][0]-points[j][0]) + (points[i][1]-points[j][1])*(points[i][1]-points[j][1]))})
		}
	}

	_, ans := Kruskal(len(edges), edges)

	tree := make([][2][2]float64, 0)

	for _, e := range ans {
		tree = append(tree, [2][2]float64{points[e.to], points[e.from]})
	}

	return tree
}
