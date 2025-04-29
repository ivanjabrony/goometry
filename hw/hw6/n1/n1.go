package hw6_n1

import (
	"gogeom/mathobjects/voronoi"
	"gogeom/matrixes"
	"math/rand/v2"
)

const (
	bottom_range = -400
	top_range    = 400
	n            = 20

	width, height, hasNumbers = 1000, 1000, false
)

func randRange(min, max int) int {
	return rand.IntN(max-min) + min
}

func Hw6_n1_a() {
	dc := matrixes.NewDrawingContext(width, height, hasNumbers)

	points := make([][2]float64, n)

	for i := range n {
		points[i] = [2]float64{
			float64(randRange(bottom_range, top_range)),
			float64(randRange(bottom_range, top_range))}
		dc.DrawPoint(points[i][0], points[i][1], 3)
		dc.Stroke()
	}

	sites := make([]voronoi.Vertex, n)

	for i := range n {
		sites = append(sites, voronoi.Vertex{X: points[i][0], Y: points[i][1]})
	}

	bbox := voronoi.BBox{Xl: -width / 2, Xr: width / 2, Yt: -height / 2, Yb: height / 2}

	diagram := voronoi.ComputeDiagram(sites, bbox, true)

	for _, cell := range diagram.Cells {
		for _, hedge := range cell.Halfedges {
			dc.DrawLine(hedge.GetStartpoint().X, hedge.GetStartpoint().Y, hedge.GetEndpoint().X, hedge.GetEndpoint().Y)
			dc.Stroke()
		}
	}

	dc.SavePNG("hw/hw6/n1/task_a.png")
}
