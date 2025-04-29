package hw6_n2

import (
	"gogeom/mathobjects/distance"
	"gogeom/mathobjects/voronoi"
	"gogeom/matrixes"
	"image/color"
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

func Hw6_n1_b() {
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

	var min_dot1 [2]float64
	var min_dot2 [2]float64
	var min_distance float64 = 1e10

	for _, edge := range diagram.Edges {
		if edge.LeftCell != nil && edge.RightCell != nil {
			dot1 := [2]float64{edge.LeftCell.Site.X, edge.LeftCell.Site.Y}
			dot2 := [2]float64{edge.RightCell.Site.X, edge.RightCell.Site.Y}
			dist := distance.CalculateDistance2Dots(dot1, dot2)
			if dist < min_distance {
				min_dot1 = dot1
				min_dot2 = dot2
				min_distance = dist
			}
		}
	}

	dc.DrawLine(min_dot1[0], min_dot1[1], min_dot2[0], min_dot2[1])
	dc.Stroke()
	dc.SavePNG("hw/hw6/n2/task_a.png")
}

func Hw6_n1_c() {
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

	for _, edge := range diagram.Edges {
		if edge.LeftCell != nil && edge.RightCell != nil {
			dc.SetColor(color.RGBA{0, 0, 100, 100})
			dc.DrawLine(edge.LeftCell.Site.X, edge.LeftCell.Site.Y,
				edge.RightCell.Site.X, edge.RightCell.Site.Y)
			dc.Stroke()
		}
	}
	dc.SavePNG("hw/hw6/n2/task_c.png")
}
